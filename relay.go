package main

import (
	_ "github.com/go-sql-driver/mysql"
	"relay/dump"
)

func main() {
	//commandArgs := command.ParseArgs()
	//log.Println("args: ", commandArgs)

	dump.FetchMetaInfo()
}

