// | melun/reaganiwadha/Regan Iwadha Google Code In Submission |

package main

import (
	"net/http"
	"strings"
)

var pathMap = make(map[string]string)

func redirect(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	path = strings.TrimPrefix(path, "/")
	if path == "" {
		w.Write([]byte("Welcome to my link shortener!, type some predefined path for redirections"))
	} else if pathMap[path] == "" {
		w.Write([]byte("Well that sucks... But there is no predefined path for that"))
	} else {
		http.Redirect(w, r, pathMap[path], 301)
	}
	print(path)
}

func main() {
	//Declare your maps here
	pathMap["google"] = "https://google.com"
	pathMap["go"] = "https://golang.org"
	pathMap["github"] = "https://github.com/reaganiwadha"
	pathMap["gci"] = "https://codein.withgoogle.com/"
	pathMap["pl"] = "https://pl.wikipedia.org/wiki/"

	http.HandleFunc("/", redirect)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
