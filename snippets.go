func splitByWhitespace(s string) ( sl []string)  {

    sl = regexp.MustCompile(`[\s]+`).Split(s, 2255)
    for i := 0 ; i < len(sl) ; i++ {
	    //fmt.Println(sl[i])
    }    
    return 
}
