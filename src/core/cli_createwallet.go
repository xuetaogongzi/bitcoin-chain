package core

import "fmt"

func (cli *CLI) createWallet(nodeID string) {
	walltes, _ := NewWallets(nodeID)
	address := walltes.createWallet()
	walltes.SaveToFile(nodeID)

	fmt.Printf("Your new address:%s\n", address)
}
