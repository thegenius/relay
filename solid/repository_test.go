package solid

import (
	"github.com/siddontang/go-log/log"
	"testing"
	"time"
)



func TestStore(t *testing.T) {
	loopCount := 1000000
	repository := NewRepository("leveldb")

	start := time.Now()
	for i:=0; i<loopCount; i++ {
		Store(repository,"")
	}
	elapsed := time.Since(start)
	log.Info("used:", elapsed)
}
