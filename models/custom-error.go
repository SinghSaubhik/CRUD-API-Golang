package models

type Error struct {
	error string
}

func (e Error) Error() string {
	return e.error
}

func (e Error) NewError(err string) Error {
	e.error = err
	return e
}
