package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

var totalAmount float64=1_000_000
func main() {

	for{

		prompt,err := userPrompt("welcome to the bank:")
		if err !=nil{
			log.Fatal(err)
		}
		fmt.Println("input:", prompt)
		match(prompt)
	}
}

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

func match(promptValue int){
	switch promptValue{


		case 1:
		checkAmount("Available amount:")
		// fmt.Print("1")
		case 2:
		amount,_,err:=depositAmount("enter amount to deposit:")
		if err !=nil{
			log.Fatal(err)
		}
		fmt.Println("amount deposited:",amount)
		fmt.Print("2")
		case 3:
		amount:=withdrawAmount("enter amount to withdraw:")
		fmt.Print(amount)
		case 4:
		fmt.Print("program exited:")
		os.Exit(0)
		default:
		fmt.Print("Program ended abruptly")
	}

}

func withdrawAmount(prompt string)float64{
	var withdrawAmount float64
	fmt.Print(prompt)
	fmt.Scan(&withdrawAmount)
	fmt.Print("total amount:",totalAmount)
	remainingAmount:=totalAmount-withdrawAmount
	if remainingAmount < totalAmount{
		fmt.Println("Please enter valid amount,amount in account too low to withdraw")
		os.Exit(0)//ends the program abruptly
	}else{

	totalAmount=remainingAmount
	return totalAmount
	}
	// fmt.Println("remaining Amount:",remainingAmount)
}


func checkAmount(prompt string){
	fmt.Print(prompt)
	fmt.Printf("bank amount: %0.2f",totalAmount)
}


func depositAmount(prompt string)(amount float64,condition bool,err error){
	fmt.Println(prompt)
	fmt.Print("")
	buffer:=bufio.NewReader(os.Stdin)
	bufferString,err:=buffer.ReadString('\n')
	if err !=nil{
		log.Fatal(err)
	}
	amountDepositedString:=strings.TrimSpace(bufferString)
	amountDeposited,err:=strconv.ParseFloat(amountDepositedString, 64)
	if err !=nil{
		log.Fatal(err)
	}
	// negNUmber:=math.Inf(0)
	negNumber:=0
	if amountDeposited <= float64(negNumber){
		fmt.Print("enter valud number greater than 0")
		return float64(negNumber),false,nil

	}
	fmt.Printf("amount deposited:%f",amountDeposited)
	totalAmount+=amountDeposited

	return totalAmount,true,nil
}
