package intcode

type IoMgr interface {
	Read() int
	Write(int)
}

type PreparedIO struct {
	Stdout []int
	Stdin []int
}

func (io *PreparedIO) Read() int {
	x := io.Stdin[0]
	io.Stdin = io.Stdin[1:]
	return x
}

func (io *PreparedIO) Write(i int) {
	io.Stdout = append(io.Stdout, i)
}

type ChanIO struct {
	Stdout chan int
	Stdin chan int
}

func GenChanIO(buf int) *ChanIO {
	return &ChanIO{
		Stdout: make(chan int, buf),
		Stdin: make(chan int, buf),
	}
}

func (io *ChanIO) Read() int {
	return <-io.Stdin
}

func (io *ChanIO) Write(i int) {
	io.Stdout <- i
}