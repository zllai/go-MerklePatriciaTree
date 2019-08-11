package kvstore

import "github.com/syndtr/goleveldb/leveldb"

type LevelDB struct {
	backend *leveldb.DB
}

func NewLevelDB(path string) (*LevelDB, error) {
	backend, err := leveldb.OpenFile(path, nil)
	if err != nil {
		return nil, err
	}
	return &LevelDB{backend}, nil
}

func (db *LevelDB) Get(key []byte) ([]byte, error) {
	return db.backend.Get(key, nil)
}

func (db *LevelDB) Put(key, value []byte) error {
	return db.backend.Put(key, value, nil)
}

func (db *LevelDB) Has(key []byte) bool {
	has, _ := db.backend.Has(key, nil)
	return has
}

func (db *LevelDB) Delete(key []byte) error {
	return db.backend.Delete(key, nil)
}

func (db *LevelDB) BatchPut(kvs [][2][]byte) error {
	batch := new(leveldb.Batch)
	for i := range kvs {
		batch.Put(kvs[i][0], kvs[i][1])
	}
	return db.backend.Write(batch, nil)
}

func (db *LevelDB) Close() {
	db.backend.Close()
}
