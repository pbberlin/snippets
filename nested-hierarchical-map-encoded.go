package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

func main() {

	m := map[string]interface{}{}

	m["a-string"] = "A String"
	m["an-int"] = 32168
	m["a-float"] = 1.00440
	m["tm"] = time.Now()

	submap := map[string]interface{}{}
	submap["smk1"] = 3223
	submap["smk2"] = "Hosianna"

	m["submap"] = submap

	fmt.Printf("%#v\n\n", m)

	b, err := json.MarshalIndent(m, "-->", "\t")
	if err != nil {
		fmt.Println("error:", err)
	}
	os.Stdout.Write(b)

}
