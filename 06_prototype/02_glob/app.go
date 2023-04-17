package prototype

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

type CouponTpl struct {
	Name string
	Rule string
}

type Act struct {
	Name      string
	CouponTpl *CouponTpl
}

func (a *Act) DeepCopy() *Act {
	b := bytes.Buffer{}
	e := gob.NewEncoder(&b)
	_ = e.Encode(a)

	d := gob.NewDecoder(&b)
	newAct := Act{}
	_ = d.Decode(&newAct)
	return &newAct
}

func NewAct(name, rule string, act *Act) *Act {
	newAct := act.DeepCopy()
	newAct.Name = name
	newAct.CouponTpl.Rule = rule
	return newAct
}

func App() {
	act1 := Act{
		"大促活动",
		&CouponTpl{"满减券", "99-10-10"},
	}
	fmt.Printf("act1 name=%s coupon.name=%s coupon.rule=%s\n", act1.Name, act1.CouponTpl.Name, act1.CouponTpl.Rule)

	act2 := NewAct("新活动", "20-10-10", &act1)
	fmt.Printf("act1 name=%s coupon.name=%s coupon.rule=%s\n", act1.Name, act1.CouponTpl.Name, act1.CouponTpl.Rule)
	fmt.Printf("act2 name=%s coupon.name=%s coupon.rule=%s\n", act2.Name, act2.CouponTpl.Name, act2.CouponTpl.Rule)

	fmt.Println(act1.CouponTpl.Rule == act2.CouponTpl.Rule)

}
