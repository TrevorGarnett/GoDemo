package gofiles

import (
	"fmt"

	table "github.com/rodaine/table"
)

type Class struct {
	Students      []*Student
	TestWeighting int
	TotalTests    int
	HWWeighting   int
	TotalHWs      int
	ClassAverage  float32
}

func (a *Class) FindScore() {
	var total float32 = 0
	var average float32 = 0
	for _, element := range a.Students {
		average = (float32(a.TestWeighting)*element.TestAverage + float32(a.HWWeighting)*element.HwAverage) / float32(100)
		total += average
		// Element is a student
		element.setAverage(average)
	}
	a.ClassAverage = total / float32(len(a.Students))
}

func (a *Class) PrintSummary() {
	// Print Summary
	fmt.Println("\n\nGrade Report --- ", len(a.Students), " Students found in File")
	fmt.Println("Test Weight:", a.TestWeighting, "%")
	fmt.Println("Homework Weight:", a.HWWeighting, "%")
	fmt.Printf("Overall Average: %.1f %c \n", a.ClassAverage, '%')
	fmt.Println()
	tbl := table.New("Student Names", "Tests", "Homeworks", "Average", "Notes")
	for _, element := range a.Students {
		name := element.LastName + ", " + element.FirstName
		t := fmt.Sprintf("%f", element.TestAverage) + "(" + fmt.Sprintf("%d", len(element.Tests)) + ")"
		h := fmt.Sprintf("%f", element.HwAverage) + "(" + fmt.Sprintf("%d", len(element.Hw)) + ")"
		var note string = ""
		if len(element.Tests) < a.TotalTests {
			if len(element.Hw) < a.TotalHWs {
				note = "** Missing multiple assignments **"
			} else {
				note = "** May be missing a test **"
			}
		} else if len(element.Hw) < a.TotalHWs {
			note = "** May be missing a homework **"
		}
		tbl.AddRow(name, t, h, element.Average, note)
	}
	tbl.Print()
	fmt.Println()
}
