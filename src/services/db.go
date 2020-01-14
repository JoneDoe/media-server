package services

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/boltdb/bolt"

	"istorage/config"
	"istorage/models"
)

type Engine struct {
	db     *bolt.DB
	Bucket string
}

func InitDb() *Engine {
	db, err := bolt.Open(config.Config.Db.Database, 0600, nil)
	if err != nil {
		log.Fatal(err)
	}

	storage := &Engine{db, config.Config.Db.Bucket}
	storage.CreateBucket()

	return storage
}

func (s *Engine) Close() {
	s.db.Close()
}

func (s *Engine) CreateRecord(attachment *models.Attachment) error {
	err := s.db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(s.Bucket))

		buf, err := json.Marshal(attachment.ToJson())
		if err != nil {
			return err
		}

		return bucket.Put([]byte(attachment.Uuid), buf)
	})

	defer s.Close()

	return err
}

func (s *Engine) DeleteRecord(uuid string) error {
	err := s.db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(s.Bucket))

		return bucket.Delete([]byte(uuid))
	})

	defer s.Close()

	return err
}

func (s *Engine) GetRecord(uuid string) (map[string]interface{}, error) {
	data := make(map[string]interface{})

	err := s.db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(s.Bucket))
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

func (s *Engine) CreateBucket() {
	s.db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(s.Bucket))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}
		return nil
	})
}
