package main

import (
	"fmt"
	"go-zero-demo/common"
	"os"
	"sync"
)

//func say(s string) {
//	fmt.Printf("data:%v\n", s)
//}
//
//func main() {
//	num := "2"
//	defer say(num)
//	num = "4"
//	println("-------------->%v", num)
//
//	var m map[int]string
//	m = make(map[int]string, 10)
//	m[2] = "value"
//
//	delete(m, 2)
//
//	fmt.Printf("map data:%v", m[2])
//
//	for k, v := range m {
//		fmt.Printf("range map key:%v \n,map value:%v \n", k, v)
//	}
//
//	newT := Te{"名字", 18}
//	fmt.Printf("Te =---------->:%d \n", newT.age)
//}

var wait sync.WaitGroup

type Te struct {
	Name string
	age  int
	aa   any
}

func (t Te) changeName(n string) {
	(t).Name = n
	fmt.Printf("---------->tName:%v \n", t.Name)
}

func (t Te) GetName() string {
	return t.Name
}

func init() {
	fmt.Println("init")
}

func main() {
	os.ReadFile("....")
	te := Te{
		age:  10,
		Name: "name1",
	}
	//ty := reflect.TypeOf(te)
	var name = "123"
	name1 := "222222"
	fmt.Printf(name)
	fmt.Printf(name1)
	fmt.Printf("first name --------> %v \n", te.Name)
	te.changeName("changeName")
	fmt.Printf("second name --------> %v \n", te.Name)
	fmt.Println(common.Version)

	myList := []int{3, 1, 2}
	fmt.Printf("%v \n", myList)
	xiecheng()
	var command string
	//bo := true
	fmt.Scanln(&command)
	//fmt.Println(&command)
	//fmt.Print("zcc")
	//fmt.Print(bo)
	//addr, _ := net.ResolveIPAddr("tcp", "127.0.0.1:8029")
	//http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
	//	fmt.Println(request.URL.Path)
	//	writer.Write([]byte("zxczxc"))
	//})
	//http.ListenAndServe("127.0.0.1:8209", nil)
	//r := gin.Default()
	//r.GET("/", func(c *gin.Context) {
	//	c.String(200, "Hello, Geektutu")
	//})
	//r.Run() // listen and serve on 0.0.0.0:8080
	//database.DB{}
	//
	//addr, _ := net.ResolveTCPAddr("tcp", "127.0.0.1:8029")
	//connection, err := net.ListenTCP("tcp", addr)
	//if err != nil {
	//	fmt.Println("发生链接错误", err)
	//}
	//for {
	//	conn, err := connection.Accept()
	//	if err != nil {
	//		break
	//	}
	//	conn.Write([]byte("asdas"))
	//}
}

type Te2 struct {
	Te
	More string
}

type Te22 interface {
	say(n string)
}

func xiecheng() {
	//r := gin.Default()
	//r.GET("/testGet", controller.UserController{}.Test)
	//r.Run(":8090")
	//runtime.GC()
	getG

}

//go gin database grpc go-zero
