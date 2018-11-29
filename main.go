package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
	"tank/rest"
)

func main() {

	//将运行时参数装填到config中去。
	rest.PrepareConfigs()

	rest.CONTEXT.Init()
	defer rest.CONTEXT.Destroy()

	http.Handle("/", rest.CONTEXT.Router)

	dotPort := fmt.Sprintf(":%v", rest.CONFIG.ServerPort)

	info := fmt.Sprintf("App started at http://localhost%v", dotPort)
	rest.LogInfo(info)

	err1 := http.ListenAndServe(dotPort, nil)
	if err1 != nil {
		log.Fatal("ListenAndServe: ", err1)
	}
}
