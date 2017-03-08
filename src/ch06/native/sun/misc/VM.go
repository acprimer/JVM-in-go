package misc

import (
	"ch06/native"
	"ch06/rtda"
	"ch06/instructions/base"
)

func init() {
	native.Register("sun/misc/VM",
		"initialize",
		"()V", initialize)
}

func initialize(frame *rtda.Frame) {
	classLoader := frame.Method().Class().Loader()
	jlSysClass := classLoader.LoadClass("java/lang/System")
	initSysClass := jlSysClass.GetStaticMethod("initializeSystemClass", "()V")
	base.InvokeMethod(frame, initSysClass)
}