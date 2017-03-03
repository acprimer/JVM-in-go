package constants

import (
	"ch05/instructions/base"
	"ch05/rtda"
)

type NOP struct {
	base.NoOperandsInstruction
}

func (self *NOP) Execute(frame *rtda.Frame) {
	// do nothing
}
