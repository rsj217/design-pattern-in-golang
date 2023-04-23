package iterator

import "fmt"

type Coupon struct {
	Id int64
}

type CouponList []*Coupon

func (cl CouponList) Iterator() func() (*Coupon, bool) {
	idx := 0
	return func() (val *Coupon, ok bool) {
		if len(cl) <= idx {
			return
		}
		val, ok = cl[idx], true
		idx++
		return
	}

}

func client() {
	couponList := CouponList{
		&Coupon{1},
		&Coupon{2},
		&Coupon{3},
		&Coupon{4},
	}

	iter := couponList.Iterator()
	for {
		if coupon, ok := iter(); ok {
			fmt.Printf("couponId=%d\n", coupon.Id)
		}
		break
	}
}
