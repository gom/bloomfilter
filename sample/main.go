package main

import (
	"fmt"

	"github.com/gom/bloomfilter"
)

func print(b bool) {
	fmt.Printf("%x\n", b)
}

func main() {
	bf := bloomfilter.New(10000, 3)
	s := "test string foobar hogehuga"

	bf.Clear()
	bf.Add(s)
	bf.Add("foo").Add("bar")
	print(bf.Check(s))
	print(bf.Check("other string"))
	print(bf.CheckAndAdd("hoge"))
	print(bf.Check("hoge"))
	bf.Clear()
	print(bf.Check(s))
	fmt.Printf("%x\n", bf.FalsePositiveRate(100))
}
