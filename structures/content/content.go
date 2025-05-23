package content

func Structure_content() {
	/*
			   Definition:
			   A struct in Go is a composite data type that "groups together variables under one name".
			   Each variable inside a struct is called a field. Each field has:

			   a name (identifier)
			   a type.

			   Basic struct syntax:
			   type MyStruct struct {
			       x int
			       y int
			       name string
			   }
			   MyStruct is the type name.
			   x, y are int fields.
			   name is a string field.

			-> Empty Struct
			-> Struct With Fields
			-> Embedded Fields
			-> Naming Conflicts
			-> Promoted Fields
			-> Struct Tags - Tags are strings, used by libraries like encoding/json, reflect, protobuf, etc.
			-> Recursive Struct Limitation
					type T5 struct{ f *T5 }         // Pointer, allowed.
					type T6 struct{ f func() T6 }   // Function, allowed.
					type T7 struct{ f [10][]T7 }    // Slice, allowed (slice can be nil).

		Concept							Meaning
		Empty struct					Uses no memory, can act as a marker.
		Embedded fields					Add types directly, auto-named, promote methods/fields.
		Promoted methods				Allow easy access to embedded type's fields/methods.
		Tag strings						Metadata, usually for encoding/decoding.
		Recursive structs				Must avoid self-referencing via arrays/structs; pointers & slices are okay.
	*/
}
