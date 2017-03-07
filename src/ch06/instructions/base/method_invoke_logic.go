package base

import (
	"ch06/rtda"
	"ch06/rtda/heap"
)

func InvokeMethod(frame *rtda.Frame, method *heap.Method) {
	thread := frame.Thread()
	newFrame := thread.NewFrame(method)
	thread.PushFrame(newFrame)
	argSlotSlot := int(method.ArgSlotCount())
	if argSlotSlot > 0 {
		for i := argSlotSlot - 1; i >= 0; i-- {
			slot := frame.OperandStack().PopSlot()
			newFrame.LocalVars().SetSlot(uint(i), slot)
		}
	}
}
