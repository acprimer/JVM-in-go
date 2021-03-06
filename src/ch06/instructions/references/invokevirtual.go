package references

import (
	"ch06/instructions/base"
	"ch06/rtda"
	"ch06/rtda/heap"
	"fmt"
)

type INVOKE_VIRTUAL struct {
	base.Index16Instruction
}

func (self *INVOKE_VIRTUAL) Execute(frame *rtda.Frame) {
	currentClass := frame.Method().Class()
	methodRef := currentClass.ConstantPool().GetConstant(self.Index).(*heap.MethodRef)
	resolvedMethod := methodRef.ResolvedMethod()
	if resolvedMethod.IsFlagSet(heap.ACC_STATIC) {
		panic("java.lang.IncompatibleClassChangeError")
	}

	ref := frame.OperandStack().GetRefFromTop(resolvedMethod.ArgSlotCount() - 1)
	if ref == nil {
		// hack println
		if methodRef.Name() == "println" {
			_println(frame.OperandStack(), methodRef.Descriptor())
			return
		}
		panic("java.lang.NullPointerExecption")
	}

	if resolvedMethod.IsFlagSet(heap.ACC_PROTECTED) &&
		resolvedMethod.Class().IsSuperClassOf(currentClass) &&
		resolvedMethod.Class().GetPackageName() != currentClass.GetPackageName() &&
		ref.Class() != currentClass && !ref.Class().IsSubClassOf(currentClass) {
		panic("java.lang.IllegalAccessError")
	}

	methodToBeInvoked := heap.LookupMethodInClass(ref.Class(),
		methodRef.Name(), methodRef.Descriptor())

	if methodToBeInvoked == nil || methodToBeInvoked.IsFlagSet(heap.ACC_ABSTRACT) {
		panic("java.lang.AbstractMethodError")
	}

	base.InvokeMethod(frame, methodToBeInvoked)
}

func _println(stack *rtda.OperandStack, descriptor string) {
	switch descriptor {
	case "(Z)V":
		fmt.Printf("%v\n", stack.PopInt() != 0)
	case "(C)V":
		fmt.Printf("%c\n", stack.PopInt())
	case "(B)V":
		fmt.Printf("%v\n", stack.PopInt())
	case "(S)V":
		fmt.Printf("%v\n", stack.PopInt())
	case "(I)V":
		fmt.Printf("%v\n", stack.PopInt())
	case "(J)V":
		fmt.Printf("%v\n", stack.PopLong())
	case "(F)V":
		fmt.Printf("%v\n", stack.PopFloat())
	case "(D)V":
		fmt.Printf("%v\n", stack.PopDouble())
	case "(Ljava/lang/String;)V":
		jStr := stack.PopRef()
		goStr := heap.GoString(jStr)
		fmt.Println(goStr)
	default:
		panic("println: " + descriptor)
	}
	stack.PopRef()
}
