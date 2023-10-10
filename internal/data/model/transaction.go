package model

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"
)

type Transaction struct {
	Description string
	Date        string
	Amount      int
}

func (ts *Transaction) Hash() string {
	// make sure hash id won't conflict
	t := time.Now()

	h := sha256.New224()
	h.Write([]byte(fmt.Sprintf("%s-%s-%d-%s", ts.Description, ts.Date, ts.Amount, t)))
	return hex.EncodeToString(h.Sum(nil))
}
