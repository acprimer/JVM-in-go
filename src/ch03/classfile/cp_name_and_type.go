package classfile

import "fmt"

type ConstantNameAndTypeInfo struct {
	cp              ConstantPool
	nameIndex       uint16
	descriptorIndex uint16
}

func (self *ConstantNameAndTypeInfo) readInfo(reader *ClassReader) {
	self.nameIndex = reader.readUint16()
	self.descriptorIndex = reader.readUint16()
}

func (self *ConstantNameAndTypeInfo) NameAndType() (string, string) {
	return self.cp.getUtf8(self.nameIndex), self.cp.getUtf8(self.descriptorIndex)
}

func (self *ConstantNameAndTypeInfo) GetConstantInfo() string {
	lineNumber := fmt.Sprintf("#%d.#%d", self.nameIndex, self.descriptorIndex)
	name, desc := self.NameAndType()
	return fmt.Sprintf("%-19s%-15s// %v:%v", "NameAndType", lineNumber, name, desc)
}
