package main

import (
	"fmt"
	"os"
	"flag"
)

func  flagargs(){
	var user string
	var password string
	var host string
	var port int

	flag.StringVar(&user,"u","root","账号，默认为root")
	flag.StringVar(&password,"p","","密码默认为空")
	flag.StringVar(&host,"h","localhost","主机名，默认为localhost")
	flag.IntVar(&port,"P",3306,"端口号，默认为3306")
	flag.Parse()
	fmt.Printf("user：%v\npassword：%v\nhost：%v\nport：%v\n", user, password, host, port)
}
func  main()  {
	flagargs()
	fmt.Println("命令行的参数有",len(os.Args))
	for i,v:= range os.Args{
		fmt.Printf("args[%v]=%v\n",i,v)
	}
}
