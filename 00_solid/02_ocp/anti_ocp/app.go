package anti_ocp

import "fmt"

type Cp int

const (
	Fliggy Cp = iota
	CTrip
)

type Scene int

const (
	SearchScene Scene = iota
	ActScene
)

type Coupon struct {
	name  string
	cp    Cp
	scene Scene
}

type Filter struct {
}

func (f Filter) FilterByCp(coupons []Coupon, cp Cp) []*Coupon {
	ret := make([]*Coupon, 0)
	for i, v := range coupons {
		if v.cp == cp {
			ret = append(ret, &coupons[i])
		}
	}
	return ret
}

func (f Filter) FilterByScene(coupons []Coupon, scene Scene) []*Coupon {
	ret := make([]*Coupon, 0)
	for i, v := range coupons {
		if v.scene == scene {
			ret = append(ret, &coupons[i])
		}
	}
	return ret
}

func (f Filter) FilterByCpAndScene(coupons []Coupon, cp Cp, scene Scene) []*Coupon {
	ret := make([]*Coupon, 0)
	for i, v := range coupons {
		if v.cp == cp && v.scene == scene {
			ret = append(ret, &coupons[i])
		}
	}
	return ret
}

func client() {
	f1 := Coupon{"满减券", Fliggy, SearchScene}
	f2 := Coupon{"折扣券", Fliggy, ActScene}
	c := Coupon{"满减券", CTrip, SearchScene}

	coupons := []Coupon{f1, f2, c}
	fmt.Println("Fliggy Coupons: ")
	f := Filter{}
	for _, v := range f.FilterByCp(coupons, Fliggy) {
		fmt.Printf("%s is Fliggy\n", v.name)
	}
}
