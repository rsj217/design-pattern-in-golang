package dsp

import "fmt"

type Cacher interface {
	get()
	set()
}

type Memory struct {
}

func (r *Memory) get() {
	fmt.Println("Memory get")
}

func (r *Memory) set() {
	fmt.Println("Memory set")
}

type Redis struct {
}

func (r *Redis) get() {
	fmt.Println("redis get")
}

func (r *Redis) set() {
	fmt.Println("redis set")
}

type CacheManger struct{}

func (cm *CacheManger) GetCache(c Cacher) {
	c.get()
}
func (cm *CacheManger) SetCache(c Cacher) {
	c.set()
}

func client() {
	cm := &CacheManger{}
	cm.GetCache(&Redis{})
	cm.GetCache(&Memory{})
}
