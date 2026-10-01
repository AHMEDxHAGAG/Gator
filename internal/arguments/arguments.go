// Package arguments
package arguments

type Arguments struct {
	Name string
	Args []string
}

const (
	namePos = 0
	argPos  = 1
)

func NewArgument(args []string) Arguments {
	name := args[namePos]
	var arguments []string = nil
	if len(args) > argPos {
		arguments = args[argPos:]
	}
	return Arguments{
		Name: name,
		Args: arguments,
	}
}
