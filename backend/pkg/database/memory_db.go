package database

import (
	"errors"
)

type MemoryDB struct {
	storage map[uint]uint
	last_id uint
	wait    chan struct{}
}

func (db *MemoryDB) Insert(value uint) uint {
	<-db.wait
	db.last_id++
	db.storage[db.last_id] = value
	go func() {
		db.wait <- struct{}{}
	}()
	return db.last_id
}

func (db *MemoryDB) GetValueByKey(key uint) (uint, error) {
	value, exists := db.storage[key]
	if exists {
		return value, nil
	} else {
		return 0, errors.New("token is invalid")
	}
}
