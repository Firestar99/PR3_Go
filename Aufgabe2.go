/*
2.)
Modifizieren sie das nachfolgende Programm sodass die main Methode auf den Abschluss einer weiteren von ihnen erstellten Goroutine wartet bevor sie beendet
*/
package main

import (
	"fmt"
)

func main() {
	done := make(chan bool)

	fmt.Println("done")
}
