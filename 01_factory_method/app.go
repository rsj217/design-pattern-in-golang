package factory_method

import (
	"fmt"
	"strconv"
	"strings"
)

type Benefiter interface {
	GetDisAmount() int64
}

// 直减券
type Voucher struct {
	rule string // "0-10-10"
}

func (v *Voucher) GetDisAmount() int64 {
	ruleConf := strings.Split(v.rule, "-")
	discountConf, _ := strconv.Atoi(ruleConf[1])
	return int64(discountConf)
}

// 折扣券
type Coupon struct {
	rule string // "99-80-20"
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

type Factory interface {
	CreateBenefiter(rule string) Benefiter
}

type CouponFactory struct {
}

func (cf CouponFactory) CreateBenefiter(rule string) Benefiter {
	return &Coupon{rule: rule}
}

// NewVoucherFactory Voucher 工厂函数(方法) 符合 golang 的编码规则
func NewVoucherFactory(rule string) Benefiter {
	return &Voucher{rule: rule}
}

type Config struct {
	price int64
	rule  string
}

func client() {
	conf := Config{int64(100), "100-70-20"}

	coupon := CouponFactory{}.CreateBenefiter(conf.rule)
	disAmount := coupon.GetDisAmount()
	fmt.Printf("Coupon price:%d  rule: %s disAmount: %d\n", conf.price, conf.rule, disAmount)

	conf = Config{int64(100), "0-10-10"}
	voucher := NewVoucherFactory(conf.rule)
	disAmount = voucher.GetDisAmount()
	fmt.Printf("Voucher price:%d  rule: %s disAmount: %d\n", conf.price, conf.rule, disAmount)
}
