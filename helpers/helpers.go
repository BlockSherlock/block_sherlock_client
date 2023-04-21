package helpers

import (
	"bufio"
	"crypto/ecdsa"
	"fmt"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}
func StrToPK(s string) (privateKey *ecdsa.PrivateKey, address common.Address) {
	privateKey, err := crypto.HexToECDSA(s)
	if err != nil {
		return
	}
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return
	}
	address = crypto.PubkeyToAddress(*publicKeyECDSA)

	return
}
func ReadFromTerminal(startText string) string {
	fmt.Print("\n" + startText + ": ")
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	return strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(text, "\n", ""), "\r", ""), " ", "")
}
