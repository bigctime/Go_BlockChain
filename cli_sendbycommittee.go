package main

import (
	"fmt"
	"log"
)

//向委员会发送交易
func (cli *CLI) sendByCommittee(from, to string, amount int, nodeID, port string) {
	if !ValidateAddress(from) {
		log.Panic("ERROR: Sender address is not valid")
	}
	if !ValidateAddress(to) {
		log.Panic("ERROR: Recipient address is not valid")
	}

	var committeeNodes = fmt.Sprintf("localhost:%s", nodeID)

	bc := NewBlockchain(nodeID)
	UTXOSet := UTXOSet{bc}
	defer bc.db.Close()

	wallets, err := NewWallets(nodeID)
	if err != nil {
		log.Panic(err)
	}
	wallet := wallets.GetWallet(from)

	tx := NewUTXOTransaction(&wallet, to, amount, &UTXOSet)
	//主动连接委员会，发送交易
	sendTx(committeeNodes, tx)

	fmt.Println("Success!")
}
