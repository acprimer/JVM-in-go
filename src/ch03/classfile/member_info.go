package classfile

type MemberInfo struct {
	cp ConstantPool
	accessFlags uint16
	nameIndex uint16
	descriptorIndex uint16
	attributes []AttributeInfo
}

func readMembers(reader *ClassReader, cp ConstantPool) []*MemberInfo {
	memberCount := reader.readUint16()
	members := make([]*MemberInfo, memberCount)
	for i := range members {
		members[i] = readMember(reader, cp)
	}
	return members
}

func readMember(reader *ClassReader, cp ConstantPool) *MemberInfo {
	return &MemberInfo{
		cp : cp,
		accessFlags: reader.readUint16(),
		nameIndex:reader.readUint16(),
		descriptorIndex:reader.readUint16(),
		attributes:readAttributes(reader, cp),
	}
}

func (self *MemberInfo) Name() string {
	return self.cp.getUtf8(self.nameIndex)
}

func (self *MemberInfo) Descriptor() string {
	return self.cp.getUtf8(self.descriptorIndex)
}

func (self *MemberInfo) Attributes() []AttributeInfo {
	return self.attributes
}

func ParseDescriptor(desc string) string {
	switch desc {
	case "B": return "byte"
	case "C": return "char"
	case "D": return "double"
	case "F": return "float"
	case "I": return "int"
	case "J": return "long"
	case "L": return "reference"
	case "S": return "short"
	case "Z": return "boolean"
	case "[": return "reference"
	default:
		return "invalid descriptor"
	}
}