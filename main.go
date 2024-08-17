package main

import (
	"fmt"
	"math"
    "strconv"
    "github.com/pchagas72/golang-binary-translator/helper"
)

type Binary struct {
    literal     string
}

type Decoder struct{
    content     []Binary
}

func  decodeDecimal(bin Binary) int{
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

func (binCode *Decoder) decodeDecimal() []int{
    var result []int
    for _, bin := range(binCode.content){
        result = append(result, decodeDecimal(bin)) 
    }
    return result
}

func encodeDecimal(decimal int) Binary{
    var result Binary
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
            result.literal = helper.ChangeStrIndex(result.literal, "1", index)
        } else{
            result.literal = helper.ChangeStrIndex(result.literal, "0", index)
        }
        counter++
    }
    return Binary{helper.ReverseString(result.literal)}
}

func (binCode *Decoder) encodeDecimal(sequence string) {
    var decimals []int  
    var tempDecimalString string
    var binCodeBlanks Decoder
    for i, c := range(sequence){
        if c != ' '{
            tempDecimalString += string(c)
        } else{
            decimal, err := strconv.Atoi(tempDecimalString)
            helper.Check(err)
            decimals = append(decimals, decimal)
            tempDecimalString = ""
        }
        if i == len(sequence)-1{
            decimal, err := strconv.Atoi(tempDecimalString)
            helper.Check(err)
            decimals = append(decimals, decimal)
        }
    }
    for range(decimals){
        binCodeBlanks.content = append(binCodeBlanks.content, Binary{"0"})
    }
    i := 0
    for range(binCodeBlanks.content){
        binCode.content = append(binCode.content, encodeDecimal(decimals[i]))
        i++
    }
}

func (bin *Decoder) TextToBinary (content string) {
    var code []Binary
    var tempStr string
    for i, value := range(content){
        if value != ' '{
            tempStr += string(value)
        } else{
            code = append(code, Binary{tempStr})
            tempStr = ""
        }
        if i == len(content)-1{
            code = append(code, Binary{tempStr})
            tempStr = ""
        }
    }
    bin.content = code
}

func main(){
    var decimals string = "256 10 12 0 1 2 3 81 29 30"
    var decimals_int []int = []int{256, 10, 12, 0, 1, 2, 3, 81, 29, 30}
    var binCode Decoder
    binCode.encodeDecimal(decimals)
    for index, i := range(binCode.content){
        fmt.Println(decimals_int[index], " = ", i.literal)
    }
    var decoder Decoder
    decoder.TextToBinary("001 010 011 100 110")
    fmt.Println(decoder.decodeDecimal())
    fmt.Println(decoder.encodeDecimal("5"))
}
