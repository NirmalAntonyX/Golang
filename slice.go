package main
import "fmt"

func main() {
  a := []int{10,20,30}
  sum := 0
  for key,value := range a {
	  sum += value
	  fmt.Println(key, value)
  }
 fmt.Println(sum)
  
}