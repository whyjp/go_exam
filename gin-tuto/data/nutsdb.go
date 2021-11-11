package data

import (
	"encoding/json"
	"log"

	"github.com/xujiajun/nutsdb"
)

var (
	db     *nutsdb.DB
	bucket string

	fileDir string
)

func init() {
	opt := nutsdb.DefaultOptions
	fileDir = "./data/nutsdb"

	opt.Dir = fileDir
	opt.SegmentSize = 1024 * 1024 // 1MB
	db, _ = nutsdb.Open(opt)
}

func delete(bucket string, targetKey string) {
	if err := db.Update(
		func(tx *nutsdb.Tx) error {
			key := []byte(targetKey)
			return tx.Delete(bucket, key)
		}); err != nil {
		log.Fatal(err)
	}
}

func put(bucket string, targetKey string, val []byte) error {
	if err := db.Update(
		func(tx *nutsdb.Tx) error {
			key := []byte(targetKey)
			return tx.Put(bucket, key, val, 0)
		}); err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}
func Save(name string, keyR string, v interface{}) error {
	b, err := json.Marshal(v)
	if err != nil {
		log.Println("save error: ", err)
		return err
	}
	return put(name, keyR, b)
}
func readString(bucket string, targetKey string) string {
	var value []byte
	if err := db.View(func(tx *nutsdb.Tx) error {
		key := []byte(targetKey)
		e, err := tx.Get(bucket, key)
		if err != nil {
			return err
		}
		log.Println("val:", string(e.Value))
		value = e.Value
		return nil
	}); err != nil {
		log.Println(err)
	}
	return string(value)
}

func read(bucket string, targetKey string) interface{} {
	var value interface{}
	if err := db.View(
		func(tx *nutsdb.Tx) error {
			key := []byte(targetKey)
			e, err := tx.Get(bucket, key)
			if err != nil {
				return err
			}
			log.Println("val:", e.Value)
			value = e.Value
			return nil
		}); err != nil {
		log.Println(err)
	}
	return value
}
