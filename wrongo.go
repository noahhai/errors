package wrongo

type formatFunc func(e *err) string

var defaultFormatFunc formatFunc = func(e *err) string {
	errString := e.msg
	curr := e
	for curr.child != nil {
		if curr.child.msg != "" {
			errString += "\n"
			errString += curr.child.msg
		}
		curr = curr.child
	}
	return errString
}

type err struct {
	child  *err
	msg    string
	format formatFunc
}

func New(msg string) *err {
	return &err{msg: msg}
}

func From(e error) *err {
	if asErr, ok := e.(*err); ok {
		return asErr
	}
	return New(e.Error())
}

func (e *err) Add(msg string) *err {
	n := err{msg: msg}
	n.child = e
	return &n
}

func Or(err1, err2 error) *err {
	if err1 != nil {
		return From(err1)
	} else if err2 != nil {
		return From(err2)
	}
	return nil
}

func (e *err) Or(err *err) *err {
	if e == nil {
		return err
	}
	return e
}

func (e *err) SetFormatter(f formatFunc) {
	e.format = f
}

func (e *err) Error() string {
	if e.format != nil {
		return e.format(e)
	}
	return defaultFormatFunc(e)
}
