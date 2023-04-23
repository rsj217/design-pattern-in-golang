package adapter

import (
	"fmt"
	"strconv"
	"strings"
)

type Benefiter interface {
	GetStepRule() string
}

var _ Benefiter = (*Coupon)(nil)

type Coupon struct {
	amountList   []int64
	discountList []int64
}

func (c *Coupon) addAmountRule(amountPrice, discountPrice int64) {
	c.amountList = append(c.amountList, amountPrice*100)
	c.discountList = append(c.discountList, discountPrice*100)

}
func (c *Coupon) GetStepRule() string {
	ans := make([]string, 0, len(c.amountList))
	for idx, v := range c.amountList {
		rule := fmt.Sprintf("%d-%d", v/100, c.discountList[idx]/100)
		ans = append(ans, rule)
	}
	return strings.Join(ans, ",")
}

// var _ Benefiter = (*Voucher)(nil) 编译失败
type Voucher struct {
	discount int64
}

func (v *Voucher) GetDiscountPrice() int64 {
	return v.discount / 100
}

var _ Benefiter = (*VoucherCouponAdapter)(nil)

type VoucherCouponAdapter struct {
	voucher *Voucher
}

func (v *VoucherCouponAdapter) GetStepRule() string {
	discountPrice := v.voucher.GetDiscountPrice()
	return fmt.Sprintf("0-%d", discountPrice)
}

func GetDisAmountFromStepRule(amount int64, b Benefiter) int64 {
	rule := b.GetStepRule()
	rulesList := strings.Split(rule, ",")
	for idx := len(rulesList) - 1; 0 <= idx; idx-- {
		r := rulesList[idx]
		info := strings.Split(r, "-")
		amountConf, _ := strconv.Atoi(info[0])
		discountConf, _ := strconv.Atoi(info[1])
		if int64(amountConf)*100 <= amount {
			return int64(discountConf)
		}
	}
	return 0
}

func client() {
	coupon := &Coupon{}
	coupon.addAmountRule(9, 2)
	coupon.addAmountRule(49, 5)
	coupon.addAmountRule(99, 10)
	amount := GetDisAmountFromStepRule(8, coupon)
	fmt.Printf("coupon discount=%d\n", amount)
	amount = GetDisAmountFromStepRule(100*100, coupon)
	fmt.Printf("coupon discount=%d\n", amount)

	voucherCoupon := &VoucherCouponAdapter{
		&Voucher{20 * 100},
	}
	amount = GetDisAmountFromStepRule(0, voucherCoupon)
	fmt.Printf("voucher discount=%d\n", amount)
}
