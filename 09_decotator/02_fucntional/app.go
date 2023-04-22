package decorator

import "fmt"

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
		ans := fn(id)
		fmt.Printf("After Fetch: LogWrapper record a log for coupon.id=%d\n", ans.id)
		return ans
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
