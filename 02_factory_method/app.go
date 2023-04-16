package factory_method

import (
	"fmt"
	"strconv"
	"strings"
)

type Benefiter interface {
	GetDiscount() int64
}

// 直减券
type Voucher struct {
	rule string // "0-10-10"
}

func (v *Voucher) GetDiscount() int64 {
	ruleConf := strings.Split(v.rule, "-")
	discountConf, _ := strconv.Atoi(ruleConf[1])
	return int64(discountConf)
}

// 折扣券
type Coupon struct {
	rule string // "99-80-20"
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

type Factory interface {
	CreateBenefiter(string) Benefiter
}

type VoucherFactory struct {
	Factory
}

func (vf VoucherFactory) CreateBenefiter(rule string) Benefiter {
	return &Voucher{rule: rule}
}

type CouponFactory struct {
	Factory
}

func (cf CouponFactory) CreateBenefiter(rule string) Benefiter {
	return &Coupon{rule: rule}
}

type Config struct {
	price int64
	rule  string
}

func App() {
	conf := Config{int64(100), "100-70-20"}

	coupon := CouponFactory{}.CreateBenefiter(conf.rule)
	discount := coupon.GetDiscount()
	fmt.Printf("Coupon price:%d  rule: %s discount: %d\n", conf.price, conf.rule, discount)

	conf = Config{int64(100), "0-10-10"}
	voucher := VoucherFactory{}.CreateBenefiter(conf.rule)
	discount = voucher.GetDiscount()
	fmt.Printf("Voucher price:%d  rule: %s discount: %d\n", conf.price, conf.rule, discount)
}
