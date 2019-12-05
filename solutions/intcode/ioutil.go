package intcode

type Ioutil struct {
	Stdout []int
	Stdin []int
}

func (io *Ioutil) Read() int {
	x := io.Stdin[0]
	io.Stdin = io.Stdin[1:]
	return x
}

func (io *Ioutil) Write(i int) {
	io.Stdout = append(io.Stdout, i)
}
