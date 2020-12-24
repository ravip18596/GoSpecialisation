package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

/*
Write a program which reads information from a file and represents it in a slice of structs. Assume that there is a text file which contains a series of names.
Each line of the text file has a first name and a last name, in that order, separated by a single space on the line.

Your program will define a name struct which has two fields, fname for the first name, and lname for the last name.
Each field will be a string of size 20 (characters).

Your program should prompt the user for the name of the text file. Your program will successively read each line of the text file
and create a struct which contains the first and last names found in the file. Each struct created will be added to a slice,
and after all lines have been read from the file, your program will have a slice containing one struct for each line in the file.
After reading all lines from the file, your program should iterate through your slice of structs and print the first and last names found in each struct.
 */

/*
In order to test this program, pls enter absolute path of the text file as input after the prompt
 */
type User struct {
	fname,lname string
}
func main() {
	fmt.Println("enter absolute path of the file name below")
	var fileName string
	_,err := fmt.Scanf("%s",&fileName)
	if err != nil{
		log.Fatal("error reading file name from stdin using scanf. err is ",err)
		return
	}
	file,err := os.Open(fileName)
	if err != nil{
		log.Fatal("error opening file. err is ",err)
		return
	}
	users := []User{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan(){
		scannerbyteArr := scanner.Bytes()
		byteArr := make([]byte,20)
		copy(byteArr,scannerbyteArr)
		for index,b:=range byteArr{
			if rune(b) == '\n'{
				byteArr = byteArr[:index]
				break
			}
		}
		names := strings.Split(string(byteArr)," ")
		if len(names)>=2 {
			users = append(users, User{
				fname: names[0],
				lname: names[1],
			})
		}
	}
	if err:=file.Close();err != nil{
		fmt.Println("err is ",err)
	}
	for _,user := range users{
		fmt.Println(user)
	}
}

func (u User) String() string {
	return fmt.Sprintf("First name - %s, Last name - %s",u.fname,u.lname)
}