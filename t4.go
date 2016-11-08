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
	// _, trie := makeTrieDb()
	db, err := ethdb.NewLDBDatabase("/tmp74/adba", 256, 0)
	if err != nil {
		panic(err)
	}
	for it := db.NewIterator(); it.Next(); {
		fmt.Println("key = ", it.Key())
		fmt.Println("value = ", it.Value())
	}

/*
	key := common.ParseData("dog")
	//value := trie.Get(key)
	value, err := db.Get(key)
	if err != nil {
		fmt.Println("Unexpected error: %v", err)
	}
	fmt.Println(value)
*/
}
