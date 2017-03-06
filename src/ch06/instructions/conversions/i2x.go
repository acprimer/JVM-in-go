package conversions

import (
	"ch06/instructions/base"
	"ch06/rtda"
)

type I2B struct{ base.NoOperandsInstruction }
type I2C struct{ base.NoOperandsInstruction }
type I2S struct{ base.NoOperandsInstruction }
type I2L struct{ base.NoOperandsInstruction }
type I2F struct{ base.NoOperandsInstruction }
type I2D struct{ base.NoOperandsInstruction }

func (self *I2B) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	x := stack.PopInt()
	y := int32(int8(x))
	stack.PushInt(y)
}

func (self *I2C) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	x := stack.PopInt()
	y := int32(uint16(x))
	stack.PushInt(y)
}

func (self *I2S) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	x := stack.PopFloat()
	y := int32(int16(x))
	stack.PushInt(y)
}

func (self *I2L) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	x := stack.PopInt()
	y := int64(x)
	stack.PushLong(y)
}

func (self *I2F) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	x := stack.PopInt()
	y := float32(x)
	stack.PushFloat(y)
}

func (self *I2D) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	x := stack.PopFloat()
	y := float64(x)
	stack.PushDouble(y)
}