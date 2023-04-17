package _1_deepcopy

import "fmt"

type CouponTpl struct {
	Name string
	Rule string
}

func (ct CouponTpl) DeepCopy() *CouponTpl {
	return &CouponTpl{
		ct.Name,
		ct.Rule,
	}
}

func NewCouponTpl(name, rule string) *CouponTpl {
	return &CouponTpl{name, rule}
}

type Act struct {
	Name      string
	CouponTpl *CouponTpl
}

func (a *Act) DeepCopy() *Act {
	return &Act{
		a.Name,
		a.CouponTpl.DeepCopy(),
	}
}

func Copy() {
	act1 := Act{
		"大促活动",
		&CouponTpl{"满减券", "99-10-10"},
	}

	fmt.Printf("act1 name=%s coupon.name=%s coupon.rule=%s\n", act1.Name, act1.CouponTpl.Name, act1.CouponTpl.Rule)

	act2 := act1
	fmt.Printf("act2 name=%s coupon.name=%s coupon.rule=%s\n", act2.Name, act2.CouponTpl.Name, act2.CouponTpl.Rule)

	act2.Name = "新活动"
	act2.CouponTpl.Rule = "20-10-10"

	fmt.Printf("act1 name=%s coupon.name=%s coupon.rule=%s\n", act1.Name, act1.CouponTpl.Name, act1.CouponTpl.Rule)
	fmt.Printf("act2 name=%s coupon.name=%s coupon.rule=%s\n", act2.Name, act2.CouponTpl.Name, act2.CouponTpl.Rule)

	fmt.Println(act1.CouponTpl.Rule != act1.CouponTpl.Rule)

}

func DeepCopy() {
	act1 := Act{
		"大促活动",
		&CouponTpl{"满减券", "99-10-10"},
	}

	fmt.Printf("act1 name=%s coupon.name=%s coupon.rule=%s\n", act1.Name, act1.CouponTpl.Name, act1.CouponTpl.Rule)

	act2 := act1.DeepCopy()
	fmt.Printf("act2 name=%s coupon.name=%s coupon.rule=%s\n", act2.Name, act2.CouponTpl.Name, act2.CouponTpl.Rule)

	act2.Name = "新活动"
	act2.CouponTpl.Rule = "20-10-10"

	fmt.Printf("act1 name=%s coupon.name=%s coupon.rule=%s\n", act1.Name, act1.CouponTpl.Name, act1.CouponTpl.Rule)
	fmt.Printf("act2 name=%s coupon.name=%s coupon.rule=%s\n", act2.Name, act2.CouponTpl.Name, act2.CouponTpl.Rule)

	fmt.Println(act1.CouponTpl.Rule == act2.CouponTpl.Rule)

}
