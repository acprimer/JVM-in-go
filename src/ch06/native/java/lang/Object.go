package lang

import (
	"ch06/native"
	"ch06/rtda"
)

func init() {
	native.Register("java/lang/Object",
		"getClass",
		"()Ljava/lang/Class;", getClass)
}

func getClass(frame *rtda.Frame) {
	this := frame.LocalVars().GetRef(0)
	class := this.Class().JClass()
	frame.OperandStack().PushRef(class)
}