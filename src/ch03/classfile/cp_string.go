package classfile

import "fmt"

type ConstantStringInfo struct {
	cp          ConstantPool
	stringIndex uint16
}

func (self *ConstantStringInfo) readInfo(reader *ClassReader) {
	self.stringIndex = reader.readUint16()
}

func (self *ConstantStringInfo) String() string {
	return self.cp.getUtf8(self.stringIndex)
}

func (self *ConstantStringInfo) GetConstantInfo() string {
	lineNumber := fmt.Sprintf("#%d", self.stringIndex)
	return fmt.Sprintf("%-19s%-15v// %v", "String", lineNumber, self.String())
}