package service

import (
	"net"
)

type Service struct {
	*net.MX
}

func New() *Service {
	return &Service{}
}

func (domain string) hasMX() bool {
	MXRecords, err := net.LookupMX(domain)
}
