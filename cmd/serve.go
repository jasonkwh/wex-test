package cmd

import (
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/jasonkwh/wex-test/internal/data/pgx"
	"github.com/jasonkwh/wex-test/internal/exchangerate"
	"github.com/jasonkwh/wex-test/internal/server"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "starts the purchase transaction server",
	Run:   serve,
}

func serve(cmd *cobra.Command, args []string) {
	// make grateful close pool
	var clPool []io.Closer

	zl, err := initZapLogger()
	if err != nil {
		log.Fatal("unable to start zap logger")
	}

	// start db repo
	re, err := pgx.CreatePurchaseRepository(cfg.Database, zl)
	if err != nil {
		zl.Fatal("unable to create purchase repository", zap.Error(err))
	}

	// create exchange rate retriever
	ret := exchangerate.NewRetriever(&http.Client{}, cfg.ExchangeRate.Within)

	// start server
	srv, err := server.NewServer(cfg.Server, re, ret, zl)
	if err != nil {
		zl.Fatal("unable to start the purchase transaction server", zap.Error(err))
	}
	go func() {
		if err := srv.Run(); err != nil && err != http.ErrServerClosed {
			zl.Fatal("purchase transaction server failed to serve", zap.Error(err))
		}
	}()
	clPool = append(clPool, srv)

	zl.Info("server started")

	// handle shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	if err := gracefulClose(clPool); err != nil {
		zl.Error("failed to close the server", zap.Error(err))
	}
	os.Exit(0)
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
