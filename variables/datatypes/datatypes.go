package datatypes

/*
-------------------Type     = TypeName [ TypeArgs ] | TypeLit | "(" Type ")" .

			TypeName - name of the variable
			TypeArgs - int/slice/map
			typelist - go generic's
			type 	 - var f func(x (int)) // x is of type int, parentheses don't change that

-------------------TypeName = identifier | QualifiedIdent

identifier     - Type from the current package - Person, int
QualifiedIdent - Type from an imported package - time.Time

-------------------TypeArgs = "[" TypeList [ "," ] "]" .
[ ... ]    - Encloses type arguments
TypeList   - One or more types separated by commas
[ "," ]    - Optional trailing comma (Go allows this)

-------------------TypeList = Type { "," Type } .
func MakeTuple[A, B any](a A, b B) (A, B) {
    return a, b
}
MakeTuple[int, string](1, "hi") // 👈 TypeList = int, string

-------------------TypeLit = ArrayType | StructType | PointerType | FunctionType | InterfaceType |SliceType | MapType | ChannelType .
|------------------|---------------------------------|--------------------------------|
| Type Literal     | Example                         | Description                    |
|------------------|---------------------------------|--------------------------------|
| ArrayType        | `[3]int`                        | Fixed-length array             |
| SliceType        | `[]string`                      | Dynamic slice                  |
| StructType       | `struct { Name string }`        | Anonymous struct               |
| PointerType      | `*int`                          | Pointer to int                 |
| FunctionType     | `func(int) string`              | Function signature             |
| InterfaceType    | `interface { Read([]byte) }`    | Anonymous interface            |
| MapType          | `map[string]int`                | Map from string to int         |
| ChannelType      | `chan int`, `<-chan int`        | Channel for communication      |
|------------------|---------------------------------|--------------------------------|



*/
