package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	// create empty list

	student_list := []string{} // Initialize an empty list
	fmt.Println(student_list)

	//OPEN TEXT FILE
	//SOURCE: https://gosamples.dev/read-file/#:~:text=The%20simplest%20way%20of%20reading%20a%20text%20or,the%20file%20line%20by%20line%20or%20in%20chunks.
	f, err := os.Open("class_list.txt")

	//if error occurs, print out error
	if err != nil {
		log.Fatal(err)
	}
	//close file
	defer f.Close()

	//scan text file (line by line, not by comma or space)
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		// print the name by line
		// fmt.Println(scanner.Text())
		// but also add to a list
		name := (scanner.Text())
		fmt.Println(name)
		student_list = append(student_list, name)
		fmt.Println(len(student_list))
		// fmt.Println(scanner.Text())
	}
	//if there is an error, print the error
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	//iterate through list of students
	// for i := 0; i < len(student_list(); i++ {
	// 	fmt.Print(student_list[i])
	// }
	//PRINT COMPLETE LIST OF STUDENTS
	fmt.Println(student_list)
}
