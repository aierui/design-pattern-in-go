package factorymethod

import (
	"fmt"
	"testing"
)

func TestGunFactory(t *testing.T) {
	ak47, _ := getGun("ak47")
	maverick, _ := getGun("maverick")
	printDetails(ak47)
	printDetails(maverick)
}

func printDetails(g iGun) {
	fmt.Printf("Gun: %s", g.getName())
	fmt.Println()
	fmt.Printf("Power: %d", g.getPower())
	fmt.Println()
}
