package singleton

import (
	"fmt"
	"sync"
	"time"
)

type DBConn struct{}

var once sync.Once
var dbConn *DBConn

func GetDB() *DBConn {
	if dbConn == nil {
		once.Do(func() {
			dbConn = &DBConn{}
			fmt.Println("dbConn create")
		})
	} else {
		fmt.Println("dbConn already create")
	}
	return dbConn
}

func App() {
	for i := 0; i < 100; i++ {
		go func() {
			GetDB()
		}()
	}
	time.Sleep(time.Second * 2)
}
