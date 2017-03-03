package classfile

import "fmt"

type ConstantInvokeDynamicInfo struct {
	cp *ConstantPool
	bootstrapMethodAttrIndex uint16
	nameAndTypeIndex uint16
}

func (self *ConstantInvokeDynamicInfo) readInfo(reader *ClassReader) {
	self.bootstrapMethodAttrIndex = reader.readUint16()
	self.nameAndTypeIndex = reader.readUint16()
}

func (self *ConstantInvokeDynamicInfo) NameAndType() (string, string) {
	return self.cp.getNameAndType(self.nameAndTypeIndex)
}

func (self *ConstantInvokeDynamicInfo) GetConstantInfo() string {
	return fmt.Sprintf("Fieldref\t\t#%v", self.nameAndTypeIndex)
}

//func (self *ConstantInvokeDynamicInfo) BoostrapMethodInfo() (uint16, []uint16) {
//	bmAttr := self.cp.cf.
//}

type ConstantMethodHandleInfo struct {
	referenceKind uint8
	referenceIndex uint16
}

func (self *ConstantMethodHandleInfo) readInfo(reader *ClassReader) {
	self.referenceKind = reader.readUint8()
	self.referenceIndex = reader.readUint16()
}

func (self *ConstantMethodHandleInfo) GetConstantInfo() string {
	return fmt.Sprintf("Fieldref\t\t#%v", self.referenceIndex)
}

func (self *ConstantMethodHandleInfo) ReferenceKind() uint8 {
	return self.referenceKind
}
func (self *ConstantMethodHandleInfo) ReferenceIndex() uint16 {
	return self.referenceIndex
}

type ConstantMethodTypeInfo struct {
	descriptorIndex uint16
}

func (self *ConstantMethodTypeInfo) readInfo(reader *ClassReader) {
	self.descriptorIndex = reader.readUint16()
}

func (self *ConstantMethodTypeInfo) GetConstantInfo() string {
	return fmt.Sprintf("MethodType\t\t#%v", self.descriptorIndex)
}