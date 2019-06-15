package client

// PrintCapable indicates the object can print to stdout
type PrintCapable interface {
	Print(i ...interface{})
	Printf(format string, i ...interface{})
	Println(i ...interface{})
}

// PrintCapable indicates the object can print to stderr
type PrintErrCapable interface {
	PrintErr(i ...interface{})
	PrintErrf(format string, i ...interface{})
	PrintErrln(i ...interface{})
}
