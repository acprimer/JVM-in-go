package main

import "fmt"
import (
	"ch03/classpath"
	"ch03/classfile"
	"strings"
)

func main() {
	cmd := parseCmd()
	if cmd.versionFlag {
		fmt.Println("version 0.0.3")
	} else if cmd.helpFlag || cmd.class == "" {
		printUsage()
	} else {
		startJVM(cmd)
	}
}

func startJVM(cmd *Cmd) {
	cp := classpath.Parse(cmd.XjreOption, cmd.cpOption)
	className := strings.Replace(cmd.class, ".", "/", -1)
	cf := loadClass(className, cp)
	fmt.Println(cmd.class)
	printClassInfo(cf)
}

func loadClass(className string, cp *classpath.Classpath) *classfile.ClassFile {
	classData, _, err := cp.ReadClass(className)
	if err != nil {
		panic(err)
	}
	cf, err := classfile.Parse(classData)
	if err != nil {
		panic(err)
	}
	return cf
}

func printClassInfo(cf *classfile.ClassFile) {
	// magic
	fmt.Printf("magic number: 0x%X\n", cf.MagicNumber())

	// version
	fmt.Printf("minjor version: %v\n", cf.MinorVersion())
	fmt.Printf("major version: %v\n", cf.MajorVersion())
	// flags
	flags := cf.AccessFlags()
	fmt.Printf("access flags: 0x%x %v\n", flags, classfile.ParseAccessFlag(flags))

	// constant pool
	fmt.Println("Constant pool:")
	for i := 1; i < len(cf.ConstantPool()); i++ {
		if cf.ConstantPool()[i] == nil {
			continue
		}
		fmt.Printf("%5s = %v\n", fmt.Sprintf("#%d", i), cf.ConstantPool()[i].GetConstantInfo())
	}

	fmt.Printf("this class: %v\n", cf.ClassName())
	fmt.Printf("super class: %v\n", cf.SuperClassName())
	fmt.Printf("interfaces: %v\n", cf.InterfaceNames())

	// fields
	fmt.Printf("fields count: %v\n", len(cf.Fields()))
	for _, f := range cf.Fields() {
		fmt.Printf("  %s %s ", classfile.ParseDescriptor(f.Descriptor()), f.Name())
		for _, attr := range f.Attributes() {
			switch f.Descriptor() {
			case "Z":
				constantAttr := attr.(*classfile.ConstantValueAttribute)
				constantInfo := cf.ConstantPool()[constantAttr.ConstantValueIndex()].(*classfile.ConstantIntegerInfo)
				fmt.Printf("%s: int %v\n", attr.PrintInfo(), constantInfo.Val())
				break
			default:
				fmt.Println()
			}
		}
	}

	// methods
	fmt.Printf("methods count: %v\n", len(cf.Methods()))
	for _, f := range cf.Methods() {
		fmt.Printf("  %s %d\n", f.Name(), len(f.Attributes()))
		for _, attr := range f.Attributes() {
			fmt.Printf("%s\n", attr.PrintInfo())
		}
	}

	// attributes
	fmt.Printf("attribute count: %v\n", len(cf.Attributes()))
	for _, f := range cf.Attributes() {
		fmt.Printf("%s\n", f.PrintInfo())
	}
}
