package rtda

import "ch06/classfile"

type MemberRef struct {
	SysRef
	name string
	desriptor string
}

func (self *MemberRef) copyMemberRefInfo(refInfo *classfile.ConstantMemberrefInfo) {
	self.className = refInfo.ClassName()
	self.name, self.desriptor = refInfo.NameAndDescriptor()
}
