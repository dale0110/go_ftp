package main

import (
  "time"
	"fmt"
)

func main() {
		t := time.Now
		fmt.Println(t)
		fmt.Println("%d \n",t.Year())
	    fmt.Println("%d \n",t.Month())
	    fmt.Println("%d \n",t.Day())
}
