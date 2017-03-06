package conversions

import (
	"ch06/instructions/base"
	"ch06/rtda"
)

type L2D struct{ base.NoOperandsInstruction }
type L2F struct{ base.NoOperandsInstruction }
type L2I struct{ base.NoOperandsInstruction }

func (self *L2D) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	x := stack.PopDouble()
	y := float64(x)
	stack.PushDouble(y)
}

func (self *L2F) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	x := stack.PopDouble()
	y := float32(x)
	stack.PushFloat(y)
}

func (self *L2I) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	x := stack.PopDouble()
	y := int32(x)
	stack.PushInt(y)
}