package references

import (
	"ch06/instructions/base"
	"ch06/rtda"
	"ch06/rtda/heap"
)

type INVOKE_INTERFACE struct {
	index uint
	count uint8
	zero  uint8
}

func (self *INVOKE_INTERFACE) FetchOperands(reader *base.BytecodeReader) {
	self.index = uint(reader.ReadInt16())
	self.count = reader.ReadUint8()
	self.zero = reader.ReadUint8()
}

func (self *INVOKE_INTERFACE) Execute(frame *rtda.Frame) {
	currentClass := frame.Method().Class()
	methodRef := currentClass.ConstantPool().GetConstant(self.index).(*heap.InterfaceMethodRef)
	resolvedMethod := methodRef.ResolvedInterfaceMethod()
	if resolvedMethod.IsFlagSet(heap.ACC_STATIC) ||
		resolvedMethod.IsFlagSet(heap.ACC_PRIVATE) {
		panic("java.lang.IncompatibleClassChangeError")
	}

	ref := frame.OperandStack().GetRefFromTop(resolvedMethod.ArgSlotCount() - 1)
	if ref == nil {
		panic("java.lang.NullPointerExecption")
	}

	methodToBeInvoked := heap.LookupMethodInClass(ref.Class(),
		methodRef.Name(), methodRef.Descriptor())

	if methodToBeInvoked == nil || methodToBeInvoked.IsFlagSet(heap.ACC_ABSTRACT) {
		panic("java.lang.AbstractMethodError")
	}

	if !methodToBeInvoked.IsFlagSet(heap.ACC_PUBLIC) {
		panic("java.lang.AbstractMethodError")
	}

	base.InvokeMethod(frame, methodToBeInvoked)
}