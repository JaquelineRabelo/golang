package main
import "fmt"

func main() {
	fmt.Printf("Type: %T - Value: %v\n", true, true)  //bool 
	fmt.Printf("Type: %T - Value: %v\n","jaqueline", "jaqueline")  //string
	fmt.Printf("Type: %T - Value: %v\n", "steph", "steph")  //string
	fmt.Printf("Type: %T - Value: %v\n", 1, 1) //int 
	fmt.Printf("Type: %T - Value: %v\n", "1", "1") // string
	fmt.Printf("Type: %T - Value: %v\n", 1.233, 1.233)  //flot

	fmt.Printf("Type: %T - Value: %v\n", "Eu programo em Go", "Eu programo em Go")

	

}
//Tipos:
// bool (true/false) 
// string - sequência de bytes
// int - números inteiros
//float (float64/float32)- número decimal
