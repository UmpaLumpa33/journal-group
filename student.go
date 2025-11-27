package main

import "fmt"

type Student struct {
	FirstName string
	LastName  string
	ThirdName string
	Grades    []int
}

func (s Student) AverageGrade() float64 {
	if len(s.Grades) == 0 {
		return 0
	}
	sum := 0
	for _, grade := range s.Grades {
		sum += grade
	}
	return float64(sum) / float64(len(s.Grades))
}

func (s Student) FullName() string {
	return s.LastName + " " + s.FirstName + " " + s.ThirdName
}

func (s Student) DisplayInfo() {
	fmt.Printf("%s: ", s.FullName())
	for i, grade := range s.Grades {
		if i > 0 {
			fmt.Print(", ")
		}
		fmt.Print(grade)
	}
	fmt.Printf(" (Средний: %.2f)\n", s.AverageGrade())
}
