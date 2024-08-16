package main

import (
	"fmt"
	"math"
    "strconv"
)

type binary struct {
    literal     string
}

type binaryCode struct{
    content     []binary
}

func findMaxWithIndex(array []int) (int, int){
    var max_value int = array[0]
    var max_value_index int
    for index, i := range(array){
        if i > max_value{
            max_value = i
            max_value_index = int(index)
        }
    }
    return max_value, max_value_index
}

func changeStrIndex(content string, newContent string, index int) string{
    var newString string
    for i, c := range(content){
        if i == index{
            newString += newContent 
            continue
        }
        newString += string(c)
    }
    return newString
}

func check(err error){
    if err != nil{
        fmt.Println(err)
    }
}

func reverse_string(s string) string {
    runes := []rune(s)
    for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
        runes[i], runes[j] = runes[j], runes[i]
    }
    return string(runes)
}

func (bin *binary) toDecimal() int{
    var result int
    literal := bin.literal
    for index, value := range(literal){
        i := len(literal) - index
        if value == '1'{
            result += int(math.Pow(2,float64(i-1)))
        }
    }
    return result
}

func (binCode *binaryCode) toDecimal() []int{
    var result []int
    for _, bin := range(binCode.content){
        result = append(result, bin.toDecimal()) 
    }
    return result
}

func decimalToBinary(decimal int) binary{
    var result binary
    var biggestTwoExp int = 1
    var twoExponents []int = []int{1}
    var counter int = 0
    for biggestTwoExp < decimal{
        biggestTwoExp += int(math.Pow(2, float64(counter)))
        counter++
        twoExponents = append(twoExponents, biggestTwoExp) 
    }
    for range(twoExponents){
        result.literal += "0"
    }
    counter = 0
    for range(twoExponents){
        index := len(twoExponents)-counter-1
        if decimal >= twoExponents[index]{
            decimal -= twoExponents[index]
            result.literal = changeStrIndex(result.literal, "1", index)
        } else{
            result.literal = changeStrIndex(result.literal, "0", index)
        }
        counter++
    }
    return binary{reverse_string(result.literal)}
}

func (binCode *binaryCode) decimalToBinary(sequence string) {
    var decimals []int  
    var tempDecimalString string
    var binCodeBlanks binaryCode
    for i, c := range(sequence){
        if c != ' '{
            tempDecimalString += string(c)
        } else{
            decimal, err := strconv.Atoi(tempDecimalString)
            check(err)
            decimals = append(decimals, decimal)
            tempDecimalString = ""
        }
        if i == len(sequence)-1{
            decimal, err := strconv.Atoi(tempDecimalString)
            check(err)
            decimals = append(decimals, decimal)
        }
    }
    for range(decimals){
        binCodeBlanks.content = append(binCodeBlanks.content, binary{"0"})
    }
    i := 0
    for range(binCodeBlanks.content){
        binCode.content = append(binCode.content, decimalToBinary(decimals[i]))
        i++
    }
}

func (bin *binaryCode) stringToBinary (content string) {
    var code []binary
    var tempStr string
    for i, value := range(content){
        if value != ' '{
            tempStr += string(value)
        } else{
            code = append(code, binary{tempStr})
            tempStr = ""
        }
        if i == len(content)-1{
            code = append(code, binary{tempStr})
            tempStr = ""
        }
    }
    bin.content = code
}

func main(){
    var decimals string = "256 10 12 0 1 2 3"
    var decimals_int []int = []int{256, 10, 12, 0, 1, 2, 3}
    var binCode binaryCode
    binCode.decimalToBinary(decimals)
    for index, i := range(binCode.content){
        fmt.Println(decimals_int[index], " = ", i.literal)
    }
}
