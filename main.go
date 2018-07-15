package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	data := ""

	if len(os.Args) > 1 {
		// get the supplied args
		for _, s := range os.Args[1:] {
			data += s + " "
		}

	}

	if len(data) == 0 {
		// no data so trying stdin
		s := bufio.NewScanner(os.Stdin)
		for s.Scan() {
			data = s.Text()
		}
	}

	data = strings.Trim(data, " ")
	numbers := strings.Split(data, " ")

	for _, n := range numbers {
		i, err := strconv.ParseInt(n, 10, 64)
		if err != nil {
			panic(err)
		}

		h := fmt.Sprintf("%#x", i)
		clean := strings.Replace(h, "0x", "", 1)
		fmt.Printf("%v ", clean)
	}

	fmt.Printf("\n")
}
