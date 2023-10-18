package main
import "fmt"

func main() {
	var name string
	var age int
	// take input until space appears
	fmt.Print("Enter your name: ")
	fmt.Scan(&name)
	fmt.Printf("Name: %s\n", name)

	// take multiple inputs
	fmt.Println("Enter your name and age:")
	fmt.Scan(&name, &age)
	fmt.Printf("Name: %s\nAge: %d\n", name, age)

	// take input until new line
	fmt.Println("Enter your name and age:")
	fmt.Scanln(&name, &age)
	fmt.Printf("Name: %s\nAge: %d\n", name, age)
	
	// take input with format specifiers
	fmt.Println("Enter your name and age:")
	fmt.Scanf("%s %d", &name, &age)
	fmt.Printf("Name: %s\nAge: %d\n", name, age)

	// take float and Boolean input
	var temperature float32
	var sunny bool
	fmt.Println("Enter the current temperature:")
	fmt.Scanf("%g", &temperature)
	fmt.Println("Is the day sunny?")
	fmt.Scanf("%t", &sunny)
	fmt.Printf("Current temperature: %g\nIs the day Sunny? %t\n", temperature, sunny)
}
