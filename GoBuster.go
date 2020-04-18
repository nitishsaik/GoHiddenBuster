package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime"
)


/*
Change os.open("/usr/share/wordlist/**.txt) or any custom location

change URL String ="xxxxxxxx" with custom domain 
Feel free to contribute



*/
func main() {

	runtime.GOMAXPROCS(4)

	file, err := os.Open("./xxxx.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		var URL string = "https://xxxxxxxxxx" + scanner.Text()
		resp, err := http.Get(URL)
		if err != nil {
			log.Fatal(err)
		}
		if resp.StatusCode == 200 {
			fmt.Println(resp.Status, URL)
		}

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}
