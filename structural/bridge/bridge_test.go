package bridge

import (
	"fmt"
	"testing"
)

func TestBridge(t *testing.T) {
	hpPrinter := &hp{}
	epsonPrinter := &epson{}
	macComputer := &mac{}
	macComputer.setPrinter(hpPrinter)
	macComputer.print()

	fmt.Println()

	macComputer.setPrinter(epsonPrinter)
	macComputer.print()

	fmt.Println()

	winComputer := &windows{}
	winComputer.setPrinter(hpPrinter)
	winComputer.print()

	fmt.Println()

	winComputer.setPrinter(epsonPrinter)
	winComputer.print()
	fmt.Println()

	/*
		Print request for mac
		Printing by a HP Printer

		Print request for mac
		Printing by a EPSON Printer

		Print request for windows
		Printing by a HP Printer

		Print request for windows
		Printing by a EPSON Printer
	*/
}
