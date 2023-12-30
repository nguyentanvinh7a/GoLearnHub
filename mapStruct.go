package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	number int
	name   string
	isMale bool
	subjects []string
}

type People struct {
	name string
	age int
}

type Teacher struct {
	People
	number int `Maximum can't over 100`
	subjects []string
}

func main() {
	studentNameAgeMap := make(map[string]int)
	studentNameAgeMap = map[string]int{
		"John": 32,
		"Rob":  28,
	}
	studentNameAgeMap["Sam"] = 31
	delete(studentNameAgeMap, "Rob")
	fmt.Println(studentNameAgeMap)
	m := map[[3]string]int{}
	fmt.Println(m)
	_, isExist := studentNameAgeMap["Rob"]
	fmt.Println(isExist)

	studentYuh := Student{
		1,
		"Yuh",
		true,
		[]string{"Math", "English"},
	}
	fmt.Println(studentYuh.name)

	studentYuh1 := Student{}
	studentYuh1.name = "Yuh"
	studentYuh1.number = 1
	studentYuh1.isMale = true
	studentYuh1.subjects = []string{"Math", "English"}
	fmt.Println(studentYuh1.name)

	studentYuh2 := struct { name string }{name: "Yuh"}
	fmt.Println(studentYuh2.name)

	teacherYuh := Teacher{}
	teacherYuh.name = "Yuh"
	teacherYuh.age = 30
	teacherYuh.number = 1
	teacherYuh.subjects = []string{"Math", "English"}
	t := reflect.TypeOf(teacherYuh)
	field, _ := t.FieldByName("number")
	fmt.Println(field.Tag)
}