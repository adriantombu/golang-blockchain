package main

import (
	"testing"
)

func TestNewData(t *testing.T) {
	sender := "adrian"
	receiver := "carole"
	message := "1'm 5uch 4 933k!"
	amount := uint(1337)

	got := newData(sender, receiver, message, amount)

	if got.sender != sender {
		t.Errorf("got %s and expected %s", got.sender, sender)
	}

	if got.receiver != receiver {
		t.Errorf("got %s and expected %s", got.receiver, receiver)
	}

	if got.message != message {
		t.Errorf("got %s and expected %s", got.message, message)
	}

	if got.amount != amount {
		t.Errorf("got %d and expected %d", got.amount, amount)
	}
}

func TestNewDataSameSenderReceiver(t *testing.T) {
	sender := "adrian"
	message := "1'm 5uch 4 933k!"
	amount := uint(1337)

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("TestNewWrongData should have panicked!")
		}
	}()

	newData(sender, sender, message, amount)
}

func TestNewBlock(t *testing.T) {
	var prevHash string
	index := 0
	data := newData("adrian", "carole", "1'm 5uch 4 933k!", 1337)

	got := newBlock(index, data, prevHash)

	if got.index != index {
		t.Errorf("index - got %d and expected %d", got.index, index)
	}

	if got.prevHash != prevHash {
		t.Errorf("prevHash - got %s and expected %s", got.prevHash, prevHash)
	}

	if got.timestamp == 0 {
		t.Errorf("timestamp - got %d and expected %d", got.timestamp, 0)
	}

	if got.hash == "" {
		t.Errorf("hash - got %s and expected %s", got.hash, "")
	}

	if got.data.sender != data.sender {
		t.Errorf("sender - got %s and expected %s", got.data.sender, data.sender)
	}

	if got.data.receiver != data.receiver {
		t.Errorf("receiver - got %s and expected %s", got.data.receiver, data.receiver)
	}

	if got.data.message != data.message {
		t.Errorf("message - got %s and expected %s", got.data.message, data.message)
	}

	if got.data.amount != data.amount {
		t.Errorf("amount - got %d and expected %d", got.data.amount, data.amount)
	}
}

func TestAddBlock(t *testing.T) {
	sender := "adrian"
	receiver := "carole"
	message := "1'm 5uch 4 933k!"
	amount := uint(1337)

	data := newData(sender, receiver, message, amount)

	var blockchain Blockchain
	blockchain.addBlock(data)
	blockchain.addBlock(data)
	blockchain.addBlock(data)
	blockchain.addBlock(data)
	blockchain.addBlock(data)

	if len(blockchain.chain) != 5 {
		t.Errorf("blockchain length - got %d and expected %d", len(blockchain.chain), 5)
	}
}

func TestChainIsValid(t *testing.T) {
	sender := "adrian"
	receiver := "carole"
	message := "1'm 5uch 4 933k!"
	amount := uint(1337)

	data := newData(sender, receiver, message, amount)

	var blockchain Blockchain
	blockchain.addBlock(data)
	blockchain.addBlock(data)

	if blockchain.isValid() != true {
		t.Error("blockchain validity - got false and expected true")
	}
}

func TestChainIsNotValid(t *testing.T) {
	sender := "adrian"
	receiver := "carole"
	message := "1'm 5uch 4 933k!"
	amount := uint(1337)

	data := newData(sender, receiver, message, amount)

	var blockchain Blockchain
	blockchain.addBlock(data)
	blockchain.addBlock(data)

	blockchain.chain[0].hash = "A RANDOM HASH"

	if blockchain.isValid() != false {
		t.Error("blockchain validity - got true and expected false")
	}
}
