package intcode

type IoMgr interface {
	Read() int64
	Write(int64)
}

type PreparedIO struct {
	Stdout []int64
	Stdin []int64
}

func (io *PreparedIO) Read() int64 {
	x := io.Stdin[0]
	io.Stdin = io.Stdin[1:]
	return x
}

func (io *PreparedIO) Write(i int64) {
	io.Stdout = append(io.Stdout, i)
}

type ChanIO struct {
	Stdout chan int64
	Stdin chan int64
}

func GenChanIO(buf int64) *ChanIO {
	return &ChanIO{
		Stdout: make(chan int64, buf),
		Stdin: make(chan int64, buf),
	}
}

func (io *ChanIO) Read() int64 {
	return <-io.Stdin
}

func (io *ChanIO) Write(i int64) {
	io.Stdout <- i
}