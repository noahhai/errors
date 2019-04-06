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

func (e *err) Add(msg string) *err {
	n := err{msg: msg}
	n.child = e
	return &n
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
