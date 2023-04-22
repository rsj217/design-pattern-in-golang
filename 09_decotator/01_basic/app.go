package decorator

import "fmt"

type Fetcher interface {
	Fetch(id int64) *Coupon
}

type Coupon struct {
	id int64
}

type CouponManager struct{}

func (cm *CouponManager) Fetch(id int64) *Coupon {
	fmt.Println("fetch data from db")
	return &Coupon{id}
}

type LogWrapper struct {
	fetcher Fetcher
}

func (rcw *LogWrapper) Fetch(id int64) *Coupon {
	fmt.Println("Before Fetch: ")
	ans := rcw.fetcher.Fetch(id)
	fmt.Printf("After Fetch: LogWrapper record a log for coupon.id=%d\n", ans.id)
	return ans
}

func client1() {
	cm := CouponManager{}
	coupon := cm.Fetch(10)
	fmt.Printf("CouponManager.Fetch  coupon.id=%d\n", coupon.id)
}

func client2() {
	var cm Fetcher = &CouponManager{}
	cm = &LogWrapper{
		fetcher: cm,
	}
	coupon := cm.Fetch(10)
	fmt.Printf("CouponManager.Fetch  coupon.id=%d\n", coupon.id)

}
