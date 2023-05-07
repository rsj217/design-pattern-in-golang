package composite

import "fmt"

type Benefiter interface {
	GetDisAmount() int64
}

type Coupon struct {
	disAmount int64
}

func (c *Coupon) GetDisAmount() int64 {
	return c.disAmount
}

type Card struct {
	coupons []*Coupon
}

func (c *Card) GetDisAmount() int64 {
	ans := int64(0)
	for _, v := range c.coupons {
		ans += v.GetDisAmount()
	}
	return ans
}

type Benefit struct {
	benefits []Benefiter
}

func (b *Benefit) Add(benefiter Benefiter) {
	b.benefits = append(b.benefits, benefiter)
}

func (b *Benefit) GetDisAmount() int64 {
	ans := int64(0)
	for _, v := range b.benefits {
		ans += v.GetDisAmount()
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

	disAmount := benefit.GetDisAmount()
	fmt.Printf("benefit disAmount=%d", disAmount)

}
