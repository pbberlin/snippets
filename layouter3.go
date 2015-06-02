package main

import "github.com/pbberlin/tools/transposablematrix"

func main() {

	transposablematrix.InitTB()
	transposablematrix.Main3()
	// transposablematrix.MainTestNoEasternNeighbor()
	// transposablematrix.MainStairyShrinky()

	<-transposablematrix.TermBoxDone
}
