package main

import (
	"encoding/json"
	"fmt"
	"os"
)
type User struct {
	Age int `json:"age"`
	Name string `json:"name"`
}
type Server struct {
	//ID 不会导出到json中
	ID int `json:"-"`
	//ServerName的值会进行二次json编码
	ServerName string `json:"serverName"`
	ServerName2 string `json:"serverName2,string"`
	//ServerIP为空，则不输出到JSON串中
	ServerIP string `json:"serverIP,omitempty"`
}
func main(){
	// 1 字符串转json-map
	org := "[{\"age\":18,\"name\":\"zhangsan\"},{\"age\":20,\"name\":\"lisi\"}]"
	var orgResponse []map[string]interface{}
	json.Unmarshal([]byte(org), &orgResponse)
	for _,v:=range orgResponse{
		fmt.Println("-------------------------")
		for key,value:=range v{
			fmt.Printf("key:%v,value:%v\n",key,value)
		}
		fmt.Println("-------------------------")
	}
	// 2 转struct
	var users  []User
	json.Unmarshal([]byte(org),&users)
	fmt.Println("-------------------------")
	for _,v:=range users{
		fmt.Printf("age:%v,name:%v\n",v.Age,v.Name)
	}
	fmt.Println("-------------------------")
	// 3 struct转json
	s := Server{
		ID:3,
		ServerName:`Go:"1.0"`,
		ServerName2:`Go:"2.0"`,
		ServerIP:``,
	}
	b,_ :=json.Marshal(s)
	os.Stdout.Write(b)
}