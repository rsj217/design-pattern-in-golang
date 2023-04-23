package bridge

import (
	"fmt"
	"strconv"
	"strings"
)

type Benefiter interface {
	GetDiscount() int64
}

type Voucher struct {
	rule string // "0-10-10"
}

func NewVoucher(rule string) *Voucher {
	return &Voucher{rule: rule}
}

func (v *Voucher) GetDiscount() int64 {
	ruleConf := strings.Split(v.rule, "-")
	discountConf, _ := strconv.Atoi(ruleConf[1])
	return int64(discountConf)
}

type Coupon struct {
	rule string // "99-80-20"
}

func NewCoupon(rule string) *Coupon {
	return &Coupon{rule: rule}
}

func (c *Coupon) GetDiscount() int64 {
	ruleConf := strings.Split(c.rule, "-")
	priceConf, _ := strconv.Atoi(ruleConf[0])
	rebateConf, _ := strconv.Atoi(ruleConf[1])
	maxConf, _ := strconv.Atoi(ruleConf[2])
	discount := priceConf * rebateConf / 100
	if discount < maxConf {
		return int64(discount)
	}
	return int64(maxConf)
}

type BizProder interface {
	GetDiscount() int64
	SetBenefiter(Benefiter)
}

type Hotel struct {
	benefiter Benefiter
}

func (h *Hotel) GetDiscount() int64 {
	return h.benefiter.GetDiscount()
}

func (h *Hotel) SetBenefiter(benefiter Benefiter) {
	h.benefiter = benefiter
}

type Scenic struct {
	benefiter Benefiter
}

func (h *Scenic) GetDiscount() int64 {
	return h.benefiter.GetDiscount()
}

func (h *Scenic) SetBenefiter(benefiter Benefiter) {
	h.benefiter = benefiter
}

type Config struct {
	price int64
	rule  string
}

func client() {
	conf := &Config{int64(100), "100-70-20"}
	coupon := NewCoupon(conf.rule)

	conf = &Config{int64(100), "0-10-10"}
	voucher := NewVoucher(conf.rule)

	hotel := Hotel{}
	hotel.SetBenefiter(coupon)
	discount := hotel.GetDiscount()
	fmt.Printf("hotel coupon discount=%d\n", discount)

	hotel.SetBenefiter(voucher)
	discount = hotel.GetDiscount()
	fmt.Printf("hotel voucher discount=%d\n", discount)

	scenic := Scenic{}
	scenic.SetBenefiter(coupon)
	discount = scenic.GetDiscount()
	fmt.Printf("scenic coupon discount=%d\n", discount)

	scenic.SetBenefiter(voucher)
	discount = scenic.GetDiscount()
	fmt.Printf("scenic voucher discount=%d\n", discount)

}
