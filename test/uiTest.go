package main

import (
	"fmt"
	"github.com/gosuri/uilive"
	"time"
)

func main(){
	writer := uilive.New()
	// start listening for updates and render
	writer.Start()

	for i:=0; i<100; i++{
		fmt.Fprintf(writer, "%d/%d\n", i, 100)
		time.Sleep(time.Millisecond * 5)
	}

	fmt.Fprintln(writer, "Finished: Downloaded 100GB")
	writer.Stop() // flush and stop rendering
}
