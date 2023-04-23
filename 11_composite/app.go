package composite

import "fmt"

type Benefiter interface {
	GetDiscount() int64
}

type Coupon struct {
	discount int64
}

func (c *Coupon) GetDiscount() int64 {
	return c.discount
}

type Card struct {
	coupons []*Coupon
}

func (c *Card) GetDiscount() int64 {
	ans := int64(0)
	for _, v := range c.coupons {
		ans += v.GetDiscount()
	}
	return ans
}

type Benefit struct {
	benefits []Benefiter
}

func (b *Benefit) Add(benefiter Benefiter) {
	b.benefits = append(b.benefits, benefiter)
}

func (b *Benefit) GetDiscount() int64 {
	ans := int64(0)
	for _, v := range b.benefits {
		ans += v.GetDiscount()
	}
	return ans
}

func client() {
	c1 := &Coupon{10}
	c2 := &Coupon{5}
	card1 := &Card{
		[]*Coupon{{9}, {9}, {9}},
	}
	card2 := &Card{
		[]*Coupon{{5}},
	}

	benefit := Benefit{}
	benefit.Add(c1)
	benefit.Add(c2)
	benefit.Add(card1)
	benefit.Add(card2)

	discount := benefit.GetDiscount()
	fmt.Printf("benefit discount=%d", discount)

}
