package anti_dsp

import "fmt"

type Redis struct {
}

func (r *Redis) get() {
	fmt.Println("redis get")
}

func (r *Redis) set() {
	fmt.Println("redis set")
}

type CacheManger struct{}

func (cm *CacheManger) GetCache(redis *Redis) {
	redis.get()
}
func (cm *CacheManger) SetCache(redis *Redis) {
	redis.set()
}

func client() {
	cm := &CacheManger{}
	cm.GetCache(&Redis{})
}
