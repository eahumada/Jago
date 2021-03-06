package jago

import "fmt"

/*178 (0XB2)*/
func GETSTATIC(opcode uint8, f *Frame, t *Thread, c *Class, m *Method) {
	index := f.index16()
	field := c.constantPool[index].(*FieldRef).ResolvedField()
	f.push(field.class.staticVars[field.index])
}

/*179 (0XB3)*/
func PUTSTATIC(opcode uint8, f *Frame, t *Thread, c *Class, m *Method) {
	index := f.index16()
	value := f.pop()
	fieldref := c.constantPool[index].(*FieldRef)
	class := fieldref.ResolvedClass()
	field := fieldref.ResolvedField()
	class.staticVars[field.index] = value
}

/*180 (0XB4)*/
func GETFIELD(opcode uint8, f *Frame, t *Thread, c *Class, m *Method) {
	index := f.index16()
	objectref := f.pop().(ObjectRef)

	f.push(f.getField(objectref, index))
}

/*181 (0XB5)*/
func PUTFIELD(opcode uint8, f *Frame, t *Thread, c *Class, m *Method) {
	index := f.index16()
	value := f.pop()
	objectref := f.pop().(ObjectRef)

	f.putField(objectref, index, value)
}

/*182 (0XB6)*/
func INVOKEVIRTUAL(opcode uint8, f *Frame, t *Thread, c *Class, m *Method) {
	index := f.index16()
	method := c.constantPool[index].(*MethodRef).ResolvedMethod()

	if method.isStatic() {
		Fatal("Not an instance method")
	}
	parameterCount := len(method.parameterDescriptors) + 1 // with an extra objectref: this
	params := make([]Value, parameterCount)
	for i := parameterCount-1; i >= 0; i-- {
		params[i] = f.pop()
	}
	// get objectref and target method
	objectref := params[0].(ObjectRef)
	if objectref.IsNull() {
		Fatal("NullPointerException")
	}
	overridenMethod := objectref.class.FindMethod(method.name, method.descriptor)
	if method.isNative() {
		result := t.invokeNativeMethod(overridenMethod, params...)
		if overridenMethod.returnDescriptor != JVM_SIGNATURE_VOID {
			f.push(result)
		}
	} else {
		frame := NewStackFrame(overridenMethod)
		// pass parameters
		for j := 0; j < parameterCount; j++ {
			frame.storeVar(uint(j), params[j])
		}

		t.pushFrame(frame)
	}
}

// like invokevirtual with objectref, but don't find along the inheritance
/*183 (0XB7)*/
func INVOKESPECIAL(opcode uint8, f *Frame, t *Thread, c *Class, m *Method) {
	index := f.index16()

	method := c.constantPool[index].(*MethodRef).ResolvedMethod()
	parameterCount := len(method.parameterDescriptors) + 1 // with an extra objectref: this
	params := make([]Value, parameterCount)
	for i := parameterCount-1; i >= 0; i-- {
		params[i] = f.pop()
	}

	if method.isNative() {
		result := t.invokeNativeMethod(method, params...)
		if method.returnDescriptor != JVM_SIGNATURE_VOID {
			f.push(result)
		}
	} else {
		frame := NewStackFrame(method)
		// pass parameters
		for j := 0; j < parameterCount; j++ {
			frame.storeVar(uint(j), params[j])
		}

		t.pushFrame(frame)
	}
}

/*184 (0XB8)*/
func INVOKESTATIC(opcode uint8, f *Frame, t *Thread, c *Class, m *Method) {
	index := f.index16()

	methodref := c.constantPool[index].(*MethodRef)
	class := methodref.ResolvedClass()
	// do class initialization if any
	clinits := class.Initialize()
	for _, clinit := range clinits { t.pushFrame(NewStackFrame(clinit))}
	if len(clinits) > 0 {
		f.pc = 88888888 // revert this instruction so as to execute again
		return
	}

	method := methodref.ResolvedMethod()
	parameterCount := len(method.parameterDescriptors)
	params := make([]Value, parameterCount)
	// read parameters
	for i := parameterCount - 1; i >= 0; i-- {
		params[i] = f.pop()
	}

	if method.isNative() {
		result := t.invokeNativeMethod(method, params...)
		if method.returnDescriptor != JVM_SIGNATURE_VOID {
			f.push(result)
		}
	} else {
		frame := NewStackFrame(method)
		// pass parameters
		for j := parameterCount - 1; j >= 0; j-- {
			frame.storeVar(uint(j), params[j])
		}

		t.pushFrame(frame)
	}
}

/*185 (0XB9)*/
func INVOKEINTERFACE(opcode uint8, f *Frame, t *Thread, c *Class, m *Method) {
	panic(fmt.Sprintf("Not implemented for opcode %d\n", opcode))
}

/*186 (0XBA)*/
func INVOKEDYNAMIC(opcode uint8, f *Frame, t *Thread, c *Class, m *Method) {
	panic(fmt.Sprintf("Not implemented for opcode %d\n", opcode))
}

/*187 (0XBB)*/
func NEW(opcode uint8, f *Frame, t *Thread, c *Class, m *Method) {
	index := f.index16()
	class := c.constantPool[index].(*ClassRef).ResolvedClass().(*Class)
	objectref := class.NewObject()
	f.push(objectref)
}

// array component type
const (
	T_BOOLEAN	    = 4
	T_CHAR	        = 5
	T_FLOAT	        = 6
	T_DOUBLE	    = 7
	T_BYTE	        = 8
	T_SHORT	        = 9
	T_INT	        = 10
	T_LONG	        = 11
)

/*188 (0XBC)*/
func NEWARRAY(opcode uint8, f *Frame, t *Thread, c *Class, m *Method) {
	atype := uint8(m.code[f.pc+1])
	var componentDescriptor string
	switch atype {
	case T_CHAR: componentDescriptor = JVM_SIGNATURE_CHAR
	case T_BYTE: componentDescriptor = JVM_SIGNATURE_BYTE
	case T_SHORT: componentDescriptor = JVM_SIGNATURE_SHORT
	case T_INT: componentDescriptor = JVM_SIGNATURE_INT
	case T_LONG: componentDescriptor = JVM_SIGNATURE_LONG
	case T_FLOAT: componentDescriptor = JVM_SIGNATURE_FLOAT
	case T_DOUBLE: componentDescriptor = JVM_SIGNATURE_DOUBLE
	case T_BOOLEAN: componentDescriptor = JVM_SIGNATURE_BOOLEAN
	default:
		Fatal("Invalid atype value")
	}
	count := f.pop().(Int)
	arrayClass := BOOTSTRAP_CLASSLOADER.CreateClass(JVM_SIGNATURE_ARRAY + componentDescriptor).(*ArrayClass)
	arrayref := arrayClass.NewArray(count)
	f.push(arrayref)
}

/*189 (0XBD)*/
func ANEWARRAY(opcode uint8, f *Frame, t *Thread, c *Class, m *Method) {
	index := f.index16()
	count := f.pop().(Int)

	var arrayClass *ArrayClass
	componentType := c.constantPool[index].(*ClassRef).ResolvedClass()
	switch componentType.(type) {
	case *Class:
		arrayClass = BOOTSTRAP_CLASSLOADER.CreateClass(JVM_SIGNATURE_ARRAY + JVM_SIGNATURE_CLASS + componentType.Name() + JVM_SIGNATURE_ENDCLASS).(*ArrayClass)
	case *ArrayClass:
		arrayClass = BOOTSTRAP_CLASSLOADER.CreateClass(JVM_SIGNATURE_ARRAY + componentType.Name()).(*ArrayClass)
	default:
		Fatal("Not a class or array class")
	}
	arrayref := arrayClass.NewArray(count)

	f.push(arrayref)
}

/*190 (0XBE)*/
func ARRAYLENGTH(opcode uint8, f *Frame, t *Thread, c *Class, m *Method) {
	f.push(f.pop().(ArrayRef).length)
}

/*191 (0XBF)*/
func ATHROW(opcode uint8, f *Frame, t *Thread, c *Class, m *Method) {
	panic(fmt.Sprintf("Not implemented for opcode %d\n", opcode))
}

/*192 (0XC0)*/
func CHECKCAST(opcode uint8, f *Frame, t *Thread, c *Class, m *Method) {
	panic(fmt.Sprintf("Not implemented for opcode %d\n", opcode))
}

/*193 (0XC1)*/
func INSTANCEOF(opcode uint8, f *Frame, t *Thread, c *Class, m *Method) {
	panic(fmt.Sprintf("Not implemented for opcode %d\n", opcode))
}

/*194 (0XC2)*/
func MONITORENTER(opcode uint8, f *Frame, t *Thread, c *Class, m *Method) {
	panic(fmt.Sprintf("Not implemented for opcode %d\n", opcode))
}

/*195 (0XC3)*/
func MONITOREXIT(opcode uint8, f *Frame, t *Thread, c *Class, m *Method) {
	panic(fmt.Sprintf("Not implemented for opcode %d\n", opcode))
}
