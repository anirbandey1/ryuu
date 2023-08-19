package main

import (
	"flag"
	"fmt"

	"github.com/anirbandey1/ryuu/repos"
)


func main()  {
    

    flag.Parse()

    args:= flag.Args()

    var firstNonFlagArg string

    if len(args) > 0 {
        firstNonFlagArg = args[0]

    }else {
        fmt.Println("No non-flag argument provided")
    }

    switch firstNonFlagArg {
        case "search":
            fmt.Println("Search")
            if len(args) >= 2{
                fmt.Printf("Search : %s \n",args[1])
                repos.SearchAppimage(args[1])
            }else {
                fmt.Println("Please provide the name of the app you are searching for")
            }
        case "install":
            fmt.Println("Install")
        default:
            fmt.Println("Please provide a valid non-flag argument")

    }



    
}
