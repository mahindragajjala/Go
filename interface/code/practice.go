package code

type datad interface {
	Addfunc() int
	Minus() int
}

func Practice_interface(d datad) {
	d.Addfunc()
}
func Addfunc() float32 {
	return 32.00
}
func Minus() float32 {
	return 32.00
}
