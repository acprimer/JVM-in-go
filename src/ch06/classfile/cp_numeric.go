package classfile

import (
	"math"
	"fmt"
)

// integer
type ConstantIntegerInfo struct {
	val int32
}

func (self *ConstantIntegerInfo) readInfo(reader *ClassReader) {
	bytes := reader.readUint32()
	self.val = int32(bytes)
}

func (self *ConstantIntegerInfo) Value() int32 {
	return self.val
}

func (self *ConstantIntegerInfo) GetConstantInfo() string {
	return fmt.Sprintf("%-19s%v", "Integer", self.val)
}

// float
type ConstantFloatInfo struct {
	val float32
}

func (self *ConstantFloatInfo) readInfo(reader *ClassReader) {
	bytes := reader.readUint32()
	self.val = math.Float32frombits(bytes)
}

func (self *ConstantFloatInfo) Value() float32 {
	return self.val
}

func (self *ConstantFloatInfo) GetConstantInfo() string {
	return fmt.Sprintf("%-19s%vf", "Float", self.val)
}

// long
type ConstantLongInfo struct {
	val int64
}

func (self *ConstantLongInfo) readInfo(reader *ClassReader) {
	bytes := reader.readUint64()
	self.val = int64(bytes)
}

func (self *ConstantLongInfo) Value() int64 {
	return self.val
}

func (self *ConstantLongInfo) GetConstantInfo() string {
	return fmt.Sprintf("%-19s%vl", "Long", self.val)
}

// double
type ConstantDoubleInfo struct {
	val float64
}

func (self *ConstantDoubleInfo) readInfo(reader *ClassReader) {
	bytes := reader.readUint64()
	self.val = math.Float64frombits(bytes)
}

func (self *ConstantDoubleInfo) Value() float64 {
	return self.val
}

func (self *ConstantDoubleInfo) GetConstantInfo() string {
	return fmt.Sprintf("%-19s%vd", "Double", self.val)
}