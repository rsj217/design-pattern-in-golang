package proxy

import "fmt"

type Coupon struct {
	id int64
}

func Fetch(id int64) *Coupon {
	fmt.Println("fetch data from db")
	return &Coupon{id}
}

type GetFunc func(id int64) *Coupon

func RedisCacheProxy(fn GetFunc) GetFunc {
	return func(id int64) *Coupon {
		fmt.Println("\tredis cache")
		if id == 10 {
			fmt.Println("\t\tredis cache hit")
			return &Coupon{id}
		}
		fmt.Println("\tredis cache missing")
		ans := fn(id)
		return ans
	}
}

func LocalCacheProxy(fn GetFunc) GetFunc {
	return func(id int64) *Coupon {
		fmt.Println("\tlocal cache")
		if id == 20 {
			fmt.Println("\t\tlocal cache hit")
			return &Coupon{id}
		}
		fmt.Println("\tlocal cache missing")
		ans := fn(id)
		return ans
	}
}

func client1() {
	coupon := Fetch(21)
	fmt.Printf("coupon.id=%d\n", coupon.id)

	couponRedisProxy := RedisCacheProxy(Fetch)
	coupon = couponRedisProxy(10)
	fmt.Printf("coupon.id=%d\n", coupon.id)

	couponLocalProxy := LocalCacheProxy(couponRedisProxy)
	coupon = couponLocalProxy(20)
	fmt.Printf("coupon.id=%d\n", coupon.id)
}

func client2() {
	coupon := LocalCacheProxy(RedisCacheProxy(Fetch))(21)
	fmt.Printf("coupon.id=%d\n", coupon.id)
}
