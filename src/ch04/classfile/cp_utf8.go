package classfile

import "fmt"

type ConstantUtf8Info struct {
	str string
}

func (self *ConstantUtf8Info) readInfo(reader *ClassReader) {
	length := uint32(reader.readUint16())
	bytes := reader.readBytes(length)
	self.str = decodeMUTF8(bytes)
}

func decodeMUTF8(bytes []byte) string {
	return string(bytes)
}

func (self *ConstantUtf8Info) GetConstantInfo() string {
	return fmt.Sprintf("%-19s%v", "Utf8", self.str)
}