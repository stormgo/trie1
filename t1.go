package main

import (
	//"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethdb"
	"github.com/ethereum/go-ethereum/trie"
)


// makeTestTrie create a sample test trie to test node-wise reconstruction.
func makeTestTrie() (ethdb.Database, *trie.Trie, map[string][]byte) {
	// Create an empty trie
	db, _ := ethdb.NewMemDatabase()
	trie, _ := trie.New(common.Hash{}, db)

	// Fill it with some arbitrary data
	content := make(map[string][]byte)
	for i := byte(0); i < 255; i++ {
		// Map the same data under multiple keys
		key, val := common.LeftPadBytes([]byte{1, i}, 32), []byte{i}
		content[string(key)] = val
		trie.Update(key, val)

		key, val = common.LeftPadBytes([]byte{2, i}, 32), []byte{i}
		content[string(key)] = val
		trie.Update(key, val)

		// Add some other data to inflate th trie
		for j := byte(3); j < 13; j++ {
			key, val = common.LeftPadBytes([]byte{j, i}, 32), []byte{j, i}
			content[string(key)] = val
			trie.Update(key, val)
		}
	}
	trie.Commit()

	// Return the generated trie
	return db, trie, content
}


func main() {

makeTestTrie()
//_, mytrie, _ := makeTestTrie()


/*
    var trie Trie
	key := make([]byte, 32)
	value := common.FromHex("0x61")
	fmt.Println(value)
	trie.Update(key, value)
	value = trie.Get(key)
	fmt.Println(value)
*/
}
