package codego

/*
In Go, polymorphism and abstraction are *related* concepts but not the same. Let's break it down clearly:



 🔶 1. Abstraction (Hiding details and exposing behavior)

* Goal: Expose only *what is necessary* and hide the internal working.
* In Go: Achieved using interfaces.
* Example:


type Shape interface {
    Area() float64  // abstraction: caller doesn't care how Area is calculated
}


Here, the caller knows `Area()` exists, but not how it’s implemented.



 🔷 2. Polymorphism (Many forms of behavior)

* Goal: Different types respond differently to the same interface.
* In Go: Achieved when multiple types implement the same interface, and we call their methods through that interface.

Example continued:


type Circle struct {
    Radius float64
}

func (c Circle) Area() float64 {
    return 3.14 * c.Radius * c.Radius
}

type Rectangle struct {
    Width, Height float64
}

func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}

func PrintArea(s Shape) {
    fmt.Println("Area:", s.Area())
}


Now, both `Circle` and `Rectangle` abstract the `Area()` method and provide their own polymorphic implementations.



 ✅ In Summary:

| Concept      | Meaning                             | How Go supports it   |
| Abstraction  | Hiding internal logic               | Interfaces           |
| Polymorphism | One interface, many implementations | Structs + Interfaces |


+--------------------------+
|        INTERFACE         |
|    type Shape interface  |
|    {                     |
|       Area() float64     | <-- Abstraction: defines behavior
|    }                     |
+-----------+--------------+
            |
            |                   (Polymorphism: different structs implement the same interface)
   +--------+--------+        +---------------------+
   |                 |        |                     |
+--v--+         +----v---+  +--v-----------------+ +--v------------------+
|Circle         |Rectangle |  | func (c Circle)   | | func (r Rectangle) |
|Radius float64 |Width, Ht |  | Area() float64 {  | | Area() float64 {   |
+-------------- +----------+  | return πr²        | | return w × h       |
                              +-------------------+ +--------------------+

         ⬇ Used like this
         func PrintArea(s Shape) {
             fmt.Println(s.Area())
         }

         // Accepts Circle or Rectangle → Polymorphism

*/
