package main

import "fmt"

type name string

type human struct {
	fname name
	lname name
	age   int
}

type richMan struct {
	person
	money int
}

type person struct {
	human
	location   string
	occupation string
}

type studentRoute struct{}

func (sr *studentRoute) GetAll() []person {
	return nil
}
func (sr *studentRoute) GetOne(id int) person {
	return person{}
}
func (sr *studentRoute) Edit(id int) person {
	return person{}
}
func (sr *studentRoute) DeleteOne(id int) bool {
	return true
}
func (sr *studentRoute) create(p person) person {
	return person{}
}

func main() {
	routeForStudent := studentRoute{}

	routeForStudent.

	//structs
	/*
				student1 := person{
					fname: "Khalil",
					lname: "Shams",
					age:   41,
				}

				var student2 person
				student2.fname = "Sanusi"
				student2.lname = "Saidu"
				student2.age = 34

			student3 := person{
				"Aminu", "Sani", 45,
			}

		student1 := person{
			fname: "Khalil",
			lname: "Shams",
			age:   41,
		}
	*/

	var student1 person
	student1.fname = "Jamilu"
	student1.lname = "Musa"
	student1.age = 23
	student1.location = "Kano"
	student1.occupation = "Engineer"

	fmt.Printf("%v\n", student1.toString())
	fmt.Printf("%v\n", student1.toHausa())

}

func (p person) toString() string {
	return fmt.Sprintf("firstname:%v lastname:%v age %v location %v ocupation %v\n", p.fname, p.lname, p.age, p.location, p.occupation)
}

func (p person) toHausa() string {
	return fmt.Sprintf("Sunana:%v, sunan babana:%v shekaru na kuma %v, ina zama a %v aikina kuma %v\n", p.fname, p.lname, p.age, p.location, p.occupation)
}
