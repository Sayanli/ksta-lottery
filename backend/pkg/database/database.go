package database

type DB interface {
	Insert(value uint) uint

	GetValueByKey(key uint) (uint, error)
}

func NewDB() *MemoryDB {
	mDB := &MemoryDB{
		storage: make(map[uint]uint),
		last_id: 0,
		wait:    make(chan struct{}),
	}
	go func() {
		mDB.wait <- struct{}{}
	}()
	return mDB
}
