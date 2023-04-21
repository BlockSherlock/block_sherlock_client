package main

import (
	"client/eth"
	"client/helpers"
	"client/server"
	"fmt"
	"os"
	"strings"
	"sync"
)

func main() {
	var wallets sync.Map
	port := helpers.ReadFromTerminal("enter your free port")
	fmt.Println(port)
	apiKey := helpers.ReadFromTerminal("enter your apiKey")

	memory := server.Memory{Wallets: &wallets, ApiKey: apiKey}
	loadedWallets := eth.LoadWallets(&wallets)
	if len(loadedWallets) == 0 {
		fmt.Println("no wallets in './w' folder")
		return
	}
	os.WriteFile("forBlockSherlock_bot.txt", []byte(strings.Join(loadedWallets, "\n")), 0644)

	server.ServerStart(port, 1000, 1000, &memory, apiKey)
}
