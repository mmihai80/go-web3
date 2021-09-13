package wallet

import (
	"github.com/mmihai80/go-web3/dto"
)

type WalletInterface interface {
	Sign(tx dto.TransactionParameters) (string, error)
}
