package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"net/http"
	"encoding/json"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	wallet "github.com/sunnyRK/GO-WALLET/Wallet"
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
	Timestamp string
}

var wallets wallet.Wallets
var chain Blockchain

func (b *Block) DeriveHash() {
	info := bytes.Join([][]byte{b.PrevHash}, []byte{})
	hash := sha256.Sum256(info)
	b.Hash = hash[:]
}

func CreateBlock(token int, prevHash []byte, senderAddress string, recipientAddress string) *Block {
	t := time.Now()
	block := &Block{[]byte{}, token, prevHash, senderAddress, recipientAddress, t.String()}
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

func addWallet(w http.ResponseWriter, r *http.Request) {
	// wallets := &wallet.Wallets{}
	w.Header().Set("Content-Type", "application/json")
	address := wallets.AddWallet()
	_ = json.NewDecoder(r.Body).Decode(address)
	addresses := wallet.GetAllAddresses()
	addresses = append(addresses, address)
	json.NewEncoder(w).Encode(&address)
}

func getAllWalletAddresses(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	addresses := wallet.GetAllAddresses()
	json.NewEncoder(w).Encode(addresses)
}

func getAllWalletDetails(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(wallets)
}

func TransferToken(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	token, _ := strconv.Atoi(vars["val"])
	sender := vars["sender"]
	recipient := vars["recipient"]
	_ = json.NewDecoder(r.Body).Decode(chain)
	if wallets.Wallets[sender].Token >= token {
		wallets.Wallets[sender].Token = wallets.Wallets[sender].Token - token
		wallets.Wallets[recipient].Token = wallets.Wallets[recipient].Token + token
		chain.AddBlock(token, sender, recipient)
	}
	json.NewEncoder(w).Encode(&chain.blocks)
}

func main() {
	chain.blocks = append(chain.blocks, Genesis())
	fmt.Println(chain)

	router := mux.NewRouter()
	router.HandleFunc("/addwallet", addWallet).Methods("POST")
	router.HandleFunc("/getAllWalletAddresses", getAllWalletAddresses).Methods("GET")
	router.HandleFunc("/getAllWalletDetails", getAllWalletDetails).Methods("GET")
	router.HandleFunc("/transferToken/{val}/{sender}/{recipient}", TransferToken).Methods("POST")
	http.ListenAndServe(":8000", router)
}

