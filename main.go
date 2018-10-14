package main

import "encoding/json"
import "net/http"
import "fmt"
import "strconv"

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
	student{"B003", "James", 27},
	student{"B004", "Richard", 30},
}

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
		keys, id := r.URL.Query()["id"]
		var result []byte
		var err error

		if id {
			for _, each := range data {
				if each.ID == keys[0] {
					result, err = json.Marshal(each)

					if err != nil {
						http.Error(w, err.Error(), http.StatusInternalServerError)
						return
					}

					w.Write(result)
					return
				}
			}
		}

		http.Error(w, "User not found", http.StatusBadRequest)
		return
	}

	http.Error(w, "", http.StatusBadRequest)
}

func addUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method == "GET" {
		pId := r.FormValue("id")
		pName := r.FormValue("name")
		pGrade, err_1 := strconv.Atoi(r.FormValue("grade"))

		if err_1 != nil {
			http.Error(w, err_1.Error(), http.StatusInternalServerError)
			return
		}

		data = append(data, student{pId, pName, pGrade})

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

func mainroot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Try <ol><li>http://<host>/user/list</li><li>http://<host>user/get-by-id?id=<id></li><li>http://<host>user/add?id=<id>&name=<name>&grade=<grade></li></ol>")
}

func main() {
	http.HandleFunc("/user/list", users)
	http.HandleFunc("/user/get-by-id", user)
	http.HandleFunc("/user/add", addUser)
	http.HandleFunc("/", mainroot)

	fmt.Println("Starting web server at http://localhost:8080/")
	http.ListenAndServe(":8080", nil)
}
