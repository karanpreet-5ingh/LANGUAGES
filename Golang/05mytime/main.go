package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("welcome to time study of golang")

	present_time := time.Now()
	fmt.Println(present_time)
	fmt.Println(present_time.Format("01-02-2006 15:04:05 Monday"))
	// Above mention time is the compulsary time that you have to enter to let the code understand which format you are asking

	// yahan pr bss ek date create ki h 2020 february 10 or associated time h 23:23:00
	created_date := time.Date(2020, time.February, 10, 23, 23, 0, 0, time.UTC)

	fmt.Println(created_date)
	fmt.Println(created_date.Format("01-02-2006 Monday"))
	//yahan pr created date ko format kiya h or bol rahe hain k is format me dikha date.

}
