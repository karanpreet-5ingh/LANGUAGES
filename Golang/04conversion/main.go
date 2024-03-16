package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to our pizza app")
	fmt.Printf("Please rate our Pizza between 1 and 5 ")

	// ek new reader create kr raha hai using bufio jo ki read krta rahega unless you hit a new line

	reader := bufio.NewReader(os.Stdin)

	input, _ := reader.ReadString('\n')

	fmt.Println("Thanks for rating, ", input)

	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64) //so for conversion we will use this santax strconv to initiate converting and ParseFLoat to convert into float value it has 2 parameters 1 is Input and 2nd is bit size which we usually use 64

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Print("Added 1 to your rating:", numRating+1)
	}

}
