package models

import (
	"crypto/ecdsa"
)

type Wallet struct {
	PrivateKey *ecdsa.PrivateKey
}
