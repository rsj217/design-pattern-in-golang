package facade

import "fmt"

type Checker interface {
	check() bool
}

type Stock struct {
	actId int64
}

func (s *Stock) check() bool {
	fmt.Printf("[Stock module] check actId=%d left stock\n", s.actId)
	return true
}

type Auth struct {
	session string
}

func (u *Auth) check() bool {
	fmt.Printf("[Auth module] check session=%s login status\n", u.session)
	return true
}

type Risk struct {
	device string
}

func (r *Risk) check() bool {
	fmt.Printf("[Risk module] check device=%s risk status\n", r.device)
	return true
}

type Facade struct {
	stock *Stock
	auth  *Auth
	risk  *Risk
}

func (f *Facade) Check() bool {
	return f.auth.check() && f.stock.check() && f.risk.check()
}

func client() {
	actId := int64(101)
	session := "session217"
	device := "ios171"
	facade := &Facade{
		&Stock{actId},
		&Auth{session},
		&Risk{device},
	}

	if facade.Check() {
		fmt.Println("send coupon")
	}

}
