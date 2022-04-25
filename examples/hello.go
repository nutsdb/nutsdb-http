package main

import (
	"log"

	"nutshttp"

	"github.com/xujiajun/nutsdb"
)

func main() {
	opt := nutsdb.DefaultOptions
	opt.Dir = "/tmp/nutsdb"
	db, err := nutsdb.Open(opt)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	// Enable auth
	nutshttp.EnableAuth = true
	nutshttp.SetSecret("TbI7O6yEdEYa")
	go func() {
		if err := nutshttp.Enable(db); err != nil {
			panic(err)
		}
	}()

	testDB(db)

	select {}
}

func testDB(db *nutsdb.DB) {
	var (
		bucket = "bucket001"

		key   = []byte("foo")
		value = []byte("bar")
	)

	if err := db.Update(func(tx *nutsdb.Tx) error {
		if err := tx.SAdd(bucket, key, value); err != nil {
			return err
		}

		if err := tx.SAdd(bucket, key, []byte("bar2")); err != nil {
			return err
		}

		_ = tx.RPush(bucket, []byte("key1"), []byte("value1"))
		_ = tx.RPush(bucket, []byte("key1"), []byte("value2"))

		return nil

	}); err != nil {
		log.Fatal(err)
	}

	if err := db.View(func(tx *nutsdb.Tx) error {
		items, err := tx.SMembers(bucket, key)
		if err != nil {
			return err
		}

		for _, item := range items {
			log.Printf("item: %s", item)
		}
		return nil
	}); err != nil {

		log.Fatal(err)
	}
}
