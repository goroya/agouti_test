package main

import (
	"encoding/csv"
	"github.com/sclevine/agouti"
	"fmt"
	"log"
	"flag"
	"os"
	"time"
)



func main() {
	flag.Parse()

	file2, err := os.Create("./test.csv")
	defer file2.Close()

	line := []string{"apple", "orange", "lemon"}
	writer := csv.NewWriter(file2)

	writer.Write(line)
	writer.Flush()

	fmt.Println("Hello")
	driver := agouti.PhantomJS()
	if err := driver.Start(); err != nil {
		log.Fatalf("Failed to start driver:%v", err)
	}
	defer driver.Stop()

	page, err := driver.NewPage(agouti.Browser("phantomjs"))
	if err != nil {
		log.Fatalf("Failed to open page:%v", err)
	}

	if err := page.Navigate("https://www.w3schools.com/html/html_examples.asp"); err != nil {
		log.Fatalf("Failed to navigate:%v", err)
	}
	time.Sleep(1 * time.Second)
	fmt.Println(page.Find("#footer"))
	fmt.Println("-------")

	//page.Screenshot("./test.jpg")
}
