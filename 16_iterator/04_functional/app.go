package iterator

import "fmt"

type Coupon struct {
	Id int64
}

type CouponList []*Coupon

func (cl CouponList) Do(fn func(coupon *Coupon)) {
	for _, v := range cl {
		fn(v)
	}
}

func Client() {
	couponList := CouponList{
		&Coupon{1},
		&Coupon{2},
		&Coupon{3},
		&Coupon{4},
	}

	couponList.Do(func(coupon *Coupon) {
		fmt.Printf("couponId=%d\n", coupon.Id)
	})
}
