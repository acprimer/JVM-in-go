package heap

import (
	"ch06/classfile"
	"strings"
)

type Class struct {
	AccessFlags
	name              string
	superClassName    string
	interfaceNames    []string
	constantPool      *ConstantPool
	fields            []*Field
	methods           []*Method
	loader            *ClassLoader
	superClass        *Class
	interfaces        []*Class
	instanceSlotCount uint
	staticSlotCount   uint
	staticVars        Slots
}

func newClass(cf *classfile.ClassFile) *Class {
	class := &Class{}
	class.accessFlags = cf.AccessFlags()
	class.name = cf.ClassName()
	class.superClassName = cf.SuperClassName()
	class.interfaceNames = cf.InterfaceNames()
	class.constantPool = newConstantPool(class, cf.ConstantPool())
	class.fields = newFields(class, cf.Fields())
	class.methods = newMethods(class, cf.Methods())
	return class
}

func (self *Class) ConstantPool() *ConstantPool {
	return self.constantPool
}

func (self *Class) StaticVars() Slots {
	return self.staticVars
}

func (self *Class) SuperClass() *Class {
	return self.superClass
}

func (self *Class) isAccessibleTo(other *Class) bool {
	return self.IsFlagSet(ACC_PUBLIC) || self.GetPackageName() == other.GetPackageName()
}

func (self *Class) GetPackageName() string {
	if i := strings.LastIndex(self.name, "/"); i >= 0 {
		return self.name[:i]
	}
	return ""
}

func (self *Class) NewObject() *Object {
	return newObject(self)
}

func (self *Class) GetMainMethod() *Method {
	return self.getMainMethod("main", "([Ljava/lang/String;)V")
}

func (self *Class) getMainMethod(name, descriptor string) *Method {
	for _, method := range self.methods {
		if method.IsFlagSet(ACC_STATIC) &&
			method.name == name &&
			method.descriptor == descriptor {
			return method
		}
	}
	return nil
}
