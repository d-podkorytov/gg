package main

import (
	"fmt"
	"os"
	"os/exec"
        "runtime" 
        "strings"
        "bufio" 
)

func main() {

scanner := bufio.NewScanner(os.Stdin)
if scanner.Err() != nil { fmt.Printf("Error: %v\n",nil) }

for scanner.Scan() {
    // get line from stdin
    s:=strings.Trim(scanner.Text(),"\t")
    fmt.Printf("GG STDIN LINE '%v'\n",s)
 
 if strings.Contains(scanner.Text(), "go get") {
                //fmt.Printf("str_len=%v s='%v'\n",str_len,s)
                words:=strings.Split(s, " ")
                fmt.Printf("GG words =%v\n",words) 
                fmt.Printf("GG PACKAGE ='%v'\n",words[2])                  
                fmt.Printf("GG GO  ='%v'\n",words[0])                  
                fmt.Printf("GG GET ='%v'\n",words[1])                  
                
                //if len(words) == 0 {break}
                //fmt.Printf("words=%v\n",words)
                if words[0] == "go" && words[1] == "get" {  
		fmt.Fprintf(os.Stderr, "OS: %v %v %v\n", runtime.GOOS,os.Args,s,words[2])
                    fmt.Printf("GG WILL FETCH PACKAGE %v\n",words[2])
                    cmd := exec.Command("go","get",words[2])
                    cmd.Stdout = os.Stdout
                    cmd.Stderr = os.Stderr
                
                    e := cmd.Run()
                    if (e!= nil) {fmt.Printf("error %v\n",e)}
                                                         }
               }
	}
}