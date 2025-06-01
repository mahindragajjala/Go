package interfacesegregation

import "fmt"

type Allinterface interface {
	ForPrint()
	ForScan()
	ForFax()
}
type ForOldPrinter struct {
}

func (o ForOldPrinter) ForPrint() {
	fmt.Println("Printing...")
}

type ForScannerPrinter struct {
}

func (f ForScannerPrinter) ForPrint() {
	fmt.Println("Printing...")

}
func (f ForScannerPrinter) ForScan() {
	fmt.Println("Scanning...")

}

type ForMultiFunctionDevice struct {
}

func (m ForMultiFunctionDevice) ForPrint() {
	fmt.Println("Printing...")
}
func (m ForMultiFunctionDevice) ForScan() {
	fmt.Println("Scanning...")
}
func (m ForMultiFunctionDevice) ForFax() {
	fmt.Println("Faxing...")
}
func NormalInterface() {
	ForOldPrinter := ForOldPrinter{}
	ForOldPrinter.ForPrint()
	ForScannerPrinter := ForScannerPrinter{}
	ForScannerPrinter.ForPrint()
	ForScannerPrinter.ForScan()
	ForMultiFunctionDevice := ForMultiFunctionDevice{}
	ForMultiFunctionDevice.ForPrint()
	ForMultiFunctionDevice.ForScan()
	ForMultiFunctionDevice.ForScan()
}
