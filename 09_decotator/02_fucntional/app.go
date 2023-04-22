package decorator

import (
	"fmt"
	"time"
)

type Fetcher interface {
	Fetch(id int64) *Coupon
}
type GetFunc func(id int64) *Coupon

type Coupon struct {
	id int64
}

func Fetch(id int64) *Coupon {
	fmt.Println("fetch data from db")
	return &Coupon{id}
}

func LogWrapper(fn GetFunc) GetFunc {
	return func(id int64) *Coupon {
		fmt.Println("Before Fetch: ")
		defer fmt.Println()
		ans := fn(id)
		fmt.Printf("After Fetch: LogWrapper record a log for coupon.id=%d\n", ans.id)
		return ans
	}
}

func ProfileWrapper(fn GetFunc) GetFunc {
	return func(id int64) *Coupon {
		start := time.Now()
		defer func() {
			fmt.Printf("It takes %v ms", time.Since(start).Milliseconds())
		}()
		time.Sleep(time.Millisecond * 200)
		return fn(id)
	}
}

func client1() {
	coupon := Fetch(10)
	fmt.Printf("CouponManager.Fetch  coupon.id=%d\n", coupon.id)
}

func client2() {
	coupon := LogWrapper(Fetch)(10)
	fmt.Printf("CouponManager.Fetch  coupon.id=%d\n", coupon.id)

}

func client3() {
	coupon := ProfileWrapper(LogWrapper(Fetch))(10)
	fmt.Printf("CouponManager.Fetch  coupon.id=%d\n", coupon.id)
}
