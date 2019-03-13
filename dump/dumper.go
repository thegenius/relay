package dump

import (
	"github.com/siddontang/go-mysql/canal"
	"github.com/siddontang/go-mysql/client"
	"log"
	"os"
)

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func prepareDir(dirName string) {
	exists, err := PathExists(dirName)
	if err != nil {
		log.Fatalln("error: ", err)
	}
	if !exists {
		mkdirError := os.MkdirAll(dirName, os.ModePerm)
		if mkdirError != nil {
			log.Fatalln("failed to create dir: ", mkdirError)
		}
	}
}

type MyEventHandler struct {
	canal.DummyEventHandler
}

func (h *MyEventHandler) OnRow(e *canal.RowsEvent) error {
	log.Printf("%s %v\n", e.Action, e.Rows)
	return nil
}

func FetchMetaInfo() {
	log.Printf("fetch")
	conn, err := client.Connect("127.0.0.1:3306", "canal", "123456", "test")
	if err != nil {
		log.Fatal("error: ", err)
	}

	err = conn.Ping()
	if err != nil {
		log.Fatal("failed to ping")
	}

	// Insert
	r, queryError := conn.Execute(`select COLUMN_NAME,DATA_TYPE,COLUMN_COMMENT 
from information_schema.COLUMNS where table_name = 'entity_basic_info' and table_schema = 'test'`)
	if queryError != nil {
		log.Fatal("error: ", queryError)
	}
	log.Print(r)
}

func Dump() {




	// Open connection to database
	//username := "root"
	//password := "123456"
	//hostname := "localhost"
	//port := "3306"
	//dbname := "test"

	dumpDir := "dumps"  // you should create this directory
	prepareDir(dumpDir)

	cfg := canal.NewDefaultConfig()
	cfg.Addr = "127.0.0.1:3306"
	cfg.User = "canal"
	cfg.Password = "123456"
	// We only care table canal_test in test db
	cfg.Dump.TableDB = "test"
	cfg.Dump.Tables = []string{"entity_basic_info"}


	c, err := canal.NewCanal(cfg)
	if err != nil {
		log.Fatal("failed to open canal!")
	}



	//func (h *MyEventHandler) String() string {
	//	return "MyEventHandler"
	//}

	// Register a handler to handle RowsEvent
	c.SetEventHandler(&MyEventHandler{})

	// Start canal
	c.Run()
}
