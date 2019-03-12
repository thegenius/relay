package dump

import (
	"database/sql"
	"fmt"
	"github.com/JamesStewy/go-mysqldump"
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

func Dump() {

	// Open connection to database
	username := "root"
	password := "123456"
	hostname := "localhost"
	port := "3306"
	dbname := "test"

	dumpDir := "dumps"  // you should create this directory
	prepareDir(dumpDir)
	dumpFilenameFormat := fmt.Sprintf("%s-20060102T150405", dbname)   // accepts time layout string and add .sql at the end of file

	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", username, password, hostname, port, dbname))
	if err != nil {
		fmt.Println("Error opening database: ", err)
		return
	}

	// Register database with mysqldump
	dumper, err := mysqldump.Register(db, dumpDir, dumpFilenameFormat)
	if err != nil {
		fmt.Println("Error registering databse:", err)
		return
	}

	// Dump database to file
	resultFilename, err := dumper.Dump()
	if err != nil {
		fmt.Println("Error dumping:", err)
		return
	}
	fmt.Printf("File is saved to %s", resultFilename)

	// Close dumper and connected database
	dumper.Close()
}
