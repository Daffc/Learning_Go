package main

import (
	"encoding/xml"
	"fmt"
)

type Plant struct {
	XMLName xml.Name `xml:"plant"`   // Tags 'plant' as the main struct element.
	Id      int      `xml:"id,attr"` // Tags 'Id' as a attribute of 'Plant' instead of an element.
	Name    string   `xml:"name"`    // Tages 'Name' as 'name' element.
	Origin  []string `xml:"origin"`  // Tages 'Origin' as 'origin' element.
}

func (p Plant) String() string {
	return fmt.Sprintf("Plant id=%v, name=%v, origin=%v", p.Id, p.Name, p.Origin)
}

func main() {
	coffee := &Plant{Id: 27, Name: "Coffee"}
	coffee.Origin = []string{"Ehiotia", "Brasil"}
	out, _ := xml.MarshalIndent(coffee, " ", "  ") // Reorganizes XML content with indentation.
	fmt.Println(string(out))

	fmt.Println(xml.Header + string(out)) // Printing XML header and content.

	// Recover XML string to struct variable.
	var p Plant
	if err := xml.Unmarshal(out, &p); err != nil {
		panic(err)
	}
	fmt.Println(p)

	tomato := &Plant{Id: 81, Name: "Tomato"}
	tomato.Origin = []string{"Mexico", "California"}

	type Nesting struct {
		XMLName xml.Name `xml:"nesting"`            // Tages as 'nesting' the main struct
		Plants  []*Plant `xml:"parent>child>plant"` // Tags indicates that elements in 'Plants' should be nested by <parent> and <child> XML elements.
	}

	nesting := &Nesting{}
	nesting.Plants = []*Plant{coffee, tomato}

	out, _ = xml.MarshalIndent(nesting, " ", "  ")
	fmt.Println(string(out))
}
