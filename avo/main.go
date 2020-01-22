package main

import (
	. "github.com/mmcloughlin/avo/build"
	. "github.com/mmcloughlin/avo/operand"
)

type ctx struct {
	rot16 Mem
	rot8  Mem
	iv    Mem
}

func main() {
	// iv := GLOBL("iv", RODATA|NOPTR)
	// for n, v := range []U32{
	// 	0x6A09E667, 0xBB67AE85, 0x3C6EF372, 0xA54FF53A,
	// 	0x510E527F, 0x9B05688C, 0x1F83D9AB, 0x5BE0CD19,
	// } {
	// 	DATA(4*n, v)
	// }

	rot16 := GLOBL("rot16_shuf", RODATA|NOPTR)
	for n, v := range []U8{
		0x02, 0x03, 0x00, 0x01, 0x06, 0x07, 0x04, 0x05,
		0x0A, 0x0B, 0x08, 0x09, 0x0E, 0x0F, 0x0C, 0x0D,
		0x12, 0x13, 0x10, 0x11, 0x16, 0x17, 0x14, 0x15,
		0x1A, 0x1B, 0x18, 0x19, 0x1E, 0x1F, 0x1C, 0x1D,
	} {
		DATA(n, v)
	}

	rot8 := GLOBL("rot8_shuf", RODATA|NOPTR)
	for n, v := range []U8{
		0x01, 0x02, 0x03, 0x00, 0x05, 0x06, 0x07, 0x04,
		0x09, 0x0A, 0x0B, 0x08, 0x0D, 0x0E, 0x0F, 0x0C,
		0x11, 0x12, 0x13, 0x10, 0x15, 0x16, 0x17, 0x14,
		0x19, 0x1A, 0x1B, 0x18, 0x1D, 0x1E, 0x1F, 0x1C,
	} {
		DATA(n, v)
	}

	c := ctx{
		rot16: rot16,
		rot8:  rot8,
		// iv:    iv,
	}

	// AVX(c)
	Hash8(c)

	Generate()
}