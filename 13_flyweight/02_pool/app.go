package flyweight

import (
	"fmt"
	"sync"
)

type Conn struct {
	addr string
	port int64
}

var connPool *sync.Pool

func init() {
	connPool = &sync.Pool{
		New: func() interface{} {
			return &Conn{"localhost", 3306}
		},
	}
}

func App() {
	conn1 := connPool.Get().(*Conn)
	fmt.Printf("&conn1=%p, addr=%s port=%d\n", conn1, conn1.addr, conn1.port)
	connPool.Put(conn1)

	conn2 := connPool.Get().(*Conn)
	fmt.Printf("&conn2=%p, addr=%s port=%d\n", conn2, conn2.addr, conn2.port)

	fmt.Println("conn1==conn2", conn1 == conn2)

	conn3 := connPool.Get().(*Conn)
	fmt.Printf("&conn3=%p, addr=%s port=%d\n", conn3, conn3.addr, conn3.port)

	fmt.Println("conn2==conn3", conn2 == conn3)
}
