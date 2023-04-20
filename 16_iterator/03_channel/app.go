package iterator

import "fmt"

type Coupon struct {
	Id int64
}

type CouponList []*Coupon

func (cl CouponList) Iterator() <-chan *Coupon {

	c := make(chan *Coupon)
	go func() {
		for _, v := range cl {
			c <- v
		}
		close(c)
	}()
	return c
}

func Client() {
	couponList := CouponList{
		&Coupon{1},
		&Coupon{2},
		&Coupon{3},
		&Coupon{4},
	}

	for coupon := range couponList.Iterator() {
		fmt.Printf("couponId=%d\n", coupon.Id)
	}
}
