package main

import (
	"fmt"
	"golearn/byteoder"
	"io/ioutil"
	"log"
	"net"
	"unsafe"
)

func main() {
	l, err := net.Listen("tcp", "127.0.0.1:9555")
	if err != nil {
		log.Fatalln("Listen error : ", err)
	}
	defer l.Close()

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatalln("Accept erro : ", err)
		}

		go process(conn)
	}
}

func process(conn net.Conn) {

	// for {
	buf, err := ioutil.ReadAll(conn)
	if err != nil {
		fmt.Println("Client reader error : ", err)
	}

	p := *(**byteoder.Head)(unsafe.Pointer(&buf))
	fmt.Println(p.Magic, p.Version, p.Reserve, p.Len)

	conn.Write([]byte("Hello"))
	// }

}
