package codego

/*
Polymorphism means “many forms.”
In programming, it allows one interface or function to behave differently based on the object or input it is working with.
*/
/*
Real-Life Example of Polymorphism
Example: "Remote Control"
A single remote can control a TV, AC, or Sound System.
The button names are the same (Power, Volume), but behavior changes based on the device.
This is polymorphism — same interface, but different behaviors.
*/
/*
Go achieves polymorphism using interfaces.
*/

/* // Interface
type Animal interface {
    Sound()
}

// Struct 1
type Dog struct{}

func (d Dog) Sound() {
    fmt.Println("Bark")
}

// Struct 2
type Cat struct{}

func (c Cat) Sound() {
    fmt.Println("Meow")
}

// Main
func main() {
    var a Animal

    a = Dog{}
    a.Sound() // Bark

    a = Cat{}
    a.Sound() // Meow
}

✅ The Sound() method behaves differently depending on the underlying type of Animal.
*/
