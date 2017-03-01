package classfile

import "fmt"

type ConstantMemberrefInfo struct {
	cp               ConstantPool
	classIndex       uint16
	nameAndTypeIndex uint16
}

func (self *ConstantMemberrefInfo) readInfo(reader *ClassReader) {
	self.classIndex = reader.readUint16()
	self.nameAndTypeIndex = reader.readUint16()
}

func (self *ConstantMemberrefInfo) ClassName() string {
	return self.cp.getClassName(self.classIndex)
}

func (self *ConstantMemberrefInfo) NameAndDescriptor() (string, string) {
	return self.cp.getNameAndType(self.nameAndTypeIndex)
}

type ConstantFieldrefInfo struct {
	ConstantMemberrefInfo
}

func (self *ConstantFieldrefInfo) GetConstantInfo() string {
	name, descriptor := self.NameAndDescriptor()
	lineNumber := fmt.Sprintf("#%d.#%d", self.classIndex, self.nameAndTypeIndex)
	return fmt.Sprintf("%-19s%-15s// %v:%v", "Fieldref", lineNumber, self.ClassName() + "." + name, descriptor)
}

type ConstantMethodrefInfo struct {
	ConstantMemberrefInfo
}

func (self *ConstantMethodrefInfo) GetConstantInfo() string {
	name, descriptor := self.NameAndDescriptor()
	lineNumber := fmt.Sprintf("#%d.#%d", self.classIndex, self.nameAndTypeIndex)
	return fmt.Sprintf("%-19s%-15s// %v:%v", "Methodref", lineNumber, self.ClassName() + "." + name, descriptor)
}

type ConstantInterfaceMethodrefInfo struct {
	ConstantMemberrefInfo
}

func (self *ConstantInterfaceMethodrefInfo) GetConstantInfo() string {
	name, descriptor := self.NameAndDescriptor()
	lineNumber := fmt.Sprintf("#%d.#%d", self.classIndex, self.nameAndTypeIndex)
	return fmt.Sprintf("%-19s%-15s// %v:%v", "Interface", lineNumber, self.ClassName() + name, self.ClassName() + descriptor)
}