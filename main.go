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

	if err := page.Navigate("https://www.w3schools.com/tags/tryit.asp?filename=tryhtml_iframe"); err != nil {
		log.Fatalf("Failed to navigate:%v", err)
	}
	time.Sleep(1 * time.Second)
	fmt.Println(page.Find("#iframeResult").SwitchToFrame())
	time.Sleep(3 * time.Second)
	fmt.Println(page.Find("body > iframe").SwitchToFrame())
	time.Sleep(3 * time.Second)
	fmt.Println(page.Find("body > div.w3-hide-large.w3-hide-large.w3-padding-16 > div:nth-child(1) > a").Text())
	fmt.Println("-------")

	page.Screenshot("./test.jpg")
}
