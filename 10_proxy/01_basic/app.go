package proxy

import "fmt"

type Fetcher interface {
	Fetch(id int64) *Coupon
}

type Coupon struct {
	id int64
}

var _ Fetcher = (*DB)(nil)

type DB struct{}

func (db *DB) Fetch(id int64) *Coupon {
	fmt.Println("fetch data from db")
	return &Coupon{id}
}

type RedisCacheProxy struct {
	fetcher Fetcher
}

func (rcw *RedisCacheProxy) Fetch(id int64) *Coupon {
	fmt.Println("\tredis cache")
	if id == 10 {
		fmt.Println("\t\tredis cache hit")
		return &Coupon{id}
	}
	fmt.Println("\t\t redis cache missing")
	return rcw.fetcher.Fetch(id)
}

type LocalCacheProxy struct {
	fetcher Fetcher
}

func (l *LocalCacheProxy) Fetch(id int64) *Coupon {
	fmt.Println("\tlocal cache")
	if id == 20 {
		fmt.Println("\t\tlocal cache hit")
		return &Coupon{id}
	}
	fmt.Println("\t\tlocal cache missing")
	ans := l.fetcher.Fetch(id)
	return ans
}

func NewCacheProxy() Fetcher {
	redis := &RedisCacheProxy{&DB{}}
	return &LocalCacheProxy{redis}
}

func client1() {
	var fetcher Fetcher = &DB{}

	coupon := fetcher.Fetch(21)
	fmt.Printf("coupon.id=%d\n", coupon.id)

}

func client2() {
	var fetcher Fetcher = &DB{}

	redisCacheProxy := &RedisCacheProxy{fetcher}
	coupon := redisCacheProxy.Fetch(21)
	fmt.Printf("coupon.id=%d\n", coupon.id)
}

func client3() {
	var fetcher Fetcher = &DB{}
	redisCacheProxy := &RedisCacheProxy{fetcher}
	localCacheProxy := &LocalCacheProxy{redisCacheProxy}
	coupon := localCacheProxy.Fetch(20)
	fmt.Printf("coupon.id=%d\n", coupon.id)
}

func client4() {
	cache := NewCacheProxy()
	coupon := cache.Fetch(21)
	fmt.Printf("coupon.id=%d\n", coupon.id)
}
