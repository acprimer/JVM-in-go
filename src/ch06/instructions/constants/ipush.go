package constants

import (
	"ch06/instructions/base"
	"ch06/rtda"
)

type BIPUSH struct {
	val int8
}

func (self *BIPUSH) FetchOperands(reader *base.BytecodeReader)  {
	self.val = reader.ReadInt8()
}

func (self *BIPUSH) Execute(frame *rtda.Frame) {
	n := int32(self.val)
	frame.OperandStack().PushInt(n)
}

type SIPUSH struct {
	val int16
}

func (self *SIPUSH) FetchOperands(reader *base.BytecodeReader)  {
	self.val = reader.ReadInt16()
}

func (self *SIPUSH) Execute(frame *rtda.Frame) {
	n := int32(self.val)
	frame.OperandStack().PushInt(n)
}