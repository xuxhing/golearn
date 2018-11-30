package main

import (
	"fmt"
	"golearn/byteoder"
	"log"
	"net"
	"unsafe"
)

func main() {

	// h := &byteoder.Head{
	// 	Magic:   1,
	// 	Version: 2,
	// 	Reserve: 3,
	// 	Len:     4,
	// }

	// data := *(*[]byte)(unsafe.Pointer(h))
	// fmt.Println("[]byte is : ", data)

	conn, err := net.Dial("tcp", "127.0.0.1:9555")
	if err != nil {
		log.Fatalln("Dial error : ", err)
	}

	go process(conn)
	for {

	}
}

func process(conn net.Conn) {
	h := &byteoder.Head{
		Magic:   1,
		Version: 2,
		Reserve: 3,
		Len:     4,
	}

	l := unsafe.Sizeof(*h)
	t := &byteoder.Slice{
		Addr: uintptr(unsafe.Pointer(h)),
		Len:  int(l),
		Cap:  int(l),
	}

	d := *(*[]byte)(unsafe.Pointer(t))
	fmt.Println("[]byte is : ", d)
	conn.Write(d)

	var buf []byte
	for {
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("Read error : ", err)
		}

		fmt.Println(string(buf[0:n]))
	}
}
