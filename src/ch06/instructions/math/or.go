package math

import (
	"ch06/instructions/base"
	"ch06/rtda"
)

type IOR struct{ base.NoOperandsInstruction }
type LOR struct{ base.NoOperandsInstruction }

func (self *IOR) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v1 := stack.PopInt()
	v2 := stack.PopInt()
	result := v1 | v2
	stack.PushInt(result)
}

func (self *LOR) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v1 := stack.PopLong()
	v2 := stack.PopLong()
	result := v1 | v2
	stack.PushLong(result)
}