package references

import (
	"ch06/instructions/base"
	"ch06/rtda"
	"ch06/rtda/heap"
)

type NEW struct {
	base.Index16Instruction
}

func (self *NEW) Execute(frame *rtda.Frame) {
	cp := frame.Method().Class().ConstantPool()
	classRef := cp.GetConstant(self.Index).(*heap.ClassRef)
	class := classRef.ResolvedClass()
	if class.IsFlagSet(heap.ACC_INTERFACE) || class.IsFlagSet(heap.ACC_ABSTRACT) {
		panic("java.lang.InstantiationError")
	}
	ref := class.NewObject()
	frame.OperandStack().PushRef(ref)
}