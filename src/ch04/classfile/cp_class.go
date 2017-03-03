package classfile

import (
	"fmt"
)

type ConstantClassInfo struct {
	cp        ConstantPool
	nameIndex uint16
}

func (self *ConstantClassInfo) readInfo(reader *ClassReader) {
	self.nameIndex = reader.readUint16()
}

func (self *ConstantClassInfo) Name() string {
	return self.cp.getUtf8(self.nameIndex)
}

func (self *ConstantClassInfo) GetConstantInfo() string {
	lineNumber := fmt.Sprintf("#%d", self.nameIndex)
	return fmt.Sprintf("%-19s%-15v// %v", "Class", lineNumber, self.Name())
}