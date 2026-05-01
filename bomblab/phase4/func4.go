package main

import "fmt"

func main() {
	for i := range int32(15) {
		fmt.Printf("func4(%d, 0, 14) = %d\n", i, func4(i, 0, 14))
	}
}

func func4(a, b, c int32) int32 {
	v1 := c // %eax
	v1 -= b
	v2 := v1 // %ecx
	v2 = int32(uint32(v2) >> 0x1f)
	v1 += v2
	v1 >>= 1
	v2 = v1 + b
	if v2 <= a {
		goto L1
	}
	c = v2 - 1
	v1 = func4(a, b, c)
	v1 += v1
	goto L2
L1:
	v1 = 0
	if v2 >= a {
		goto L2
	}
	b = v2 + 1
	v1 = func4(a, b, c)
	v1 = 1 + v1 + v1
L2:
	return v1
}
