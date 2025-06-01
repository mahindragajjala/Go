package questions

/*
Write a program where a struct `Dog` implements an interface `Animal`
with `Speak()` method.
*/
type Animal interface {
	Speak()
}
type Dog struct {
}

func (d Dog) Speak() {
}
func Implements(a Animal) {
	a.Speak()
}

func Main_interface() {
	data := Dog{}
	Implements(data)
}
