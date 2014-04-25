package main

import (
	"./go.hid"
	"fmt"
	"time"
)

const vendor_id  = 0x534c
const product_id = 0x0001

func enumerate() []string {
	list, err := hid.Enumerate(vendor_id, product_id)
	var result []string
	if err == nil {
		for _, v := range list {
			if v.InterfaceNumber == 0 || v.InterfaceNumber == -1 {
				result = append(result, v.Path)
			}
		}
	}
	return result
}

func main() {
	for {

		m := enumerate()

		for k, v := range m {
			dev, err := hid.OpenPath(v)
			if err == nil {
				ms, _ := dev.ManufacturerString()
				ps, _ := dev.ProductString()
				sns, _ := dev.SerialNumberString()
				fmt.Println(k, ms, ps, sns)
			}
			dev.Close()
		}

		time.Sleep(500 * time.Millisecond)
	}
}
