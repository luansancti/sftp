package main
import (
   "fmt"
   "log"
   "path/filepath"
   "strings"
)



func main() {
   files, err := filepath.Glob("*")
   if err != nil {
	log.Fatal(err)
   }
   
   for i, s := range files {
   	fmt.Println(i,s)
	if(strings.Split(s, ".")[0] == ".") {
		fmt.Println("sdad")
	
	}
   
   }


}

