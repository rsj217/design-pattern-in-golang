package flyweight

import "fmt"

// 抽象 actId 和 actName 这部分可以共享缓存的内容
//type UserRecord struct {
//	id      int64
//	actId   int64
//	actName string
//}

type Act struct {
	Id   int64
	Name string
}

type UserRecord struct {
	id  int64
	act *Act
}

func client() {
	act := &Act{1, "大促"}
	ur1 := UserRecord{id: 1, act: act}
	ur2 := UserRecord{id: 2, act: act}

	fmt.Println(ur1.act == ur2.act)

}
