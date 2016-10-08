package main

import (
	"fmt"
)

// Build some graphs, and then print some code that can be used to generate the
// graphs in SageMath.
func main() {
	// Diamond Graph
	gDiamond := NewLowBlockTimeGraph()
	d1 := gDiamond.CreateNode(gDiamond.GenesisNode())
	d2 := gDiamond.CreateNode(gDiamond.GenesisNode())
	_ = gDiamond.CreateNode(d1, d2)
	fmt.Printf("\n# Diamond Graph\n%s", gDiamond.SageGen())

	// Pentagon Graph
	gPentagon := NewLowBlockTimeGraph()
	p1 := gPentagon.CreateNode(gPentagon.GenesisNode())
	p2 := gPentagon.CreateNode(gPentagon.GenesisNode())
	p3 := gPentagon.CreateNode(p1)
	_ = gPentagon.CreateNode(p2, p3)
	fmt.Printf("\n# Pentagon Graph\n%s", gPentagon.SageGen())

	// Double Diamond Graph
	gDDiamond := NewLowBlockTimeGraph()
	dd1 := gDDiamond.CreateNode(gDDiamond.GenesisNode())
	dd2 := gDDiamond.CreateNode(gDDiamond.GenesisNode())
	dd3 := gDDiamond.CreateNode(dd1, dd2)
	dd4 := gDDiamond.CreateNode(dd2)
	_ = gDDiamond.CreateNode(dd3, dd4)
	fmt.Printf("\n# Double Diamond Graph\n%s", gDDiamond.SageGen())

	// Impossibility Proof Graph
	ip := NewLowBlockTimeGraph()
	ip1 := ip.CreateNode(ip.GenesisNode())
	ip2 := ip.CreateNode(ip.GenesisNode())
	ip3 := ip.CreateNode(ip1)
	ip4 := ip.CreateNode(ip2)
	ip5 := ip.CreateNode(ip3)
	ip6 := ip.CreateNode(ip4, ip5)
	ip7 := ip.CreateNode(ip4)
	ip8 := ip.CreateNode(ip7)
	ip9 := ip.CreateNode(ip8)
	ip10 := ip.CreateNode(ip9)
	_ = ip.CreateNode(ip10, ip6)
	fmt.Printf("\n# Impossibility Proof Graph\n%s", ip.SageGen())

	// Abstain Graph
	a := NewLowBlockTimeGraph()
	a1 := a.CreateNode(a.GenesisNode())
	a2 := a.CreateNode(a1)
	a3 := a.CreateNode(a2)
	a4 := a.CreateNode(a3)
	a5 := a.CreateNode(a4)
	a6 := a.CreateNode(a.GenesisNode())
	a7 := a.CreateNode(a6)
	a8 := a.CreateNode(a7)
	a9 := a.CreateNode(a8)
	a10 := a.CreateNode(a9)
	a11 := a.CreateNode(a10)
	a12 := a.CreateNode(a11)
	a13 := a.CreateNode(a12)
	a14 := a.CreateNode(a13)
	a15 := a.CreateNode(a14)
	a16 := a.CreateNode(a15)
	a17 := a.CreateNode(a16)
	_ = a.CreateNode(a17, a5)
	fmt.Printf("\n# Abstain Graph\n%s", a.SageGen())

	// Leech Graph
	l := NewLowBlockTimeGraph()
	l1 := l.CreateNode(l.GenesisNode())
	l2 := l.CreateNode(l1)
	l3 := l.CreateNode(l2)
	l4 := l.CreateNode(l3)
	l5 := l.CreateNode(l4)
	l6 := l.CreateNode(l5)
	l7 := l.CreateNode(l6)
	l8 := l.CreateNode(l7)
	l9 := l.CreateNode(l8)
	la := l.CreateNode(l9)
	lb := l.CreateNode(la)
	lc := l.CreateNode(lb)
	ld := l.CreateNode(l.GenesisNode())
	le := l.CreateNode(ld, l2)
	lf := l.CreateNode(le, l3)
	lg := l.CreateNode(lf, l5)
	lh := l.CreateNode(lg, l7)
	li := l.CreateNode(lh, l8)
	lj := l.CreateNode(li, l9)
	lk := l.CreateNode(lj, lb)
	_ = l.CreateNode(lk, lc)
	fmt.Printf("\n# Leech Graph\n%s", l.SageGen())

	// Low Latency Adversary Graph
	lla := NewLowBlockTimeGraph()
	lla1 := lla.CreateNode(lla.GenesisNode())
	lla2 := lla.CreateNode(lla.GenesisNode())
	lla3 := lla.CreateNode(lla.GenesisNode())
	lla4 := lla.CreateNode(lla.GenesisNode())
	lla5 := lla.CreateNode(lla.GenesisNode())
	lla6 := lla.CreateNode(lla5)
	lla7 := lla.CreateNode(lla1, lla2, lla3, lla4, lla6)
	lla8 := lla.CreateNode(lla6)
	lla9 := lla.CreateNode(lla8)
	_ = lla.CreateNode(lla7, lla9)
	fmt.Printf("\n# Low Latenct Adversary Graph\n%s", lla.SageGen())
}
