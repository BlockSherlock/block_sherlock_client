package server

import (
	"client/helpers"
	"client/models"
	"encoding/json"
	"fmt"
	"math/big"
	"net/http"
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type Memory struct {
	Wallets *sync.Map
	ApiKey  string
}
type SignatureReqData struct {
	Addr    common.Address
	Tx      *types.Transaction
	ChainId *big.Int
}

type Response struct {
	Tx *types.Transaction
}

func (mem *Memory) signHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	w.Header().Set("Content-Type", "application/json")
	var signatureReqData SignatureReqData
	err := json.NewDecoder(r.Body).Decode(&signatureReqData)
	helpers.CheckErr(err)
	targetAccount, ok := mem.Wallets.Load(signatureReqData.Addr)
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	account := targetAccount.(models.Wallet)
	fmt.Println(signatureReqData.Addr, signatureReqData.Tx)
	fmt.Println(account, signatureReqData.Tx.GasPrice())
	signedTx, err := types.SignTx(signatureReqData.Tx, types.NewEIP155Signer(signatureReqData.ChainId), account.PrivateKey)
	helpers.CheckErr(err)
	responseBody := Response{Tx: signedTx}
	responceByte, err := json.Marshal(responseBody)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(responceByte)
}
func pingHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	w.WriteHeader(http.StatusOK)
}
