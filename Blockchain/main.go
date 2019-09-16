package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"

	wallet "github.com/sunnyradadiya/P2P/Wallet"
)

type Blockchain struct {
	blocks []*Block
}

type Block struct {
	Hash             []byte
	Token            int
	PrevHash         []byte
	SenderAddress    string
	RecipientAddress string
}

func (b *Block) DeriveHash() {
	info := bytes.Join([][]byte{b.PrevHash}, []byte{})
	hash := sha256.Sum256(info)
	b.Hash = hash[:]
}

func CreateBlock(token int, prevHash []byte, senderAddress string, recipientAddress string) *Block {
	block := &Block{[]byte{}, token, prevHash, senderAddress, recipientAddress}
	block.DeriveHash()
	return block
}

func (chain *Blockchain) AddBlock(token int, senderAddress string, recipientAddress string) {
	prevBlock := chain.blocks[len(chain.blocks)-1]
	new := CreateBlock(token, prevBlock.Hash, senderAddress, recipientAddress)
	chain.blocks = append(chain.blocks, new)

}

func Genesis() *Block {
	return CreateBlock(0, []byte{}, "", "")
}

func InitBlockchain() *Blockchain {
	return &Blockchain{[]*Block{Genesis()}}
}

func main() {

	wallets := &wallet.Wallets{}

	for j := 0; j < 3; j++ {
		wallets.AddWallet()
	}

	chain := InitBlockchain()

	addresses := wallet.GetAllAddresses()
	for i := 0; i < len(addresses)-1; i++ {
		val := 10 * (i + 1)
		if wallets.Wallets[addresses[i]].Token >= val {
			wallets.Wallets[addresses[i]].Token = wallets.Wallets[addresses[i]].Token - val
			wallets.Wallets[addresses[i+1]].Token = wallets.Wallets[addresses[i+1]].Token + val
			chain.AddBlock(val, addresses[i], addresses[i+1])
		}
	}

	fmt.Println()
	for _, block := range chain.blocks {
		fmt.Printf("Token in Block: %d\n", block.Token)
		fmt.Printf("Previous Hash: %x\n", block.PrevHash)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Printf("Sender Address: %s\n", block.SenderAddress)
		fmt.Printf("Recipient Address: %s\n", block.RecipientAddress)
		fmt.Println()
	}

	for address := range wallets.Wallets {
		fmt.Println("\naddress : ", address, "\nToken : ", wallets.Wallets[address].Token)
		fmt.Println("privatekey : ", wallets.Wallets[address].PrivateKey, "\npublickey : ", wallets.Wallets[address].PublicKey)
	}
}
