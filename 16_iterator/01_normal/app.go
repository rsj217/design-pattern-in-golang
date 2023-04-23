package iterator

import "fmt"

type Iterator interface {
	HasNext() bool
	GetNext() *Coupon
}

type Collectioner interface {
	createIterator() Iterator
}

type CouponIterator struct {
	idx        int64
	couponList CouponList
}

func (ci *CouponIterator) HasNext() bool {
	return ci.idx < int64(len(ci.couponList))
}

func (ci *CouponIterator) GetNext() *Coupon {
	if ci.HasNext() {
		coupon := ci.couponList[ci.idx]
		ci.idx++
		return coupon
	}
	return nil
}

type Coupon struct {
	Id int64
}
type CouponList []*Coupon

func (cl CouponList) createIterator() Iterator {
	return &CouponIterator{0, cl}
}

func client() {
	couponList := CouponList{
		&Coupon{1},
		&Coupon{2},
		&Coupon{3},
		&Coupon{4},
	}

	for iter := couponList.createIterator(); iter.HasNext(); {
		coupon := iter.GetNext()
		fmt.Printf("couponId=%d\n", coupon.Id)
	}
}
