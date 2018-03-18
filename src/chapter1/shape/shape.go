package shape

type Shape struct {
	Name string
}

type React struct {
	Width int
	Height int
}

type Circle struct {
	Raidus float32
}

func (r React) Area() int {
	return r.Width * r.Height
}

func (c Circle) Area() float32 {
	return 2*3.14*c.Raidus
}