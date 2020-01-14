package services

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/boltdb/bolt"

	"istorage/models"
)

const dbFile = "iStorage.db"
const filesBucket = "files"

type Store struct {
	db *bolt.DB
}

func InitDb() *Store {
	db, err := bolt.Open(dbFile, 0600, nil)
	if err != nil {
		log.Fatal(err)
	}

	storage := &Store{db: db}
	storage.CreateBucket(filesBucket)

	return storage
}

func (s *Store) Close() {
	s.db.Close()
}

func (s *Store) CreateRecord(attachment *models.Attachment) error {
	err := s.db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(filesBucket))

		buf, err := json.Marshal(attachment.ToJson())
		if err != nil {
			return err
		}

		return bucket.Put([]byte(attachment.Uuid), buf)
	})

	defer s.Close()

	return err
}

func (s *Store) DeleteRecord(uuid string) error {
	err := s.db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(filesBucket))

		return bucket.Delete([]byte(uuid))
	})

	defer s.Close()

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

	defer s.Close()

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
