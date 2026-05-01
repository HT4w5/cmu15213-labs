package main

import "fmt"

func main() {
	for i := uint64(0); i <= 0x3e9; i++ {
		fmt.Printf("fun7(0x6030f0, %#x) = %d\n", i, fun7(0x6030f0, i))
	}
}

type node struct {
	value uint64
	left  uint64
	right uint64
}

var tree = map[uint64]node{
	0x6030f0: {value: 0x24, left: 0x603110, right: 0x603130}, // n1 (36)
	0x603110: {value: 0x08, left: 0x603190, right: 0x603150}, // n21 (8)
	0x603130: {value: 0x32, left: 0x603170, right: 0x6031b0}, // n22 (50)
	0x603150: {value: 0x16, left: 0x603270, right: 0x603230}, // n32 (22)
	0x603170: {value: 0x2d, left: 0x6031d0, right: 0x603290}, // n33 (45)
	0x603190: {value: 0x06, left: 0x6031f0, right: 0x603250}, // n31 (6)
	0x6031b0: {value: 0x6b, left: 0x603210, right: 0x6032b0}, // n34 (107)
	0x6031d0: {value: 0x28, left: 0x0, right: 0x0},           // n45 (40)
	0x6031f0: {value: 0x01, left: 0x0, right: 0x0},           // n41 (1)
	0x603210: {value: 0x63, left: 0x0, right: 0x0},           // n47 (99)
	0x603230: {value: 0x23, left: 0x0, right: 0x0},           // n44 (35)
	0x603250: {value: 0x07, left: 0x0, right: 0x0},           // n42 (7)
	0x603270: {value: 0x14, left: 0x0, right: 0x0},           // n43 (20)
	0x603290: {value: 0x2f, left: 0x0, right: 0x0},           // n46 (47)
	0x6032b0: {value: 0x3e9, left: 0x0, right: 0x0},          // n48 (1001)
}

func fun7(rdi, rsi uint64) uint64 {
	var rax, rdx uint64

	if rdi == 0 {
		goto L1
	}

	rdx = tree[rdi].value
	if int32(rdx) <= int32(rsi) {
		goto L2
	}
	rdi = tree[rdi].left
	rax = fun7(rdi, rsi)
	rax += rax
	goto L3

L2:
	rax = 0
	if int32(rdx) == int32(rsi) {
		goto L3
	}
	rdi = tree[rdi].right
	rax = fun7(rdi, rsi)
	rax = 1 + rax + rax
	goto L3

L1:
	rax = 0xffffffff
L3:
	return rax
}
