package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

type User struct {
	Name string
	Age  int
	Sex  string
}

func main() {
	//simpleTemplate()
	//httpTemplate1()
	httpTemplate2()
}
func httpTemplate2() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		htm, err := template.ParseFiles("./views/user.htm")
		if err != nil {
			fmt.Fprintf(w, "ParseFiles: %v", err)
			return
		}
		err = htm.Execute(w, &User{
			Name: "咖啡",
			Age:  19,
			Sex:  "女",
		})
		if err != nil {
			fmt.Fprintf(w, "Execute: %v", err)
			return
		}
	})

	//监听服务启动
	log.Print("Starting server....")
	log.Fatal(http.ListenAndServe(":80", nil))
}

func httpTemplate1() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		htm, err := template.New("go-web").Parse(`
		User name:{{.Name}}
			  Age:{{.Age}}
			  Sex:{{.Sex}}`)
		if err != nil {
			fmt.Fprintf(w, "Parse: %v", err)
			return
		}
		err = htm.Execute(w, &User{
			Name: "咖啡",
			Age:  19,
			Sex:  "女",
		})
		if err != nil {
			fmt.Fprintf(w, "Execute: %v", err)
			return
		}
	})

	//监听服务启动
	log.Print("Starting server....")
	log.Fatal(http.ListenAndServe(":80", nil))
}

func simpleTemplate() {
	//htm, err := template.New("go-web").Parse("Hello world!{{.}}")
	htm, err := template.New("go-web").Parse(`
		User name:{{.Name}}
			  Age:{{.Age}}
			  Sex:{{.Sex}}`)
	if err != nil {
		log.Fatalf("Parse: %v", err)
	}
	user := User{
		Name: "咖啡",
		Age:  19,
		Sex:  "女",
	}
	err = htm.Execute(os.Stdout, &user)
	if err != nil {
		log.Fatalf("Execute: %v", err)
	}
}
