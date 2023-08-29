package main

import (
    "fmt"
)
 
func cumprimento(greeting, name string) string {
    return fmt.Sprintf("%v %v", greeting, name)
}
 
func cumprimentoPrefixado (p string) func(string) string {
   return func(n string) string {
        return cumprimento(p,n)
   }
} 

var (
    tarde = cumprimentoPrefixado("Boa tarde")
    manha = cumprimentoPrefixado("Bom dia")
)

func main() {
    fmt.Println(tarde("Ricardo"))
    fmt.Println(manha("Isabel"))
}