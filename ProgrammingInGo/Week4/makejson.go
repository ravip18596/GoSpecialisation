package main

/*
Problem statement
Write a program which prompts the user to first enter a name, and then enter an address. Your program should create a map and add the name
and address to the map using the keys “name” and “address”, respectively. Your program should use Marshal() to create a JSON object from the map,
and then your program should print the JSON object.
 */

/*func main() {
	fmt.Println("Enter a name")
	scanner := bufio.NewScanner(os.Stdin)
	user := make(map[string]string)
	if scanner.Scan(){
		fmt.Println("Enter your name.")
		name := scanner.Text()
		user["name"] = name
	}
	if scanner.Scan(){
		fmt.Println("Enter your address.")
		address := scanner.Text()
		user["address"] = address
	}
	obj,err := json.Marshal(user)
	if err != nil{
		log.Fatal("error in json to byte serialization. err is ",err)
		return
	}
	fmt.Println("json object is ",string(obj))
}*/