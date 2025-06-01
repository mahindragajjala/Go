package content

func Struct_Definition_Grammar() {
	/*
				StructType    = "struct" "{" { FieldDecl ";" } "}" .
				FieldDecl     = (IdentifierList Type | EmbeddedField) [ Tag ] .
				EmbeddedField = [ "*" ] TypeName [ TypeArgs ] .
				Tag           = string_lit .

		---------------------------StructType = "struct" "{" { FieldDecl ";" } "}" .---------------------------
											type Employee struct {
												Name    string
												Age     int
												Address string
												Email   string
												Salary  float64
											}

											func main() {
												emp := Employee{
													Name:    "John Doe",
													Age:     30,
													Address: "123 Main St",
													Email:   "john@example.com",
													Salary:  50000.50,
												}

												fmt.Println(emp)
											}

		---------------------------FieldDecl = (IdentifierList Type | EmbeddedField [ Tag ].---------------------------
												type Person struct {
													Name string;
													Age int;
												}

		---------------------------EmbeddedField = [ "*" ] TypeName [ TypeArgs ] .---------------------------
												type Address struct {
													City    string
													Pincode int
												}

												type Customer struct {
													Name string
													Address  // Embedded!
												}

		--------------------------------------Tag Usage./Tag = string_lit .---------------------------
												type APIUser struct {
													UserName string `json:"username"`
													Email    string `json:"email"`
												}

		------------------------------------------Empty struct.---------------------------
											var empty struct{}
											fmt.Println(empty)  // prints: {}
											or in map
											visited := map[string]struct{}{}
											visited["google.com"] = struct{}{}

		---------------------------------Invalid Recursive Struct---------------------------

		This causes infinite memory recursion.
		type Node struct {
			next Node  // ❌ Invalid: Node contains itself.
		}
		type Node struct {
			next *Node  // Pointer is allowed.
		}

		Concept	                     Real-World Usage
		struct {}	                 Marker, zero memory, signal-only types.
		Named Fields	             Models data (Employee, Product, Order).
		Embedded Fields	             Promote reuse without explicit naming.
		Tags	                     Metadata for JSON, XML, Protobuf, DB.
		Invalid Recursion	         Prevent infinite memory loops.
		Valid Recursion	             Allow pointers to same type.
	*/
}
