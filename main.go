// =====================================================================
//
// udpmask.go -
//
// Created by skywind on 2024/11/18
// Last Modified: 2024/11/18 19:45:07
//
// =====================================================================
package main

import (
	"flag"

	"github.com/skywind3000/udpmask/service"
)

func main() {
	src := flag.String("src", "", "local address, eg: 0.0.0.0:8080")
	dst := flag.String("dst", "", "destination address, eg: 8.8.8.8:443")
	mask := flag.String("mask", "", "encryption/decryption key")
	mark := flag.Uint("mark", 0, "fwmark value")
	flag.Parse()
	if src == nil || dst == nil {
		flag.Usage()
		return
	}
	if *src == "" || *dst == "" {
		flag.Usage()
		return
	}
	config := service.ServiceConfig{
		SrcAddr: *src,
		DstAddr: *dst,
		Mask:    *mask,
		Mark:    uint32(*mark),
	}
	service.StartService(config)
}
