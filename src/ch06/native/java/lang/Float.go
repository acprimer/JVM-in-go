package lang

import (
	"ch06/native"
	"ch06/rtda"
	"math"
)

const jlFloat = "java/lang/Float"

func init() {
	native.Register(jlFloat, "floatToRawIntBits",
		"(F)I",
		floatToRawIntBits)
}

func floatToRawIntBits(frame *rtda.Frame) {
	value := frame.LocalVars().GetFloat(0)
	bits := math.Float32bits(value) // todo
	frame.OperandStack().PushInt(int32(bits))
}