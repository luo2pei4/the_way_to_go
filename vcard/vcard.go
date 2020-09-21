package main

import (
	"fmt"
	"time"
)

// Address 地址结构体
type Address struct {
	contury  string
	province string
	city     string
	street   string
	num      string
}

// VCard visa卡结构体
type VCard struct {
	name     string
	birthday time.Time
	address  *Address
}

func main() {

	addr := &Address{
		contury:  "China",
		province: "SiChuan",
		city:     "ChengDu",
		street:   "XiLin",
		num:      "199",
	}

	vc := new(VCard)
	vc.name = "Luo Pei"
	vc.birthday, _ = time.Parse("2006-01-02 15:04:05", "1983-11-09 00:00:00")
	vc.address = addr

	fmt.Println(addr)
	fmt.Println(vc)
}
