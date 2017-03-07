package control

import (
	"ch06/instructions/base"
	"ch06/rtda"
)

type RETURN struct {
	base.NoOperandsInstruction
}

func (self *RETURN) Execute(frame *rtda.Frame) {
	frame.Thread().PopFrame()
}

type IRETURN struct {
	base.NoOperandsInstruction
}

func (self *IRETURN) Execute(frame *rtda.Frame) {
	thread := frame.Thread()
	currentFrame := thread.PopFrame()
	invokerFrame := thread.CurrentFrame()
	retVal := currentFrame.OperandStack().PopInt()
	invokerFrame.OperandStack().PushInt(retVal)
}

type DRETURN struct {
	base.NoOperandsInstruction
}

func (self *DRETURN) Execute(frame *rtda.Frame) {
	thread := frame.Thread()
	currentFrame := thread.PopFrame()
	invokerFrame := thread.CurrentFrame()
	retVal := currentFrame.OperandStack().PopDouble()
	invokerFrame.OperandStack().PushDouble(retVal)
}

type FRETURN struct {
	base.NoOperandsInstruction
}

func (self *FRETURN) Execute(frame *rtda.Frame) {
	thread := frame.Thread()
	currentFrame := thread.PopFrame()
	invokerFrame := thread.CurrentFrame()
	retVal := currentFrame.OperandStack().PopFloat()
	invokerFrame.OperandStack().PushFloat(retVal)
}

type LRETURN struct {
	base.NoOperandsInstruction
}

func (self *LRETURN) Execute(frame *rtda.Frame) {
	thread := frame.Thread()
	currentFrame := thread.PopFrame()
	invokerFrame := thread.CurrentFrame()
	retVal := currentFrame.OperandStack().PopLong()
	invokerFrame.OperandStack().PushLong(retVal)
}

type ARETURN struct {
	base.NoOperandsInstruction
}

func (self *ARETURN) Execute(frame *rtda.Frame) {
	thread := frame.Thread()
	currentFrame := thread.PopFrame()
	invokerFrame := thread.CurrentFrame()
	retVal := currentFrame.OperandStack().PopRef()
	invokerFrame.OperandStack().PushRef(retVal)
}
