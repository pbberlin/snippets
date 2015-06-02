// syntax highlighting: (https://github.com/mmcgrana/gobyexample/blob/master/tools/generate.go)
// by spawning a [`pygmentize`](http://pygments.org/)´  process from a Go program. 

package main

import "fmt"
import "io/ioutil"
import "os/exec"

func main() {

    dateCmd := exec.Command("date")
    outPipe1, err := dateCmd.Output()
    if err != nil {     fmt.Println("err:", err) ; panic(err) }
    fmt.Println("> date:   " , string(outPipe1))


    grepCmd1 := exec.Command("grep","Left", "-r","*")
    outPipe2, err := grepCmd1.Output()
    if err != nil {     fmt.Println("err:", err)  }
    fmt.Println("> grep1:   \n" , string(outPipe2))


    grepCmd2 := exec.Command("grep", "hello *")
    grepIn , _ := grepCmd2.StdinPipe()
    grepOut, _ := grepCmd2.StdoutPipe()
    grepCmd2.Start()


    grepIn.Write([]byte("hello grep\ngoodbye grep"))
    grepIn.Close()
    grepBytes, _ := ioutil.ReadAll(grepOut)
    grepCmd2.Wait()

    // We only collect the `StdoutPipe` results, stderr would be similar
    fmt.Println("> grep2: " , string(grepBytes))



    // Note that when spawning commands we need to
    // provide an explicitly delineated command and
    // argument array, vs. being able to just pass in one
    // command-line string. If you want to spawn a full
    // command with a string, you can use `bash`'s `-c`
    // option:
    lsCmd := exec.Command("bash", "-c", "ls -a -l -h")
    lsOut, err := lsCmd.Output()
    if err != nil {
        panic(err)
    }
    fmt.Println("> ls -a -l -h")
    fmt.Println(string(lsOut))
}
