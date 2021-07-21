package domain

type EntityNotExists struct {
	msg string
}

func (err *EntityNotExists) Error() string {
	return err.msg
}

func NewEntityNotExists(msg string) *EntityNotExists {
	return &EntityNotExists{msg: msg}
}

type InvalidArgument struct {
	msg string
}

func NewInvalidArgument(msq string) *InvalidArgument {
	return &InvalidArgument{msg: msq}
}

func (err *InvalidArgument) Error() string {
	return err.msg
}
