package solid

import (
	"github.com/siddontang/go-log/log"
	"github.com/syndtr/goleveldb/leveldb"
)

type Repository struct {
	DB *leveldb.DB
	FileName string
}

func NewRepository(fileName string) Repository {
	db, err :=leveldb.OpenFile(fileName, nil)
	if err != nil {
		log.Fatal("failed to open leveldb: ", err)
	}
	repository := Repository{
		DB: db,
		FileName: fileName,
	}
	return repository
}

func Store(repository Repository, value string) string  {
	db := repository.DB

	err := db.Put([]byte("hello"), []byte("world"), nil)
	if err != nil {
		log.Fatal("failed to put value: ", err)
	}

	data, readError := db.Get([]byte("hello"), nil)
	if readError != nil {
		log.Fatal("read error: ", readError)
	}
	return string(data)
}
