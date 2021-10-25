package iptables

type Table string

const (
	RawTABLE    Table = "raw"
	NatTable    Table = "nat"
	FilterTable Table = "filter"
	MangleTable Table = "mangle"
)

type Chain string

const (
	PostRoutingChain Chain = "POSTROUTING"
	PreRoutingChain  Chain = "PREROUTING"
	OutputChain      Chain = "OUTPUT"
	InputChain       Chain = "INPUT"
	ForwardChain     Chain = "FORWARD"
)
