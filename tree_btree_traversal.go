package main

//import "code.google.com/p/go-tour/tree"
import "tree"
import "fmt"
import "sort"
import "strings"

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, level int){
		
	p := Pos{1,1} 
	p.Level = level
	p.Value = t.Value
	ch <- p
	//fmt.Print(t.Value, " - \n")
	
	
	if l := t.Left; l != nil {
		Walk( l, level+1)
	} else {
		//fmt.Print("ls",t.Value," ")	
	}

	if r := t.Right; r != nil {
		Walk( r, level+1)
	} else {		
		//fmt.Print("rs",t.Value," ")
	}
	
	
}





func main(){
	t := tree.New(1)
	s := fmt.Sprint(t)
	fmt.Println( s )

	t = tree.New(1)
	fmt.Println( t )
	t = tree.New(1)
	fmt.Println( t )
	
   
	go Walk( t ,1)

//   poss := []*Pos{
//		{20,20},
//   }
	var poss []*Pos
	poss = make([]*Pos,10)
	
	i := 0	
	for {
		i++
		v :=  <- ch
		poss[i-1] = &v
		//fmt.Print(i, ": ", v , "  - ")
		fmt.Print(v , " - ")
		if i> 9{
			fmt.Print("\n")
			break
		}
	}	

	for i,x := range poss {
		fmt.Println(i,x)
	}
	
	fmt.Println("Organs by weight:")
	sort.Sort(ByLevel{poss})
	printPoss(poss)	
	
}


type Tree struct  {
	Left *Tree
	Value int
	Right *Tree

}

type  Pos  struct{
	Level int
	Value int
}
type Poss []*Pos
type ByLevel struct{ Poss }
func (s ByLevel) Less(i, j int) bool { return s.Poss[i].Level < s.Poss[j].Level }
func (s Poss) Len(		 ) int  { return len(s) }
func (s Poss) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

func printPoss(s []*Pos) {
	for _, o := range s {
		indent := strings.Repeat("__",o.Level)
		fmt.Print( indent ,"",  o.Value,"\n")
	}
}

var ch = make( chan Pos,10)

func Same(t1, t2 *tree.Tree) bool {
	return true
}
