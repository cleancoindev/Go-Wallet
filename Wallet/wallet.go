package Wallet

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"log"

	"github.com/mr-tron/base58"
	"golang.org/x/crypto/ripemd160"
)

type Wallet struct {
	PrivateKey ecdsa.PrivateKey
	PublicKey  []byte
	Token      int
}

type Wallets struct {
	Wallets map[string]*Wallet
}

var addresses []string

const (
	checksumLength = 4
	version        = byte(0x00)
)

func NewKeyPair() (ecdsa.PrivateKey, []byte) {
	curve := elliptic.P256()

	private, err := ecdsa.GenerateKey(curve, rand.Reader)
	if err != nil {
		log.Panic(err)
	}
	pub := append(private.PublicKey.X.Bytes(), private.PublicKey.Y.Bytes()...)
	return *private, pub
}

func MakeWallet() *Wallet {
	private, public := NewKeyPair()
	wallet := Wallet{private, public, 100}
	return &wallet
}

func (ws *Wallets) AddWallet() string {
	wallet := MakeWallet()
	address := fmt.Sprintf("%s", wallet.Address())
	fmt.Println(address)
	if ws.Wallets == nil {
		ws.Wallets = make(map[string]*Wallet)
		ws.Wallets[address] = wallet
	} else {
		ws.Wallets[address] = wallet
	}
	addresses = append(addresses, address)
	return address
}

func (w Wallet) Address() []byte {
	pubHash := PublicKeyHash(w.PublicKey)
	versionHash := append([]byte{version}, pubHash...)

	checksum := Checksum(versionHash)
	fullHash := append(versionHash, checksum...)
	address := Base58Encode(fullHash)

	// fmt.Printf("pub key: %x\n", w.PublicKey)
	// fmt.Printf("pub Hash: %x\n", pubHash)
	// fmt.Printf("address: %x\n", address)
	return address
}

func PublicKeyHash(pubKey []byte) []byte {
	pubHash := sha256.Sum256(pubKey)

	hasher := ripemd160.New()
	_, err := hasher.Write(pubHash[:])
	if err != nil {
		log.Panic(err)
	}

	publicRipMD := hasher.Sum(nil)
	return publicRipMD
}

func Checksum(payload []byte) []byte {
	firstHash := sha256.Sum256(payload)
	secondHash := sha256.Sum256(firstHash[:])
	return secondHash[:checksumLength]
}

func Base58Encode(input []byte) []byte {
	encode := base58.Encode(input)
	return []byte(encode)
}

func Base58Decode(input []byte) []byte {
	decode, err := base58.Decode(string(input[:]))
	if err != nil {
		log.Panic(err)
	}
	return decode
}

func GetAllAddresses() []string {
	return addresses
}

// func main() {
// 	wallets := &Wallets{}
// 	address := wallets.AddWallet()
// 	fmt.Println(address)
// 	for address := range wallets.Wallets {
// 		fmt.Println("address : ", address, "\nToken : ", wallets.Wallets[address].Token)
// 	}
// 	fmt.Println(wallets.Wallets)
// }
