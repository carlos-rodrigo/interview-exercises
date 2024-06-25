/*
*
You are a developer for a university. Your current project is to develop a system for students to find courses they share with friends. The university has a system for querying courses students are enrolled in, returned as a list of (ID, course) pairs.

Write a function that takes in a collection of (student ID number, course name) pairs and returns, for every pair of students, a collection of all courses they share.

Sample Input:

enrollments1 = [

	["58", "Linear Algebra"],
	["94", "Art History"],
	["94", "Operating Systems"],
	["17", "Software Design"],
	["58", "Mechanics"],
	["58", "Economics"],
	["17", "Linear Algebra"],
	["17", "Political Science"],
	["94", "Economics"],
	["25", "Economics"],
	["58", "Software Design"],

]

Sample Output (pseudocode, in any order):

find_pairs(enrollments1) =>

	{
	  "58,17": ["Software Design", "Linear Algebra"]
	  "58,94": ["Economics"]
	  "58,25": ["Economics"]
	  "94,25": ["Economics"]
	  "17,94": []
	  "17,25": []
	}

Additional test cases:

Sample Input:

enrollments2 = [

	["0", "Advanced Mechanics"],
	["0", "Art History"],
	["1", "Course 1"],
	["1", "Course 2"],
	["2", "Computer Architecture"],
	["3", "Course 1"],
	["3", "Course 2"],
	["4", "Algorithms"]

]

Sample output:

find_pairs(enrollments2) =>

	{
	  "1,0":[]
	  "2,0":[]
	  "2,1":[]
	  "3,0":[]
	  "3,1":["Course 1", "Course 2"]
	  "3,2":[]
	  "4,0":[]
	  "4,1":[]
	  "4,2":[]
	  "4,3":[]
	}

Sample Input:
enrollments3 = [

	["23", "Software Design"],
	["3", "Advanced Mechanics"],
	["2", "Art History"],
	["33", "Another"],

]

Sample output:

find_pairs(enrollments3) =>

	{
	  "23,3": []
	  "23,2": []
	  "23,33":[]
	  "3,2":  []
	  "3,33": []
	  "2,33": []
	}

All Test Cases:
find_pairs(enrollments1)
find_pairs(enrollments2)
find_pairs(enrollments3)

Complexity analysis variables:

n: number of student,course pairs in the input
s: number of students
c: total number of courses being offered (note: The number of courses any student can take is bounded by a small constant)
*
*/
package main

import "fmt"

func groupCoursesByStudent(enrollments [][]string) map[string][]string {
	studentsCourses := make(map[string][]string)
	for _, value := range enrollments {
		studentID := value[0]
		course := value[1]

		studentsCourses[studentID] = append(studentsCourses[studentID], course)
	}

	return studentsCourses
}

func matchingCourses(firstStudentCourses, secondStudentCourses []string) []string {
	matchesMap := make(map[string]bool)

	for _, firstValue := range firstStudentCourses {
		matchesMap[firstValue] = true
	}

	matches := []string{}
	for _, secondValue := range secondStudentCourses {
		if matchesMap[secondValue] {
			matches = append(matches, secondValue)
		}
	}

	return matches
}

func combinationExist(combinationResults map[string][]string, firstStudentID, secondStudentID string) bool {
	_, ok := combinationResults[firstStudentID+","+secondStudentID]
	if ok {
		return true
	}
	_, ok = combinationResults[secondStudentID+","+firstStudentID]
	return ok
}

func combineStudentsIDs(coursesByStudent []string) [][2]string {
	studentsAmount := len(coursesByStudent)

	combinations := [][2]string{}

	for i := 0; i < studentsAmount; i++ {
		for j := i + 1; j < studentsAmount; j++ {
			combinations = append(combinations, [2]string{coursesByStudent[i], coursesByStudent[j]})
		}
	}

	return combinations
}

func groupEnrollment(enrollments [][]string) map[string][]string {
	coursesByStudent := groupCoursesByStudent(enrollments)

	studentIDs := make([]string, 0, len(coursesByStudent))
	for key := range coursesByStudent {
		studentIDs = append(studentIDs, key)
	}

	studentsCombination := combineStudentsIDs(studentIDs)
	groupedCourses := make(map[string][]string)

	for _, combination := range studentsCombination {
		groupedCourses[combination[0]+","+combination[1]] = matchingCourses(coursesByStudent[combination[0]], coursesByStudent[combination[1]])
	}

	return groupedCourses
}

func main() {
	enrollments1 := [][]string{
		{"58", "Linear Algebra"},
		{"94", "Art History"},
		{"94", "Operating Systems"},
		{"17", "Software Design"},
		{"58", "Mechanics"},
		{"58", "Economics"},
		{"17", "Linear Algebra"},
		{"17", "Political Science"},
		{"94", "Economics"},
		{"25", "Economics"},
		{"58", "Software Design"},
	}

	enrollments2 := [][]string{
		{"0", "Advanced Mechanics"},
		{"0", "Art History"},
		{"1", "Course 1"},
		{"1", "Course 2"},
		{"2", "Computer Architecture"},
		{"3", "Course 1"},
		{"3", "Course 2"},
		{"4", "Algorithms"},
	}

	enrollments3 := [][]string{
		{"23", "Software Design"},
		{"3", "Advanced Mechanics"},
		{"2", "Art History"},
		{"33", "Another"},
	}

	fmt.Println(groupEnrollment(enrollments1))
	fmt.Println(groupEnrollment(enrollments2))
	fmt.Println(groupEnrollment(enrollments3))
}
