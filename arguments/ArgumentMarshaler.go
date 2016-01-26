package arguments

type ArgumentMarshaler interface {
	Set(currentArgument string) ArgsException
}
