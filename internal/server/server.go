package server

import (
	"context"
	"fmt"
	"net"

	"github.com/jasonkwh/wex-test-upstream/svc/purchasev1"
	"github.com/jasonkwh/wex-test/internal/config"
	"github.com/jasonkwh/wex-test/internal/data/model"
	"github.com/jasonkwh/wex-test/internal/data/pgx"
	"github.com/jasonkwh/wex-test/internal/utils"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type Server struct {
	repo     pgx.PurchaseRepository
	listener net.Listener
	server   *grpc.Server

	// please refer to
	// # https://stackoverflow.com/questions/65079032/grpc-with-mustembedunimplemented-method
	purchasev1.UnimplementedPurchaseServiceServer

	zl *zap.Logger
}

func NewServer(cfg config.ServerConfig, repo pgx.PurchaseRepository, zl *zap.Logger) (*Server, error) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", cfg.Port))
	if err != nil {
		return nil, fmt.Errorf("unable to listen: %v", err)
	}

	s := &Server{
		repo:     repo,
		listener: listener,
		zl:       zl,
	}

	// register grpc server
	s.server = grpc.NewServer()
	purchasev1.RegisterPurchaseServiceServer(s.server, s)

	s.zl.Info("purchase transaction server is running")
	return s, nil
}

func (s *Server) SavePurchaseTransaction(ctx context.Context, req *purchasev1.SavePurchaseRequest) (*purchasev1.GetPurchaseResponse, error) {
	id, err := s.repo.SavePurchase(ctx, model.Transaction{
		Description: req.Description,
		Date:        utils.ToFormattedDate(req.TransactionDate),
		Amount:      int(req.Amount),
	})
	if err != nil {
		return nil, err
	}

	return &purchasev1.GetPurchaseResponse{
		Description:     req.Description,
		TransactionDate: req.TransactionDate,
		Amount:          req.Amount,
		Id:              id,
	}, nil
}

func (s *Server) GetPurchaseTransaction(ctx context.Context, req *purchasev1.GetPurchaseRequest) (*purchasev1.GetPurchaseResponse, error) {
	// ts, err := s.repo.GetPurchase(ctx, req.Id)
	// if err != nil {
	// 	return nil, err
	// }

	return nil, nil
}

func (s *Server) Run() error {
	return s.server.Serve(s.listener)
}

func (s *Server) Close() error {
	var err error

	if s.server != nil {
		s.server.GracefulStop()
	}
	if s.listener != nil {
		err = s.listener.Close()
	}

	return err
}
