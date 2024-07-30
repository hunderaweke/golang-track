package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Student struct {
	Name    string
	Grades  map[string]float64
	Average float64
}

func calculateAverage(grades map[string]float64) float64 {
	tot, cnt := 0.0, len(grades)
	for _, g := range grades {
		tot += g
	}
	return tot / float64(cnt)
}

func main() {
	s := bufio.NewReader(os.Stdin)
	fmt.Print("Enter Your name Please: ")
	name, _ := s.ReadString('\n')
	name = strings.Trim(name, "\n")
	var noOfSubjects int
	fmt.Print("How many subjects do you take? ")
	fmt.Scanf("%v", &noOfSubjects)
	grades := make(map[string]float64)
	for i := 1; i <= noOfSubjects; i++ {
		var subName string
		fmt.Printf("Enter the name of subject #%d: ", i)
		fmt.Scanln(&subName)
		var subGrade float64
		for {
			fmt.Printf("Enter the grade of subject #%d (0-100): ", i)
			fmt.Scanln(&subGrade)
			if subGrade >= 0 && subGrade <= 100 {
				break
			} else {
				fmt.Println("Invalid grade. Please enter a value between 0 and 100.")
			}
		}
		grades[subName] = subGrade
	}
	stud := Student{Name: name, Grades: grades}
	if len(grades) > 0 {
		stud.Average = calculateAverage(grades)
	} else {
		stud.Average = 0
	}
	fmt.Printf("\n\n\nFinished Processing %s's data\n", stud.Name)
	fmt.Printf("Student Name: %s\tAverage: %.2f", stud.Name, stud.Average)
	for sub, g := range stud.Grades {
		fmt.Printf("\nSubject Name: %s\t\tGrade: %.2f ", sub, g)
	}
}
