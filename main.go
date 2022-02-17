package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"

	gofiles "github.com/TrevorGarnett/GoDemo/Student"
)

func main() {
	// Ask user to input file name
	fmt.Println("Please enter the name of the file you would like to use.")
	var file string
	fmt.Scanln(&file)

	// Ask user to input weighting (in %) for tests
	fmt.Println("Please enter the weight applied to tests.")
	var weight int
	fmt.Scanln(&weight)

	// Ask the user to input the number of tests
	fmt.Println("How many tests are there")
	var numTests int
	fmt.Scanln(&numTests)

	// Ask the user to input the number of homeworks
	fmt.Println("How many homeworks are there")
	var numHW int
	fmt.Scanln(&numHW)

	// open file
	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	// remember to close the file at the end of the program
	defer f.Close()

	// read the file line by line using scanner
	scanner := bufio.NewScanner(f)

	var students []*gofiles.Student
	for scanner.Scan() {
		var name string = scanner.Text()
		var names []string = strings.Split(name, " ")
		// Test Grades
		scanner.Scan()
		var test string = scanner.Text()
		var tests []string = strings.Split(test, " ")
		// Homework Grades
		scanner.Scan()
		var homework string = scanner.Text()
		var homeworks []string = strings.Split(homework, " ")

		student := gofiles.Student{
			FirstName: names[0],
			LastName:  names[1],
			Tests:     tests,
			Hw:        homeworks}

		student.AverageHomeworks()
		student.AverageTests()
		students = append(students, &student)
	}
	// Sort the students by last name and then first name.
	sort.Sort(gofiles.ByName(students))

	class := gofiles.Class{
		Students:      students,
		TestWeighting: weight,
		TotalTests:    numTests,
		HWWeighting:   (100 - weight),
		TotalHWs:      numHW}

	class.FindScore()

	class.PrintSummary()
	// for _, element := range class.Students {
	// 	fmt.Println(element)
	// }

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
