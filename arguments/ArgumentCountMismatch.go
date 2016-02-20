package arguments

/*
ArgumentCountMismatch describes an error where the number of schema elements
do not match the number of passed in arguments
*/
type ArgumentCountMismatch struct {
}

func ArgumentCountMismatchException() ArgumentCountMismatch {
	return ArgumentCountMismatch{}
}

func (exception ArgumentCountMismatch) Error() string {
	return "The number of arguments does not match the number of elements in the schema"
}
