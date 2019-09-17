package main
 
import (
    "log"
    "net/http"
    "html/template"
    "time"
)
 
func main() {
    //规则1
    http.HandleFunc("/", handler)
     
    //规则2
    http.HandleFunc("/file/", file)
 
    log.Fatal(http.ListenAndServe(":8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {	
    t, _ := template.ParseFiles("header.html")
    err := t.Execute(w, map[string]string{"Time": time.Now().Format("2006-01-02 15:04:05"), "Name": "Fox"})

    if err != nil {
        panic(err)
    }
}

func file(w http.ResponseWriter, r *http.Request) {	 
    w.Write([]byte("pattern path: /file/ "))
}
