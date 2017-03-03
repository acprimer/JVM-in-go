package math

import (
	"ch05/instructions/base"
	"ch05/rtda"
	"ch05/instructions/stack"
)

type IINC struct{
	Index uint
	Const int32
}

func (self*IINC) FetchOperands(reader *base.BytecodeReader) {
	self.Index = uint(reader.ReadInt8())
	self.Const = int32(reader.ReadInt8())
}

func (self *IINC) Execute(frame *rtda.Frame) {
	localVars := frame.LocalVars()
	val := localVars.GetInt(self.Index)
	val += self.Const
	localVars.SetInt(self.Index, val)
}