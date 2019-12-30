package main

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"time"
)

type Data struct {
	sender   string
	receiver string
	message  string
	amount   uint
}

type Block struct {
	index     int
	timestamp int64
	data      Data
	prevHash  string
	hash      string
}

func (b *Block) getHash() string {
	dataString, err := json.Marshal(b.data)
	if err != nil {
		panic(err)
	}

	hash := sha256.New()
	hash.Write([]byte(string(dataString) + b.prevHash + string(b.index) + string(b.timestamp)))

	return fmt.Sprintf("%x", hash.Sum(nil))
}

type Blockchain struct {
	chain []Block
}

func (b *Blockchain) addBlock(data Data) {
	index := len(b.chain)

	var prevHash string
	if index > 0 {
		prevHash = b.chain[index-1].hash
	}

	b.chain = append(b.chain, newBlock(index, data, prevHash))
}

func (b *Blockchain) isValid() bool {
	for i, block := range b.chain {
		if block.hash != block.getHash() {
			return false
		}

		if i > 0 && block.prevHash != b.chain[i-1].hash {
			return false
		}
	}

	return true
}

func main() {
	var blockchain Blockchain
	blockchain.addBlock(newData("adrian", "carole", "1'm 5uch 4 933k!", 1337))
	fmt.Println(blockchain.chain[0])

	blockchain.addBlock(newData("adrian", "carole", "The answer to life and everything", 42))
	fmt.Println(blockchain.chain[1])

	blockchain.addBlock(newData("carole", "adrian", "", 77))
	fmt.Println(blockchain.chain[2])

	fmt.Println("The blockchain contains ", len(blockchain.chain), " blocks")
	fmt.Println("The blockchain validity is", blockchain.isValid())
}

func newData(sender string, receiver string, message string, amount uint) Data {
	if sender == receiver {
		panic("You cannot send a message to yourself")
	}

	return Data{
		sender:   sender,
		receiver: receiver,
		message:  message,
		amount:   amount,
	}
}

func newBlock(index int, data Data, prevHash string) Block {
	block := Block{
		index:     index,
		timestamp: time.Now().Unix(),
		data:      data,
		prevHash:  prevHash,
	}
	block.hash = block.getHash()

	return block
}
