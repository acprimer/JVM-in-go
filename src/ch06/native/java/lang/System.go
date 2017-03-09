package lang

import (
	"ch06/native"
	"ch06/rtda"
	"ch06/rtda/heap"
)

const jlSystem = "java/lang/System"

func init() {
	native.Register(jlSystem,
		"arraycopy",
		"(Ljava/lang/Object;ILjava/lang/Object;II)V", arraycopy)
	native.Register(jlSystem, "setIn0", "(Ljava/io/InputStream;)V", setIn0)
	native.Register(jlSystem, "initProperties", "(Ljava/util/Properties;)Ljava/util/Properties;", initProperties)
}

func arraycopy(frame *rtda.Frame) {
	vars := frame.LocalVars()
	src := vars.GetRef(0)
	srcPos := vars.GetInt(1)
	dest := vars.GetRef(2)
	desPos := vars.GetInt(3)
	length := vars.GetInt(4)

	if src == nil || dest == nil {
		panic("java.lang.NullPointerException")
	}
	if !checkArrayCopy(src, dest) {
		panic("java.lang.ArrayStoreException")
	}
	if srcPos < 0 || desPos < 0 || length < 0 ||
		srcPos+length > src.ArrayLength() ||
		desPos+length > dest.ArrayLength() {
		panic("java.lang.IndexOutOfBoundException")
	}

	heap.ArrayCopy(src, dest, srcPos, desPos, length)
}

func checkArrayCopy(src, dest *heap.Object) bool {
	srcClass := src.Class()
	destClass := dest.Class()
	if !srcClass.IsArray() || !destClass.IsArray() {
		return false
	}
	if srcClass.ComponentClass().IsPrimitive() ||
		destClass.ComponentClass().IsPrimitive() {
		return srcClass == destClass
	}
	return true
}

func setIn0(frame *rtda.Frame) {
	vars := frame.LocalVars()
	in := vars.GetRef(0)

	sysClass := frame.Method().Class()
	sysClass.SetRefVar("in", "Ljava/io/InputStream;", in)
}

func initProperties(frame *rtda.Frame) {

}