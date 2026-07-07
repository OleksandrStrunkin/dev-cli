package env

import (
	"fmt"
	"os"
)

func PrintAll (){
	variable:= os.Environ()

	for _, v := range variable {
		fmt.Println(v)
	}
}