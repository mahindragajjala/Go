package coupled

func Tightly_coupled() {
	type Frist_structure struct {
	}
	type Second_structure struct {
		Frist_structure
	}
	/*
		This technically compiles — but it doesn’t fully reflect the idea of tight coupling.
		This is just struct embedding (composition), not method calling or dependency.

		type Printer struct{}

		func (p *Printer) Print(msg string) {
			fmt.Println(msg)
		}

		type Report struct {
			printer Printer
		}

		func (r *Report) Generate() {
			r.printer.Print("Generating Report")
		}
	*/

}

func Decouple() {
	type Frist struct {
	}
	type Second struct {
	}
	type Mix interface {
		Frist
		Second
	}
	/*
			You cannot embed structs inside an interface like this!
			Interfaces in Go only embed other interfaces or define methods, not structs.

		type Output interface {
		Print(msg string)
		}

		type Printer struct{}

		func (p *Printer) Print(msg string) {
		    fmt.Println(msg)
		}

		type Report struct {
		    output Output
		}

		func (r *Report) Generate() {
		    r.output.Print("Generating Report")
		}

	*/
}
