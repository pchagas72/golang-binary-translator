package helper

import (
    "fmt"
)

func FindMaxWithIndex(array []int) (int, int){
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

func ChangeStrIndex(content string, newContent string, index int) string{
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

func Check(err error){
    if err != nil{
        fmt.Println(err)
    }
}

func ReverseString(s string) string {
    runes := []rune(s)
    for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
        runes[i], runes[j] = runes[j], runes[i]
    }
    return string(runes)
}


