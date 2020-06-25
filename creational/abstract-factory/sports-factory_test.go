package abstractfactory

import (
	"fmt"
	"testing"
)

const (
	adidasExample string = "adidas"
	nickExample   string = "nike"
)

func TestGetSportFactory(t *testing.T) {
	adidas, _ := getSportFactory(adidasExample)
	nick, _ := getSportFactory(nickExample)

	adidasShoe := adidas.makeShoe()
	adidasShort := adidas.makeShort()

	nickShoe := nick.makeShoe()
	nickShort := nick.makeShort()

	printShoeDetail(adidasShoe)
	printShoeDetail(nickShoe)

	printShortDetail(adidasShort)
	printShortDetail(nickShort)
}

func printShoeDetail(s shoe) {
	fmt.Printf("Logo: %v\n", s.getLogo())
	fmt.Printf("Size: %v\n", s.getSize())
}

func printShortDetail(s short) {
	fmt.Printf("Logo: %v\n", s.getLogo())
	fmt.Printf("Size: %v\n", s.getSize())
}
