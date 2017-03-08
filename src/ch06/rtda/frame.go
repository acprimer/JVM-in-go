package rtda

import (
	"ch06/rtda/heap"
)

type Frame struct {
	lower        *Frame
	method       *heap.Method
	localVars    LocalVars
	operandStack *OperandStack
	thread       *Thread
	nextPC       int
}

func newFrame(thread *Thread, method *heap.Method) *Frame {
	return &Frame{
		thread:       thread,
		method:       method,
		localVars:    newLocalVars(method.MaxLocals()),
		operandStack: newOperandStack(method.MaxStack()),
	}
}

func (self *Frame) LocalVars() LocalVars {
	return self.localVars
}

func (self *Frame) OperandStack() *OperandStack {
	return self.operandStack
}

func (self *Frame) Thread() *Thread {
	return self.thread
}

func (self *Frame) Method() *heap.Method {
	return self.method
}

func (self *Frame) NextPC() int {
	return self.nextPC
}

func (self *Frame) SetNextPC(pc int) {
	self.nextPC = pc
}

func Branch(frame *Frame, offset int) {
	pc := frame.Thread().PC()
	nextPC := pc + offset
	frame.SetNextPC(nextPC)
}

func (self *Frame) RevertNextPC() {
	self.nextPC = self.thread.pc
}

func (self *Frame) Print() {
	self.localVars.Print()
	self.operandStack.Print()
}