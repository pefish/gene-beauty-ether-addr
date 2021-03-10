package util

import (
	"fmt"
	"github.com/pefish/go-jsvm/pkg/vm"
	"testing"
)

func TestName(t *testing.T) {
	jsVm, err := vm.NewVmAndLoadWithFile("/Users/pefish/Work/golang/gene-beauty-ether-addr/test/test.js")
	if err != nil {
		return
	}
	fmt.Println(jsVm.Run([]interface{}{"0xaaaaaaaaaa46yhrth"}))
	fmt.Println(jsVm.Run([]interface{}{"0xpefish46yhrth"}))
	fmt.Println(jsVm.Run([]interface{}{"0x012345aaa46yhrth"}))
	fmt.Println(jsVm.Run([]interface{}{"0x543210aaaa46yhrth"}))
	fmt.Println(jsVm.Run([]interface{}{"0x123456aaaa46yhrth"}))
	fmt.Println(jsVm.Run([]interface{}{"0x654321aaa46yhrth"}))
}
