package classfile

import "strings"

const (
	ACC_PUBLIC     = 0x0001
	ACC_FINAL      = 0x0010
	ACC_SUPER      = 0x0020
	ACC_INTERFACE  = 0x0200
	ACC_ABSTRACT   = 0x0400
	ACC_SYNTHETIC  = 0x1000
	ACC_ANNOTATION = 0x2000
	ACC_ENUM       = 0x4000
)

func ParseAccessFlag(flags uint16) string {
	flagToken := ""
	if (flags & ACC_PUBLIC) != 0 {
		flagToken += "ACC_PUBLIC, "
	}
	if (flags & ACC_FINAL) != 0 {
		flagToken += "ACC_FINAL, "
	}
	if (flags & ACC_SUPER) != 0 {
		flagToken += "ACC_SUPER, "
	}
	if (flags & ACC_INTERFACE) != 0 {
		flagToken += "ACC_INTERFACE, "
	}
	if (flags & ACC_ABSTRACT) != 0 {
		flagToken += "ACC_ABSTRACT, "
	}
	if (flags & ACC_SYNTHETIC) != 0 {
		flagToken += "ACC_SYNTHETIC, "
	}
	if len(flagToken) == 0 {
		return flagToken
	}
	return strings.TrimSuffix(flagToken, ", ")
}
