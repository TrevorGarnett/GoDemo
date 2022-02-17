package gofiles

import (
	"strconv"
)

type Student struct {
	FirstName   string
	LastName    string
	TestAverage float32
	HwAverage   float32
	Average     float32
	Tests       []string
	Hw          []string
}

func (s *Student) AverageTests() {
	var total float32 = 0
	for _, element := range s.Tests {
		num, _ := strconv.Atoi(element)
		total = total + float32(num)
	}
	s.TestAverage = total / float32((len(s.Tests)))
}

func (s *Student) AverageHomeworks() {
	var total float32 = 0
	for _, element := range s.Hw {
		num, _ := strconv.Atoi(element)
		total = total + float32(num)
	}
	s.HwAverage = total / float32((len(s.Hw)))
}

// Utlized by the Class struct to set the overall grade, based on weighting
func (s *Student) setAverage(num float32) {
	s.Average = num
}
