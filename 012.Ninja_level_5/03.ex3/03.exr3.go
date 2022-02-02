package main

import "fmt"

// Create a new type vehicle:
// underlaying type is a struct:
// the fields are: doors, color
// Create two types: trucks & sedan.
// underlaying type of each is a struct.
// embed the vehicle struct type as a field.
// give the trcuks the field: fourWheel which is bool
// give the sedan the field: luxury which is bool
// using the vehicle, truck, and sedan structs:
// using a composite literal, create a value of type truck and assign values to the fields;
// using a composite literal, create a value of type sedan and assign values to the fields.
// print out each of the values.
// print out a single field from each of the values.

type vehicle struct {
	doors int
	color string
}

type truck struct {
	vehicle
	fourWheel bool
}

type sedan struct {
	vehicle
	luxury bool
}

func main() {
	vh := vehicle{
		doors: 4,
		color: "black",
	}

	tr := truck{
		vehicle:   vh,
		fourWheel: true,
	}

	sd := sedan{
		vehicle: vh,
		luxury:  true,
	}

	fmt.Println(tr)
	fmt.Println(tr.vehicle)
	fmt.Println(tr.doors)
	fmt.Println(tr.color)
	fmt.Println(tr.fourWheel)

	fmt.Println("---------------")

	fmt.Println(sd)
	fmt.Println(sd.vehicle)
	fmt.Println(sd.doors)
	fmt.Println(sd.color)
	fmt.Println(sd.luxury)

}
