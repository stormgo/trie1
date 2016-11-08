package main

import (
	//"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethdb"
	"github.com/ethereum/go-ethereum/trie"
)

func makeTrieDb() (ethdb.Database, *trie.Trie) {
	// Create an empty trie
	//db, _ := ethdb.NewMemDatabase()
	db, _ := ethdb.NewLDBDatabase("/tmp74/adba", 256, 0)
	trie, _ := trie.New(common.Hash{}, db)
	key := common.FromHex("0x61")
	value := common.FromHex("0x62")

	trie.Update(key,value)
	trie.Commit()
	return db, trie
}

func main() {
	makeTrieDb()
}
