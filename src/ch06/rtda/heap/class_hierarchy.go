package heap

func (self *Class) IsAssignableFrom(other *Class) bool {
	s, t := other, self
	if s == t {
		return true
	}
	if !t.IsFlagSet(ACC_INTERFACE) {
		return s.IsSubClassOf(t)
	} else {
		return s.IsImplements(t)
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