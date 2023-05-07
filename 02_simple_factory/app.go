package simple_factory

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
	disAmount := coupon.GetDisAmount()
	fmt.Printf("Coupon price:%d  rule: %s disAmount: %d\n", conf.price, conf.rule, disAmount)

	conf = &Config{"Voucher", int64(100), "0-10-10"}

	voucher := SimpleFactory(conf)
	disAmount = voucher.GetDisAmount()
	fmt.Printf("Voucher price:%d  rule: %s disAmount: %d\n", conf.price, conf.rule, disAmount)
}
