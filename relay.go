package main

import (
	_ "github.com/go-sql-driver/mysql"
	"log"
	"relay/command"
	"relay/dump"
)

func main() {
	commandArgs := command.ParseArgs()
	log.Println("args: ", commandArgs)

	dump.Dump()
}

//func main() {

//}
