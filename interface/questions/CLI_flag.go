package questions

import (
	"flag"
	"fmt"
	"net"
	"os"
)

type IPFlag struct {
	IP net.IP
}

// Implement the `flag.Value` interface.
func (i *IPFlag) String() string {
	if i.IP == nil {
		return ""
	}
	return i.IP.String()
}

func (i *IPFlag) Set(value string) error {
	ip := net.ParseIP(value)
	if ip == nil {
		return fmt.Errorf("invalid IP address: %s", value)
	}
	i.IP = ip
	return nil
}
func Cli_Flag() {
	var ipFlag IPFlag

	// Register the flag.
	flag.Var(&ipFlag, "ip", "IP address to bind")

	// Parse CLI flags.
	flag.Parse()

	// Check if set
	if ipFlag.IP == nil {
		fmt.Println("No valid IP address provided.")
		os.Exit(1)
	}

	fmt.Println("Parsed IP Address:", ipFlag.IP)
}
