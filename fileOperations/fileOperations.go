package fileOperations
import(
	"fmt"
	"os"
	"log"
	"strconv"
)

func WriteBalanceToFile(fileName string){
	balanceText:=fmt.Sprint(fileName)
	os.WriteFile(fileName,  []byte(balanceText), 0644)
}

func ReadDataFile(fileName string)(balance float64){
	value,err:=os.ReadFile(fileName)
	if err !=nil{
		log.Fatal(err)
	}
	valueTxt:=string(value)
	userValue,err:=strconv.ParseFloat(valueTxt, 64)
	if err !=nil{
		log.Fatal(err)
	}
	fmt.Printf("initial bank account:%0.2f", userValue)
	return
}
