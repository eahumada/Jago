package jago

import "fmt"

/*202 (0XCA)*/
func BREAKPOINT(opcode uint8, f *Frame, t *Thread, c *Class, m *Method) {
	panic(fmt.Sprintf("Not implemented for opcode %d\n", opcode))
}

/*254 (0XFE)*/
func IMPDEP1(opcode uint8, f *Frame, t *Thread, c *Class, m *Method) {
	panic(fmt.Sprintf("Not implemented for opcode %d\n", opcode))
}

/*255 (0XFF)*/
func IMPDEP2(opcode uint8, f *Frame, t *Thread, c *Class, m *Method) {
	panic(fmt.Sprintf("Not implemented for opcode %d\n", opcode))
}
