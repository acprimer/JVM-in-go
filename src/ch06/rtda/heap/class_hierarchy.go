package heap

func (self *Class) IsAssignableFrom(other *Class) bool {
	s, t := other, self
	if s == t {
		return true
	}
	if !s.IsArray() {
		if !s.IsFlagSet(ACC_INTERFACE) {
			if !t.IsFlagSet(ACC_INTERFACE) {
				return s.IsSubClassOf(t)
			} else {
				return s.IsImplements(t)
			}
		} else {
			if !t.IsFlagSet(ACC_INTERFACE) {
				return t.isJlObject()
			} else {
				return t.IsSuperInterfaceOf(s)
			}
		}
	} else {
		if !t.IsArray() {
			if !t.IsFlagSet(ACC_INTERFACE) {
				return t.isJlObject()
			} else {
				return t.isJlCloneable() || t.isJioSerializable()
			}
		} else {
			sc := s.ComponentClass()
			tc := t.ComponentClass()
			return sc == tc || tc.IsAssignableFrom(sc)
		}
	}
}

func (self *Class) IsSubClassOf(other *Class) bool {
	for c := self.superClass; c != nil; c = c.superClass {
		if c == other {
			return true
		}
	}
	return false
}

func (self *Class) IsImplements(iface *Class) bool {
	for c := self; c != nil; c = c.superClass {
		for _, i := range c.interfaces {
			if i == iface || i.IsSubInterfaceOf(iface) {
				return true
			}
		}
	}
	return false
}

func (self *Class) IsSubInterfaceOf(iface *Class) bool {
	for _, i := range self.interfaces {
		if i == iface || i.IsSubInterfaceOf(iface) {
			return true
		}
	}
	return false
}

func (self *Class) IsSuperClassOf(other *Class) bool {
	return other.IsSubClassOf(self)
}

func (self *Class) IsSuperInterfaceOf(other *Class) bool {
	return other.IsSubInterfaceOf(self)
}