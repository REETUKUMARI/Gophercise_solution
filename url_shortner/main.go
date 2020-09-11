package main

import (
	"database/sql"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	mux := defaultMux()

	db, err := sql.Open("mysql", "reetu:Reetu8995@@(localhost:3306)/test")

	if err != nil {
		panic(err.Error())
	}

	var (
		path string
		url  string
	)
	var abc map[string]string
	abc = make(map[string]string)
	rows, err := db.Query("select * from db_url")
	if err != nil {
		log.Fatal(err)
	}
	rows.Scan()

	for rows.Next() {
		err := rows.Scan(&path, &url)

		if err != nil {
			log.Fatal(err)
		}

		abc[path] = url
		fmt.Println(abc)
	}

	yamlfile := flag.String("yaml", "url.yaml", "a yaml file in the formate of 'path, ulr'")
	flag.Parse()
	//_ = abc
	/*jsonfile := flag.String("json", "jsn.json", "a json file in the formate of 'path, ulr'")
	flag.Parse()

	jfile, err := os.Open(*jsonfile)

	if err != nil {
		log.Fatal(err)
	}
	dataj := make([]byte, 100000)
	jsn, err := jfile.Read(dataj)
	if err != nil {
		log.Fatal(err)
	}
	*/
	file, err := os.Open(*yamlfile)

	if err != nil {
		log.Fatal(err)
	}
	data := make([]byte, 100000)
	yaml, err := file.Read(data)
	if err != nil {
		log.Fatal(err)
	}

	jsn, err := ioutil.ReadFile("jsn.json")
	//yaml, err := r.readAll()
	if err != nil {
		log.Fatal(err)
	}

	// Build the MapHandler using the mux as the fallback
	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}
	mapHandler := MapHandler(pathsToUrls, mux)

	/*jsonfile, err := os.Open("jsn.json")
	if err != nil {
		fmt.Println(err)
	}
	datajsn := make([]byte, 100000)
	jsn, err := jsonfile.Read(datajsn)
	if err != nil {
		log.Fatal(err)
	}*/

	dbHandler, err := bdHandler(abc, mapHandler)

	jsnHandler, err := jsonHandler([]byte(string(jsn)), mapHandler)
	if err != nil {
		panic(err)
	}

	// Build the YAMLHandler using the mapHandler as the
	// fallback
	/*yaml := `
	- path: /urlshort
	url: https://github.com/gophercises/urlshort
	- path: /urlshort-final
	url: https://github.com/gophercises/urlshort/tree/solution
	`*/
	yamlHandler, err := YAMLHandler([]byte(string(data[:yaml])), mapHandler)
	if err != nil {
		panic(err)
	}

	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", jsnHandler)
	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", yamlHandler)
	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", dbHandler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}
