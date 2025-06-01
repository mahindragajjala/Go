package interfacesegregation

import "fmt"

/*
OldPrinter: Only prints documents.

ScannerPrinter: Can print and scan documents.

MultiFunctionDevice: Can print, scan, and fax.
*/
type Printer interface {
	Print()
}
type Scanner interface {
	Scan()
}
type Faxing interface {
	Fax()
}
type OldPrinter struct{}

func (s OldPrinter) Print() {
	fmt.Println("Printing...")
}

type ScannerPrinter struct{}

func (s ScannerPrinter) Print() {
	fmt.Println("Printing...")
}
func (s ScannerPrinter) Scan() {
	fmt.Println("Scanning...")
}

type MultiFunctionDevice struct{}

func (m MultiFunctionDevice) Print() {
	fmt.Println("Printing...")
}
func (m MultiFunctionDevice) Scan() {
	fmt.Println("Scanning...")
}
func (m MultiFunctionDevice) Fax() {
	fmt.Println("Faxing...")

}
func InterfaceSegregation() {
	oldPrinter := OldPrinter{}
	oldPrinter.Print()
	ScannerPrinter := ScannerPrinter{}
	ScannerPrinter.Print()
	ScannerPrinter.Scan()
	MultiFunctionDevice := MultiFunctionDevice{}
	MultiFunctionDevice.Print()
	MultiFunctionDevice.Scan()
	MultiFunctionDevice.Fax()
}

/*
type Bird interface {
    Fly()
    Swim()
}
What if you have a Penguin? It swims but doesn't fly. You're now forced to implement a meaningless Fly() method.
*/
