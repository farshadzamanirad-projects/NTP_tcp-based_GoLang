package main

import (
        "bufio"
        "fmt"
        "net"
        "os"
        "strings"
        "time"
)

func main() {
        arguments := os.Args
        if len(arguments) == 1 {
                fmt.Println("Please provide port number")
                return
        }

        PORT := ":" + arguments[1]
        l, err := net.Listen("tcp", PORT)
        if err != nil {
                fmt.Println(err)
                return
        }
        defer l.Close()

	for {
        c, err := l.Accept()
        if err != nil {
                fmt.Println(err)
                return
        }


                netData, err := bufio.NewReader(c).ReadString('\n')
                if err != nil {
                        fmt.Println(err)
                        return
                }
                if strings.TrimSpace(string(netData)) == "STOP" {
                        fmt.Println("Exiting TCP server!")
                        return
			        }

                fmt.Print("-> ", string(netData))
                t := time.Now()
 		        formatted := fmt.Sprintf("%d%02d%02d%02d%02d%02d",t.Year(), t.Month(), t.Day(),t.Hour(), t.Minute(), t.Second())
                c.Write([]byte(formatted))
		        }
}
