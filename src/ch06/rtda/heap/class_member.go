package heap

import "ch06/classfile"

type ClassMember struct {
	AccessFlags
	name       string
	descriptor string
	class      *Class
}

func (self *ClassMember) copyMemberInfo(memberInfo *classfile.MemberInfo) {
	self.accessFlags = memberInfo.AccessFlags()
	self.name = memberInfo.Name()
	self.descriptor = memberInfo.Descriptor()
}

func (self *ClassMember) Descriptor() string {
	return self.descriptor
}

func (self *ClassMember) Class() *Class {
	return self.class
}

func (self *ClassMember) Name() string {
	return self.name
}

func (self *ClassMember) isAccessibleTo(d *Class) bool {
	if self.IsFlagSet(ACC_PUBLIC) {
		return true
	}
	c := self.class
	if self.IsFlagSet(ACC_PROTECTED) {
		return d == c || d.IsSubClassOf(c) ||
			c.GetPackageName() == d.GetPackageName()
	}
	if !self.IsFlagSet(ACC_PRIVATE) {
		return c.GetPackageName() == d.GetPackageName()
	}
	return d == c
}
