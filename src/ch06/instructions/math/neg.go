package math

import (
	"ch06/instructions/base"
	"ch06/rtda"
)

type DNEG struct{ base.NoOperandsInstruction }
type FNEG struct{ base.NoOperandsInstruction }
type INEG struct{ base.NoOperandsInstruction }
type LNEG struct{ base.NoOperandsInstruction }

func (self *DNEG) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v1 := stack.PopDouble()
	stack.PushDouble(-v1)
}

func (self *FNEG) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v1 := stack.PopFloat()
	stack.PushFloat(-v1)
}

func (self *INEG) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v1 := stack.PopInt()
	stack.PushInt(-v1)
}

func (self *LNEG) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v1 := stack.PopLong()
	stack.PushLong(-v1)
}