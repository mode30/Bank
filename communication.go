package main
import(
	"fmt"
	"log"
	"bufio"
	"strconv"
	"strings"
	"os"
)

func userPrompt(prompt string) (promptValue int,err error) {
	// fmt.Print("welcome to the bank")
	fmt.Print(prompt)
	fmt.Println("what do you want to do:")
	fmt.Println("1.check balance:")
	fmt.Println("2.deposit money:")
	fmt.Println("3.withdraw money:")
	fmt.Println("4.exit:")
	fmt.Print("your chioce:")
	buffer := bufio.NewReader(os.Stdin)
	bufferString, err := buffer.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	input := strings.TrimSpace(bufferString)
	userInput, err := strconv.Atoi(input)
	if err != nil {
		log.Fatal(err)

	}
	return userInput,nil

}
