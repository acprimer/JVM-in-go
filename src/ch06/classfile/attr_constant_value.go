package classfile

import "fmt"

type ConstantValueAttribute struct {
	cp ConstantPool
	constantValueIndex uint16
	attrLen uint32
	descriptorIndex uint16
}

func (self *ConstantValueAttribute) readInfo(reader *ClassReader) {
	self.constantValueIndex = reader.readUint16()
}

func (self *ConstantValueAttribute) ConstantValueIndex() uint16 {
	return self.constantValueIndex
}

func (self *ConstantValueAttribute) PrintInfo() string {
	return fmt.Sprintf("ConstantValue")
}
