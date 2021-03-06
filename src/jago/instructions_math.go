package jago

import "math"

/*96 (0X60)*/
func IADD(opcode uint8, f *Frame, t *Thread, c *Class, m *Method) {
	value2 := f.pop().(Int)
	value1 := f.pop().(Int)
	f.push(value1 + value2)
}

/*97 (0X61)*/
func LADD(opcode uint8, f *Frame, t *Thread, c *Class, m *Method) {
	value2 := f.pop().(Long)
	value1 := f.pop().(Long)
	f.push(value1 + value2)
}

/*98 (0X62)*/
func FADD(opcode uint8, f *Frame, t *Thread, c *Class, m *Method) {
	value2 := f.pop().(Float)
	value1 := f.pop().(Float)
	f.push(value1 + value2)
}

/*99 (0X63)*/
func DADD(opcode uint8, f *Frame, t *Thread, c *Class, m *Method) {
	value2 := f.pop().(Double)
	value1 := f.pop().(Double)
	f.push(value1 + value2)
}

/*100 (0X64)*/
func ISUB(opcode uint8, f *Frame, t *Thread, c *Class, m *Method) {
	value2 := f.pop().(Int)
	value1 := f.pop().(Int)
	f.push(value1 - value2)
}

/*101 (0X65)*/
func LSUB(opcode uint8, f *Frame, t *Thread, c *Class, m *Method) {
	value2 := f.pop().(Long)
	value1 := f.pop().(Long)
	f.push(Long(value1 - value2))
}

/*102 (0X66)*/
func FSUB(opcode uint8, f *Frame, t *Thread, c *Class, m *Method) {
	value2 := f.pop().(Float)
	value1 := f.pop().(Float)
	f.push(Float(value1 - value2))
}

/*103 (0X67)*/
func DSUB(opcode uint8, f *Frame, t *Thread, c *Class, m *Method) {
	value2 := f.pop().(Double)
	value1 := f.pop().(Double)
	f.push(Double(value1 - value2))
}

/*104 (0X68)*/
func IMUL(opcode uint8, f *Frame, t *Thread, c *Class, m *Method) {
	value2 := f.pop().(Int)
	value1 := f.pop().(Int)
	f.push(Int(value1 * value2))
}

/*105 (0X69)*/
func LMUL(opcode uint8, f *Frame, t *Thread, c *Class, m *Method) {
	value2 := f.pop().(Long)
	value1 := f.pop().(Long)
	f.push(Long(value1 * value2))
}

/*106 (0X6A)*/
func FMUL(opcode uint8, f *Frame, t *Thread, c *Class, m *Method) {
	value2 := f.pop().(Float)
	value1 := f.pop().(Float)
	f.push(Float(value1 * value2))
}

/*107 (0X6B)*/
func DMUL(opcode uint8, f *Frame, t *Thread, c *Class, m *Method) {
	value2 := f.pop().(Double)
	value1 := f.pop().(Double)
	f.push(Double(value1 * value2))
}

/*108 (0X6C)*/
func IDIV(opcode uint8, f *Frame, t *Thread, c *Class, m *Method) {
	value2 := f.pop().(Int)
	value1 := f.pop().(Int)
	f.push(Int(value1 / value2))
}

/*109 (0X6D)*/
func LDIV(opcode uint8, f *Frame, t *Thread, c *Class, m *Method) {
	value2 := f.pop().(Long)
	value1 := f.pop().(Long)
	f.push(Long(value1 / value2))
}

/*110 (0X6E)*/
func FDIV(opcode uint8, f *Frame, t *Thread, c *Class, m *Method) {
	value2 := f.pop().(Float)
	value1 := f.pop().(Float)
	f.push(Float(value1 / value2))
}

/*111 (0X6F)*/
func DDIV(opcode uint8, f *Frame, t *Thread, c *Class, m *Method) {
	value2 := f.pop().(Double)
	value1 := f.pop().(Double)
	f.push(Double(value1 / value2))
}

/*112 (0X70)*/
func IREM(opcode uint8, f *Frame, t *Thread, c *Class, m *Method) {
	value2 := f.pop().(Int)
	value1 := f.pop().(Int)
	f.push(Int(value1 % value2))
}

/*113 (0X71)*/
func LREM(opcode uint8, f *Frame, t *Thread, c *Class, m *Method) {
	value2 := f.pop().(Long)
	value1 := f.pop().(Long)
	f.push(Long(value1 % value2))
}

/*114 (0X72)*/
func FREM(opcode uint8, f *Frame, t *Thread, c *Class, m *Method) {
	value2 := f.pop().(Float)
	value1 := f.pop().(Float)
	f.push(Float(math.Remainder(float64(value1), float64(value2))))
}

/*115 (0X73)*/
func DREM(opcode uint8, f *Frame, t *Thread, c *Class, m *Method) {
	value2 := f.pop().(Double)
	value1 := f.pop().(Double)
	f.push(Float(math.Remainder(float64(value1), float64(value2))))
}

/*116 (0X74)*/
func INEG(opcode uint8, f *Frame, t *Thread, c *Class, m *Method) {
	value := f.pop().(Int)
	f.push(Int(-value))
}

/*117 (0X75)*/
func LNEG(opcode uint8, f *Frame, t *Thread, c *Class, m *Method) {
	value := f.pop().(Long)
	f.push(Long(-value))
}

/*118 (0X76)*/
func FNEG(opcode uint8, f *Frame, t *Thread, c *Class, m *Method) {
	value := f.pop().(Float)
	f.push(Float(-value))
}

/*119 (0X77)*/
func DNEG(opcode uint8, f *Frame, t *Thread, c *Class, m *Method) {
	value := f.pop().(Double)
	f.push(Double(-value))
}

/*120 (0X78)*/
func ISHL(opcode uint8, f *Frame, t *Thread, c *Class, m *Method) {
	value2 := f.pop().(Int)
	value1 := f.pop().(Int)
	f.push(Int(value1 << uint(value2 & 0x1F))) // low 5 bits
}

/*121 (0X79)*/
func LSHL(opcode uint8, f *Frame, t *Thread, c *Class, m *Method) {
	value2 := f.pop().(Int)

	value1 := f.pop().(Long)
	f.push(Long(value1 << uint(value2 & 0x3F))) // low 6 bits
}

/*122 (0X7A)*/
func ISHR(opcode uint8, f *Frame, t *Thread, c *Class, m *Method) {
	value2 := f.pop().(Int)
	value1 := f.pop().(Int)
	f.push(Int(value1 >> uint(value2 & 0x1F))) // low 5 bits
}

/*123 (0X7B)*/
func LSHR(opcode uint8, f *Frame, t *Thread, c *Class, m *Method) {
	value2 := f.pop().(Int)

	value1 := f.pop().(Long)
	f.push(Long(value1 >> uint(value2 & 0x3F))) // low 6 bits
}

/*124 (0X7C)*/
func IUSHR(opcode uint8, f *Frame, t *Thread, c *Class, m *Method) {
	value2 := f.pop().(Int)
	value1 := f.pop().(Int)
	f.push(Int(uint32(value1) >> uint(value2 & 0x1F))) // low 5 bits
}

/*125 (0X7D)*/
func LUSHR(opcode uint8, f *Frame, t *Thread, c *Class, m *Method) {
	value2 := f.pop().(Int)

	value1 := f.pop().(Long)
	f.push(Long(uint64(value1) >> uint(value2 & 0x3F))) // low 6 bits
}

/*126 (0X7E)*/
func IAND(opcode uint8, f *Frame, t *Thread, c *Class, m *Method) {
	value2 := f.pop().(Int)
	value1 := f.pop().(Int)
	f.push(Int(value1 & value2))
}

/*127 (0X7F)*/
func LAND(opcode uint8, f *Frame, t *Thread, c *Class, m *Method) {
	value2 := f.pop().(Long)
	value1 := f.pop().(Long)
	f.push(Long(value1 & value2))
}

/*128 (0X80)*/
func IOR(opcode uint8, f *Frame, t *Thread, c *Class, m *Method) {
	value2 := f.pop().(Int)
	value1 := f.pop().(Int)
	f.push(Int(value1 | value2))
}

/*129 (0X81)*/
func LOR(opcode uint8, f *Frame, t *Thread, c *Class, m *Method) {
	value2 := f.pop().(Long)
	value1 := f.pop().(Long)
	f.push(Long(value1 | value2))
}

/*130 (0X82)*/
func IXOR(opcode uint8, f *Frame, t *Thread, c *Class, m *Method) {
	value2 := f.pop().(Int)
	value1 := f.pop().(Int)
	f.push(Int(value1 ^ value2))
}

/*131 (0X83)*/
func LXOR(opcode uint8, f *Frame, t *Thread, c *Class, m *Method) {
	value2 := f.pop().(Long)
	value1 := f.pop().(Long)
	f.push(Long(value1 ^ value2))
}

/*132 (0X84)*/
func IINC(opcode uint8, f *Frame, t *Thread, c *Class, m *Method) {
	index := f.index8()
	const_value := int8(m.code[f.pc+2])
	Trace("sfd: %d", const_value)
	value := f.loadVar(uint(index)).(Int)
	f.storeVar(uint(index), value + Int(const_value))
}
