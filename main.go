package main

import (
    "iq"
    "os"
    "fmt"
)

func printHelp() {
    fmt.Println("USAGE")
    fmt.Println("\t./205IQ u s [IQ1] [IQ2]")
    fmt.Println("")
    fmt.Println("DESCRIPTION")
    fmt.Println("\tu\t\tmean")
    fmt.Println("\ts\t\tstandard deviation")
    fmt.Println("\tIQ1\t\tminimum IQ")
    fmt.Println("\tIQ2\t\tmaximum IQ")
}

func main() {
    if iq.CheckHelp() {
        printHelp()
        os.Exit(0)
    }
    if !iq.CheckArgs() {
        printHelp()
        os.Exit(84)
    }
    iq.IQ();
}