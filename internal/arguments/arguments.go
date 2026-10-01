// Package arguments
package arguments

type Arguments struct {
	Name string
	Args []string
}

func NewArgument(args []string) Arguments {
	name := args[0]
	var arguments []string = nil
	if len(args) > 1 {
		arguments = args[1:]
	}
	return Arguments{
		Name: name,
		Args: arguments,
	}
}
