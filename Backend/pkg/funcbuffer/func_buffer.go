package funcbuffer

type FuncBuffer struct {
	buffer []func()
}

func New() *FuncBuffer {
	return &FuncBuffer{
		buffer: make([]func(), 0),
	}
}

func (fb *FuncBuffer) Add(f func()) {
	fb.buffer = append(fb.buffer, f)
}

func (fb *FuncBuffer) Execute() {
	for _, f := range fb.buffer {
		f()
	}

	fb.buffer = nil
}
