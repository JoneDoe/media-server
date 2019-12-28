package utils

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/boltdb/bolt"

	a "github.com/JoneDoe/istorage/attachment"
)

const dbFile = "iStorage.db"
const filesBucket = "files"

type Store struct {
	//blocks []*Block
	tip []byte
	db  *bolt.DB
}

type BlockchainIterator struct {
	currentHash []byte
	db          *bolt.DB
}

func openDb() {
	db, err := bolt.Open(dbFile, 0600, nil)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	//Store.
}

func (s *Store) CreateRecord(attachment *a.Attachment) error {
	return s.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(filesBucket))

		buf, err := json.Marshal(attachment.ToJson())
		if err != nil {
			return err
		}

		return b.Put([]byte(attachment.Uuid), buf)
	})
}

func (s *Store) GetRecord(uuid string) a.Attachment {
	data := a.Attachment{}

	err := s.db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(filesBucket))
		value := bucket.Get([]byte(uuid))

		err := json.Unmarshal(value, &data)
		if err != nil {
			return err
		}

		fmt.Printf("The answer is: %s\n", data)
		return nil
	})

	checkError(err)

	return data
}

/*func NewBlockchain() *Store {
	var tip []byte
	db, err := bolt.Open(dbFile, 0600, nil)
	checkError(err)

	err = db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(filesBucket))

		if b == nil {
			genesis := NewGenesisBlock()
			b, err := tx.CreateBucket([]byte(filesBucket))
			checkError(err)
			err = b.Put(genesis.Hash, genesis.Serialize())
			checkError(err)
			err = b.Put([]byte("l"), genesis.Hash)
			checkError(err)
			tip = genesis.Hash
		} else {
			tip = b.Get([]byte("l"))
		}

		return nil
	})
	checkError(err)

	bc := Store{tip, db}

	return &bc
}*/

/*func (bc *Store) Iterator() *BlockchainIterator {
	bci := &BlockchainIterator{bc.tip, bc.db}

	return bci
}

func (i *BlockchainIterator) Next() *Block {
	var block *Block

	err := i.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(filesBucket))
		encodedBlock := b.Get(i.currentHash)
		block = DeserializeBlock(encodedBlock)

		return nil
	})

	checkError(err)

	i.currentHash = block.PrevBlockHash

	return block
}*/

func checkError(err error) {
	if err != nil {
		log.Panic(err)
	}
}
