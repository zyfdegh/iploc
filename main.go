package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"gopkg.in/kataras/iris.v6"
	"gopkg.in/kataras/iris.v6/adaptors/httprouter"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func main() {
	app := iris.New()
	app.Adapt(iris.DevLogger(),
		httprouter.New())
	app.Get("/", handleGetRoot)

	app.Listen(":80")
}

func handleGetRoot(ctx *iris.Context) {
	log.Println("handle GET /")
	host := ctx.Host()
	remoteAddr := ctx.Request.RemoteAddr

	log.Printf("host: %s, remoteAddr: %s\n", host, remoteAddr)

	arr := strings.Split(remoteAddr, ":")
	if len(arr) != 2 {
		err := errors.New("len not 2")
		log.Printf("split remote addr error: %v\n", err)
		return
	}
	ip := arr[0]

	loc, err := ipLoc(ip)
	if err != nil {
		log.Printf("get loc by IP %s error: %v\n", ip)
		return
	}

	fmt.Printf("%+v\n", loc)

	tpl, err := ioutil.ReadFile("./index.tpl")
	if err != nil {
		log.Printf("read html template error: %v\n", err)
		return
	}

	t, err := template.New("index").Parse(string(tpl))
	err = t.Execute(ctx.ResponseWriter, loc)
	if err != nil {
		log.Printf("execute template error: %v\n", err)
		return
	}

	// ctx.HTML(iris.StatusOK, string(tpl))
	return
}

func ipLoc(ip string) (loc Loc, err error) {
	if len(ip) == 0 {
		return
	}

	baseURL := "http://ip.taobao.com/service/getIpInfo.php"
	api := fmt.Sprintf("%s?ip=%s", baseURL, ip)

	resp, err := http.Get(api)
	if err != nil {
		log.Printf("http get error: %v\n", err)
		return
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("read body error: %v\n", err)
		return
	}

	err = json.Unmarshal(data, &loc)
	if err != nil {
		log.Printf("decode json error: %v\n", err)
		return
	}
	return
}
