package main

import (
	"flag"
	"fmt"
	"log"
	"practice/logconfig"
	"strconv"
)

// argument:
//   - fizz この数で割り切れるときに"Fizz"と表示する
//   - buzz この数で割り切れるときに"Buzz"と表示する
type FizzBuzz struct {
	fizz int
	buzz int
}

// argument:
//   - num
//
// return:
//   - numがfizzとbuzzの両方で割り切れれば"FizzBuzz"
//   - numがfizzで割り切れれば"Fizz"
//   - numがbuzzで割り切れれば"Buzz"
//   - 上記以外は"num"
func (fzThis *FizzBuzz) fizzBuzz(num int) string {
	if num%(fzThis.fizz*fzThis.buzz) == 0 {
		return "FizzBuzz"
	} else if num%fzThis.fizz == 0 {
		return "Fizz"
	} else if num%fzThis.buzz == 0 {
		return "Buzz"
	} else {
		return strconv.Itoa(num)
	}
}

func main() {
	logger, err := logconfig.SetupLogger("logconfig/config.json")
	if err != nil {
		log.Fatalf("Failed to set log settings: %v", err)
	}
	defer logger.Close()
	log.Println("Setting log configs.. successful")

	num := flag.Int("number", 0, "positive integer")
	flag.Parse()
	fizzBuzz := FizzBuzz{3, 5}
	log.Println("---start---")
	for i := 1; i <= *num; i++ {
		msg := fizzBuzz.fizzBuzz(i)
		log.Printf("%d,%s", i, msg)
		fmt.Println(msg)
	}
	log.Println("---end---")

}
