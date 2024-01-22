package main

//
//func main() {
//
//	connection, err := net.Dial("tcp", "127.0.0.1:8029")
//	if err != nil {
//		fmt.Println("发生链接错误", err)
//	}
//	for {
//		var buf []byte = make([]byte, 1024)
//		n, err := connection.Read(buf)
//		if err == io.EOF {
//			break
//		}
//		fmt.Printf(string(buf[0:n]))
//	}
//
//}
