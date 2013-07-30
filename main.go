package main

import (
	"fmt"
	"github.com/ChuckHa/calculations/calculations"
//	"labix.org/v2/mgo/bson"
)

func main () {
	// implicitly uses the db in the connection
	c := calculations.GetCollection()

	/*
	// Save 10 new calculations
	for i := 0 ; i < 10; i++ {
		calc := calculations.NewCalculation(fmt.Sprintf("%d + 9", i))
		calc.Save(c)
	}
	*/

	/*
	// Save 1 new calculation and print it
	calc := calculations.NewCalculation("9 + 3")
	calc.Save(c)
	fmt.Println(calc)
	*/

	var err error
	// Get all objects
	var result []calculations.Calculation
	iter := c.Find(nil).Iter()
	err = iter.All(&result)
	if err != nil {
		panic(err)
	}

	// Print each object
//	fmt.Println(result)

	/*
	// Print each objects _id
	for i := 0; i < len(result); i++ {
		fmt.Println(result[i].Id.String())
	}
	*/
	for i := 0; i < len(result); i++ {
		fmt.Println(result[i].Escaped())
	}

	/*
	// Find an object by ID
	result := calculations.Calculation{}
	c.FindId(bson.ObjectIdHex("51c1bb61022c206ec22aca60")).One(&result)
	fmt.Println(result)
	*/

	/*
	// Get one object and update it
	result := calculations.Get(0, c)
	result.OS = "CentOS"
	result.Language = "Python"
	fmt.Println(result)
	*/

}
