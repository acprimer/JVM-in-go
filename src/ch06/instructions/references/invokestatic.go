package references

import (
	"ch06/instructions/base"
	"ch06/rtda"
	"ch06/rtda/heap"
)

type INVOKE_STATIC struct {
	base.Index16Instruction
}

func (self *INVOKE_STATIC) Execute(frame *rtda.Frame) {
	cp := frame.Method().Class().ConstantPool()
	methodRef := cp.GetConstant(self.Index).(*heap.MethodRef)
	resolveMethod := methodRef.ResolvedMethod()
	if !resolveMethod.IsFlagSet(heap.ACC_STATIC) {
		panic("java.lang.IncompatibleClassChangError")
	}
	base.InvokeMethod(frame, resolveMethod)
}
