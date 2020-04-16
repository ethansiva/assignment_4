package main

import (
	"fmt"
	"time"
)

var one owner = owner{firstName: "Ethan", lastName: "Siva", ownersInsurance: insurance{insurance: "Geico", 
	phoneNumber: "237-888-6963", insurancePrice: 1300, address: "1987 Chopstick dr"}, 
	ownerCar: car{price: 30000, make: "Toyota", model: "Camry"}}
var carDirectory map[owner]car = map[owner]car{one: one.ownerCar}

type car struct{
	price int
	make string
	model string
}

type insurance struct{
	insurance string
	phoneNumber string
	address string
	insurancePrice int
	
}

type owner struct{
	firstName string
	lastName string
	ownersInsurance insurance
	ownerCar car
}

func (i *car) getCarName() string{
    return i.make + " " + i.model
}

//Print all the data

func main(){
	loop:
		for{
			//go timer()   //Test with and without the timer
			var userInput int
			fmt.Print("Press 1 to print a stock\nPress 2 to sell a car\nPress 3 to buy a car\n"+
			"Press 4 query\nPress 5 to quit the program\n")
			fmt.Scanln(&userInput)
			switch userInput{
			case 1:
				fmt.Println()
				stock()
				fmt.Println()
				break
			case 2:
				fmt.Println()
				sell()
				fmt.Println()
				break
			case 3:
				fmt.Println()
				buy()
				break
			case 4:
				fmt.Println()
				query()
				break
			case 5:
				fmt.Println("Thank You!")
				break loop
			default:
				fmt.Println("Error")
			}
		}
}
func query(){
	for i,_ := range carDirectory{
		fmt.Println(i.firstName, i.lastName,"\n",i.ownersInsurance.address,"$",i.ownerCar.price,"\n", i.ownersInsurance.insurance,
		i.ownerCar.getCarName(), i.ownersInsurance.phoneNumber )
		fmt.Println()
	}
}

func stock() {
	for i,_ := range carDirectory{
		fmt.Println(i.ownerCar.getCarName())
	}
}

func addNew(firstName, lastName, oi, pn, ad, make, mod string, pcar, pi int){
	newObj := owner{firstName: firstName, lastName: lastName, ownersInsurance: insurance{insurance: oi, phoneNumber: pn, insurancePrice: pi, address: ad},
	ownerCar: car{price: pcar, make: make, model: mod}}
	carDirectory[newObj] = newObj.ownerCar
}

func deleteCar(firstName, lastName string){
	for i,_ := range carDirectory{
		if (firstName == i.firstName) && (lastName == i.lastName){
			delete(carDirectory, i)
		}
	}
}

func sell(){
	var firstName string
	var lastName string
	fmt.Print("Previous Owner' ")
	fmt.Scanln(&firstName, &lastName)
	deleteCar(firstName, lastName)

}
func buy(){
	var firstName, lastName, oi, ad, make, mod, pn string
	var pcar, pi int
	fmt.Print("Car buyer? ")
	fmt.Scanln(&firstName, &lastName)
	fmt.Print("Car price? ")
	fmt.Scanln(&pcar)
	fmt.Print("Owners address? ")
	fmt.Scanln(&ad)
	fmt.Print("Owner's insurance value? ")
	fmt.Scanln(&pi)
	fmt.Print("Owner's number? ")
	fmt.Scanln(&pn)
	fmt.Print("Make and Model? ")
	fmt.Scanln(&make, &mod)
	fmt.Print("Owner's insurance type? ")
	fmt.Scanln(&oi)
	addNew(firstName, lastName, oi, pn, ad, make, mod, pcar, pi)
}

func timer(){
	ticker := time.NewTicker(1 * time.Second)
	for  _ = range ticker.C {
		fmt.Println("ping")
	}
}