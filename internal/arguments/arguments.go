// Package arguments
package arguments

type Arguments struct {
	Name string
	Args []string
}

func NewArgument(args []string) Arguments {
	return Arguments{
		Name: args[0],
		Args: args[1:],
	}
}
