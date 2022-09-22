package main

import (
	"fmt"
	"net/http"
    //"os"
    //"io"
    //"path/filepath"
    //"database/sql"
    //_ "github.com/go-sql-driver/mysql"
    "html/template"
    "encoding/json"
)


type student struct {
    ID    string
    Name  string
    Grade int
}

var data = []student{
    student{"E001", "ethan", 21},
    student{"W001", "wick", 22},
    student{"B001", "bourne", 23},
    student{"B002", "bond", 23},
}
var positiveNumber uint8 = 69
var address = "localhost:9000"

// func  index(w http.ResponseWriter, r *http.Request) {
//     fmt.Fprintln(w, "apa kabar!")
// }

func users(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    if r.Method == "GET" {
        var result, err = json.Marshal(data)

        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }

        w.Write(result)
        return
    }

    http.Error(w, "", http.StatusBadRequest)
}

func user(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    if r.Method == "GET" {
        var id = r.FormValue("id")
        var result []byte
        var err error

        for _, each := range data {
            if each.ID == id {
                result, err = json.Marshal(each)

                if err != nil {
                    http.Error(w, err.Error(), http.StatusInternalServerError)
                    return
                }

                w.Write(result)
                return
            }
        }

        http.Error(w, "User not found", http.StatusNotFound)
        return
    }

    http.Error(w, "", http.StatusBadRequest)
}


func main() {
	//hello world
	fmt.Println("hello rezani!")

	//tes varianle
	fmt.Println("bilangan positive : ", positiveNumber)

	//tes post & get method 
	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    //     switch r.Method {
    //     case "POST":
    //         w.Write([]byte("post"))
    //     case "GET":
    //         w.Write([]byte("get"))
    //     default:
    //         http.Error(w, "", http.StatusBadRequest)
    //     }
    // })



    //web server tes from template view.html
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        var data = map[string]string{
            "Name":    "Ahmad Ridwan Rezani",
            "Message": "have a nice day",
        }

        var t, err = template.ParseFiles("view.html")
        if err != nil {
            fmt.Println(err.Error())
            return
        }
        
        t.Execute(w, data)

    })

    
    
    // keluarin response dari func 
    //http.HandleFunc("/index", index)
    http.HandleFunc("/users", users) //api all get data
    http.HandleFunc("/user", user) // api 1 data


	//cek runnning web service
    fmt.Println("server started at localhost:9000")
    http.ListenAndServe(":9000", nil)
}