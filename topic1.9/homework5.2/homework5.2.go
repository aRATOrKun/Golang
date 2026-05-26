package main

import "fmt"

type Student struct {
	Name   string
	Grades []float64
}

func (s Student) AverageGrade() float64 {
	var sum float64

	for i := range s.Grades {
		sum += s.Grades[i]
	}

	return sum / float64(len(s.Grades))
}

func main() {
	var student = Student{"Михаил", []float64{5, 5, 4, 3, 5, 4.5}}
	avg := student.AverageGrade()

	fmt.Printf("Студент %s\nСредний балл: %f\n", student.Name, avg)
}
