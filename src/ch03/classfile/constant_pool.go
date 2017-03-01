package classfile

type ConstantPool []ConstantInfo

func readConstantPool(reader *ClassReader) ConstantPool {
	cpCount := int(reader.readUint16())
	cp := make([]ConstantInfo, cpCount)
	for i := 1; i < cpCount; i++ {
		cp[i] = readConstantInfo(reader, cp)
		switch cp[i].(type) {
		case *ConstantLongInfo, *ConstantDoubleInfo:
			i++
		}
	}
	return cp
}

func (self ConstantPool) getConstantInfo(index uint16) ConstantInfo {
	if cpInfo := self[index]; cpInfo != nil {
		return cpInfo
	}
	panic("Invalid constant pool index!")
}

func (self ConstantPool) getNameAndType(index uint16) (string, string) {
	ntInfo := self.getConstantInfo(index).(*ConstantNameAndTypeInfo)
	name := self.getUtf8(ntInfo.nameIndex)
	_type := self.getUtf8(ntInfo.descriptorIndex)
	return name, _type
}

func (self ConstantPool) getClassName(index uint16) string {
	classInfo := self.getConstantInfo(index).(*ConstantClassInfo)
	return self.getUtf8(classInfo.nameIndex)
}

func (self ConstantPool) getUtf8(index uint16) string {
	utf8Info := self.getConstantInfo(index).(*ConstantUtf8Info)
	return utf8Info.str
}

func (self ConstantPool) getValue(index uint16) string {
	utf8Info := self.getConstantInfo(index).(*ConstantUtf8Info)
	return utf8Info.str
}

func (self ConstantPool) getIntegerValue(index uint16) int32 {
	intInfo := self.getConstantInfo(index).(*ConstantIntegerInfo)
	return intInfo.val
}

func (self ConstantPool) getLongValue(index uint16) int64 {
	longInfo := self.getConstantInfo(index).(*ConstantLongInfo)
	return longInfo.val
}

func (self ConstantPool) getFloatValue(index uint16) float32 {
	floatInfo := self.getConstantInfo(index).(*ConstantFloatInfo)
	return floatInfo.val
}

func (self ConstantPool) getDoubleValue(index uint16) float64 {
	doubleInfo := self.getConstantInfo(index).(*ConstantDoubleInfo)
	return doubleInfo.val
}