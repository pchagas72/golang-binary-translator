package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/pchagas72/golang-binary-translator/helper"
)

func decodeDecimal(binary string) int {
    var result int
    for index, value := range(binary){
        i := len(binary) - index
        if value == '1'{
            result += int(math.Pow(2, float64(i-1)))
        }
    }
    return result
}

func DecodeDecimals(binary string) []int {
    var code []string = strings.Split(binary, " ")
    var result []int
    for _, value := range(code){
        result = append(result, decodeDecimal(value)) 
    }
    return result
}

func EncodeDecimal(decimal int) string{
    var result string
    var biggestTwoExp int = 1
    var twoExponents []int = []int{1}
    var counter int = 0
    for biggestTwoExp < decimal{
        biggestTwoExp += int(math.Pow(2, float64(counter)))
        counter++
        twoExponents = append(twoExponents, biggestTwoExp) 
    }
    for range(twoExponents){
        result += "0"
    }
    counter = 0
    for range(twoExponents){
        index := len(twoExponents)-counter-1
        if decimal >= twoExponents[index]{
            decimal -= twoExponents[index]
            result = helper.ChangeStrIndex(result, "1", index)
        } else{
            result = helper.ChangeStrIndex(result, "0", index)
        }
        counter++
    }
    return helper.ReverseString(result)
}

func EncodeDecimalString(decimalString string) string{
    var result string
    var decimalsString []string = strings.Split(decimalString, " ")
    for _, dec := range(decimalsString){
        decInt, err := strconv.Atoi(dec)
        helper.Check(err)
        result += EncodeDecimal(decInt) + " "
    }
    return  result
}


func main(){

    var decimals []int = []int{1,2,3,4,5}
    for _, dec := range(decimals){
        fmt.Println(EncodeDecimal(dec))
    }

    fmt.Println(DecodeDecimals("100 010 001 0101"))
    fmt.Println(EncodeDecimalString("1 2 3 4 5"))
}
