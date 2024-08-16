package main

import (
	"fmt"
	"math"
)

type binary struct {
    literal     string
}

type binaryCode struct{
    content     []binary
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

func main(){
    var input []string = []string{"001", "010", "100", "110000", "0011001"}
    var tests []binary
    for _, i := range(input){
        tests = append(tests, binary{i}) 
    }
    for _, test := range(tests){
        fmt.Println(test.toDecimal())
    }
}
