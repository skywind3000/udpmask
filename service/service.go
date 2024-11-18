// =====================================================================
//
// service.go -
//
// Created by skywind on 2024/11/18
// Last Modified: 2024/11/18 19:49:04
//
// =====================================================================
package service

import (
	"log"
	"os"
	"os/signal"

	"github.com/skywind3000/udpmask/forward"
)

type ServiceConfig struct {
	SrcAddr string
	DstAddr string
	Mask    string
	Mark    uint32
}

func StartService(config ServiceConfig) int {
	server := forward.NewUdpForward()
	saddr := forward.AddressResolve(config.SrcAddr)
	daddr := forward.AddressResolve(config.DstAddr)

	logger := log.Default()
	logger.Printf("config: %v\n", config)

	if saddr == nil {
		logger.Printf("config: invalid src address: \"%s\"\n", config.SrcAddr)
		return 1
	}

	if daddr == nil {
		logger.Printf("config: invalid dst address: \"%s\"\n", config.DstAddr)
		return 1
	}

	server.SetLogger(logger)
	server.SetMark(config.Mark)

	sigch := make(chan os.Signal, 1)
	signal.Notify(sigch, os.Interrupt, os.Kill)

	err := server.Open(saddr, daddr, config.Mask)
	if err != nil {
		return 2
	}

	<-sigch
	server.Close()

	return 0
}
