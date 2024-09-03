package main

import (
	"fmt"
	"log"
	"os"
	"runtime/debug"
)

func handleError(err error) {
	if err != nil {
		fmt.Printf("ERREUR CRITIQUE : %v\n", err)
		fmt.Println("Trace de la pile :")
		debug.PrintStack()
		log.Printf("ERREUR CRITIQUE : %v\n", err)
		os.Exit(1)
	}
}