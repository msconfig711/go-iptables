package iptables

type Table string

const (
	RAW    Table = "raw"
	NAT    Table = "nat"
	Filter Table = "filter"
	Mangle Table = "mangle"
)

type Chain string

const (
	PostRouting Chain = "POSTROUTING"
	PreRouting  Chain = "PREROUTING"
	Output      Chain = "OUTPUT"
	Input       Chain = "INPUT"
	Forward     Chain = "FORWARD"
)
