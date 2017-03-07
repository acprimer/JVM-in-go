package main

import "fmt"
import "os"

type Cmd struct {
	helpFlag bool
	versionFlag bool
	cpOption string
	XjreOption string
	class string
	args []string
}

func parseCmd() *Cmd {
	cmd := &Cmd{}
	cmd.XjreOption = "/Library/Java/JavaVirtualMachines/jdk1.7.0_79.jdk/Contents/Home/jre"
	//cmd.XjreOption = "C:\\Program Files\\Java\\jre1.8.0_101"
	cmd.cpOption = "/Users/yaodh/JavaProjects/DataStructure/out/production/DataStructure"
	cmd.class = "TestJava"
	//flag.Usage = printUsage
	//flag.BoolVar(&cmd.helpFlag, "help	", false, "print help message")
	//flag.BoolVar(&cmd.helpFlag, "?", false, "print help message")
	//flag.BoolVar(&cmd.versionFlag, "version", false, "print version and exit")
	//flag.StringVar(&cmd.cpOption, "classpath", "", "classpath")
	//flag.StringVar(&cmd.cpOption, "cp", "", "classpath")
	//flag.StringVar(&cmd.XjreOption, "Xjre", "", "path to jre")
	//flag.Parse()
	//args := flag.Args()
	//if len(args) > 0 {
	//	cmd.class = args[0]
	//	cmd.args = args[1:]
	//}
	return cmd
}

func printUsage() {
	fmt.Printf("Usage : %s [-options] class [args...]\n", os.Args[0])
}