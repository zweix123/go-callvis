package main

import (
	pkg "github.com/zweix123/go-callvis/examples/main/pkg"
)

func main() {
	funcs()
	var c calls
	c.execution()
	c.invocation()
}

func funcs() {
	pkg.Exported()
}

type calls struct{}

func (calls) execution() {
	pkg.Regular()
}

func (calls) invocation() {
	pkg.T.Static()
	var i pkg.Iface = pkg.T
	i.Dynamic()
}
