package conversions

import (
	"ch05/instructions/base"
	"ch05/rtda"
)

type F2D struct{ base.NoOperandsInstruction }
type F2I struct{ base.NoOperandsInstruction }
type F2L struct{ base.NoOperandsInstruction }

func (self *F2D) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	x := stack.PopFloat()
	y := float64(x)
	stack.PushDouble(y)
}

func (self *F2I) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	x := stack.PopFloat()
	y := int32(x)
	stack.PushInt(y)
}

func (self *F2L) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	x := stack.PopFloat()
	y := int64(x)
	stack.PushLong(y)
}