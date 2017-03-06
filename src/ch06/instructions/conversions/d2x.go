package conversions

import (
	"ch06/instructions/base"
	"ch06/rtda"
)

type D2F struct{ base.NoOperandsInstruction }
type D2I struct{ base.NoOperandsInstruction }
type D2L struct{ base.NoOperandsInstruction }

func (self *D2F) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	x := stack.PopDouble()
	y := float32(x)
	stack.PushFloat(y)
}

func (self *D2I) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	x := stack.PopDouble()
	y := int32(x)
	stack.PushInt(y)
}

func (self *D2L) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	x := stack.PopDouble()
	y := int64(x)
	stack.PushLong(y)
}