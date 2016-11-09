package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethdb"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/ethereum/go-ethereum/trie"
)

func makeTrieDb() (ethdb.Database, *trie.Trie) {
	// Create an empty trie
	//db, _ := ethdb.NewMemDatabase()
	db, _ := ethdb.NewLDBDatabase("/tmp74/adba", 256, 0)
	trie, _ := trie.New(common.Hash{}, db)
	key := common.ParseData("a")
	value := common.ParseData("b")
	fmt.Println(key)
	fmt.Println(value)

	trie.Update(key,value)
	trie.Commit()
	return db, trie
}

func mydecode(input []byte) {
	mycount, _ := rlp.CountValues(input)
	fmt.Println(mycount)
/*
	var s string
	err := rlp.DecodeBytes(input, &s)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("Decoded value: %#v\n", s)
	}
*/
}


func main() {
	// makeTrieDb()

	db, err := ethdb.NewLDBDatabase("/tmp74/adba", 256, 0)
	if err != nil {
		panic(err)
	}
	var itkey, itval []byte
	for it := db.NewIterator(); it.Next(); {

		itkey = it.Key()
		mydecode(itkey)

		itval = it.Value()
		mydecode(itval)

		//fmt.Println("key = ", itkey)
		//fmt.Println("value = ", itval)
	}
}
