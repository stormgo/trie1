package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethdb"
	"github.com/ethereum/go-ethereum/trie"
)

func makeTrieDb() (ethdb.Database, *trie.Trie) {
	// Create an empty trie
	//db, _ := ethdb.NewMemDatabase()
	db, _ := ethdb.NewLDBDatabase("/tmp74/adba", 256, 0)
	trie, _ := trie.New(common.Hash{}, db)
	key := common.ParseData("dog")
	value := common.ParseData("cat")
	fmt.Println(value)

	trie.Update(key,value)
	trie.Commit()
	return db, trie
}

func main() {
	_, trie := makeTrieDb()
	key := common.ParseData("dog")
	value := trie.Get(key)
	fmt.Println(value)
}
