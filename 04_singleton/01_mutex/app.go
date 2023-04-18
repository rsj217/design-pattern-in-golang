package singleton

import (
	"fmt"
	"sync"
	"time"
)

type DBConn struct{}

var dbConn *DBConn
var mutex = &sync.Mutex{}

func GetDB() *DBConn {
	if dbConn == nil {
		mutex.Lock()
		defer mutex.Unlock()
		if dbConn == nil {
			dbConn = &DBConn{}
			fmt.Println("dbConn create")
		} else {
			fmt.Println("dbConn already create")
		}
	} else {
		fmt.Println("dbConn already create")
	}
	return dbConn
}

func Client() {
	for i := 0; i < 100; i++ {
		go func() {
			GetDB()
		}()
	}
	time.Sleep(time.Second * 2)
}
