package bridge

import (
	"fmt"
	"strconv"
	"strings"
)

type Benefiter interface {
	GetDisAmount() int64
}

type Voucher struct {
	rule string // "0-10-10"
}

func NewVoucher(rule string) *Voucher {
	return &Voucher{rule: rule}
}

func (v *Voucher) GetDisAmount() int64 {
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

func (c *Coupon) GetDisAmount() int64 {
	ruleConf := strings.Split(c.rule, "-")
	priceConf, _ := strconv.Atoi(ruleConf[0])
	rebateConf, _ := strconv.Atoi(ruleConf[1])
	maxConf, _ := strconv.Atoi(ruleConf[2])
	disAmount := priceConf * rebateConf / 100
	if disAmount < maxConf {
		return int64(disAmount)
	}
	return int64(maxConf)
}

type BizProder interface {
	GetDisAmount() int64
	SetBenefiter(Benefiter)
}

type Hotel struct {
	benefiter Benefiter
}

func (h *Hotel) GetDisAmount() int64 {
	return h.benefiter.GetDisAmount()
}

func (h *Hotel) SetBenefiter(benefiter Benefiter) {
	h.benefiter = benefiter
}

type Scenic struct {
	benefiter Benefiter
}

func (h *Scenic) GetDisAmount() int64 {
	return h.benefiter.GetDisAmount()
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
	disAmount := hotel.GetDisAmount()
	fmt.Printf("hotel coupon disAmount=%d\n", disAmount)

	hotel.SetBenefiter(voucher)
	disAmount = hotel.GetDisAmount()
	fmt.Printf("hotel voucher disAmount=%d\n", disAmount)

	scenic := Scenic{}
	scenic.SetBenefiter(coupon)
	disAmount = scenic.GetDisAmount()
	fmt.Printf("scenic coupon disAmount=%d\n", disAmount)

	scenic.SetBenefiter(voucher)
	disAmount = scenic.GetDisAmount()
	fmt.Printf("scenic voucher disAmount=%d\n", disAmount)

}
