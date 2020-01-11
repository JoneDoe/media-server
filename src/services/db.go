package services

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/boltdb/bolt"

	a "istorage/attachment"
)

const dbFile = "iStorage.db"
const filesBucket = "files"

type Store struct {
	db *bolt.DB
}

type BlockchainIterator struct {
	currentHash []byte
	db          *bolt.DB
}

func InitDb() *Store {
	db, err := bolt.Open(dbFile, 0600, nil)
	if err != nil {
		log.Fatal(err)
	}

	storage := Store{db: db}
	storage.CreateBucket(filesBucket)

	return &storage
}

func (s *Store) Close() {
	s.db.Close()
}

func (s *Store) CreateRecord(attachment *a.Attachment) error {
	err := s.db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(filesBucket))

		buf, err := json.Marshal(attachment.ToJson())
		if err != nil {
			return err
		}

		return bucket.Put([]byte(attachment.Uuid), buf)
	})

	s.Close()

	return err
}

func (s *Store) DeleteRecord(uuid string) error {
	err := s.db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(filesBucket))

		return bucket.Delete([]byte(uuid))
	})

	s.Close()

	return err
}

func (s *Store) GetRecord(uuid string) (map[string]interface{}, error) {
	data := make(map[string]interface{})

	err := s.db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(filesBucket))
		value := bucket.Get([]byte(uuid))

		err := json.Unmarshal(value, &data)
		if err != nil {
			return err
		}

		return nil
	})

	s.Close()

	return data, err
}

func (s *Store) CreateBucket(bucketName string) {
	s.db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(bucketName))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}
		return nil
	})
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
