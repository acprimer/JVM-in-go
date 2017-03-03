package classfile

type DeprecatedAttribute struct {
	MakerAttribute
}

type SyntheticAttribute struct {
	MakerAttribute
}

type MakerAttribute struct {}

func (self *MakerAttribute) readInfo(reader *ClassReader) {
	// read nothing
}

func (self *MakerAttribute) PrintInfo() string {
	return "Deprecated"
}