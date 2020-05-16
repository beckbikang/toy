package main

import "fmt"

type ctx struct {
	name string
	idx  int8
	h    handlers
}

/*
*
idx=0 func1 start
	idx = 1 func2 start
		idx = 2 func3 start
		idx = 3 return
	idx =1 func2 end
idx=0 func1 end
*
*/
func (c *ctx) next() {
	c.idx++
	fmt.Printf("idx:%d\n", c.idx)
	for c.idx < int8(len(c.h)) {
		c.h[c.idx](c)
		c.idx++
	}
}

type handler func(*ctx)

type handlers []handler

type engin struct {
	name string
	hds  handlers
}

func (e *engin) run(c *ctx) {
	fmt.Printf("engine:%s running\n", e.name)
	c.h = e.hds
	c.next()
	fmt.Printf("engine:%s end\n", e.name)
}

func main() {
	testHandler()
}

func testHandler() {

	h1 := func(c *ctx) {
		fmt.Println("h1 start")
		c.next()
		fmt.Println("h1 end")
	}

	h2 := func(c *ctx) {
		fmt.Println("h2 start")
		c.next()
		fmt.Println("h2 end")
	}

	h3 := func(c *ctx) {
		fmt.Println("h3 running...")
	}

	hds := make([]handler, 3)
	hds[0] = h1
	hds[1] = h2
	hds[2] = h3

	c1 := &ctx{
		idx:  -1,
		name: "ctx1",
	}

	engin1 := &engin{
		name: "engin1",
		hds:  hds,
	}
	engin1.run(c1)

}
