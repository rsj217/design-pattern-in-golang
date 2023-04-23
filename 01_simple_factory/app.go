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

func NewVoucher(rule string) *Voucher {
	return &Voucher{rule: rule}
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

func SimpleFactory(conf *Config) Benefiter {
	switch conf.flag {
	case "Voucher":
		return NewVoucher(conf.rule)
	case "Coupon":
		return NewCoupon(conf.rule)
	}
	panic("flag err")
}

type Config struct {
	flag  string
	price int64
	rule  string
}

func client() {
	conf := &Config{"Coupon", int64(100), "100-70-20"}

	coupon := SimpleFactory(conf)
	discount := coupon.GetDiscount()
	fmt.Printf("Coupon price:%d  rule: %s discount: %d\n", conf.price, conf.rule, discount)

	conf = &Config{"Voucher", int64(100), "0-10-10"}

	voucher := SimpleFactory(conf)
	discount = voucher.GetDiscount()
	fmt.Printf("Voucher price:%d  rule: %s discount: %d\n", conf.price, conf.rule, discount)
}
