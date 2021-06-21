package main

import (
	"github.com/xujiajun/nutsdb"
	"log"
)

func main() {
	test()
}

func test() {
	opt := nutsdb.DefaultOptions
	opt.Dir = "/tmp/nutsdb"
	db, err := nutsdb.Open(opt)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Update(func(tx *nutsdb.Tx) error {
		key := []byte("name")
		val := []byte("abc")
		if err := tx.Put("", key, val, 0); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		log.Println("nutsdb update error:", err)
	}

	err = db.View(func(tx *nutsdb.Tx) error {
		key := []byte("name")
		if entry, err := tx.Get("", key); err != nil {
			return err
		} else {
			log.Println(string(entry.Value))
		}
		return nil
	})
	if err != nil {
		log.Println("nutsdb view error:", err)
	}
}