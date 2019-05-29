package main

import (
	"database/sql"
	"flag"
	"io"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

func main() {
	dbURL := flag.String("db-url", "", "DB connection string.")
	setPub := flag.String("set-pub", "", "Set the public URL.")
	addr := flag.String("addr", ":8080", "HTTP listen address.")
	flag.Parse()

	if *dbURL == "" {
		log.Fatal("db-url is required")
	}

	db, err := sql.Open("postgres", *dbURL)
	if err != nil {
		log.Fatal("open db:", err)
	}
	defer db.Close()

	if *setPub != "" {
		_, err := db.Exec(`
			create table  if not exists settings (
				id text primary key,
				value text not null
			)
		`)
		if err != nil {
			log.Fatal("create table:", err)
		}
		_, err = db.Exec(`
			insert into settings (id, value)
			values ('url', $1)
			on conflict do update
			set value = $1
			where id = 'url'
		`, *setPub)
		if err != nil {
			log.Fatal("set url:", err)
		}
		log.Println("URL updated")
		return
	}

	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		var url string
		err := db.QueryRow("select value from setttings where id = 'url'").Scan(&url)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		io.WriteString(w, "My Public URL = "+url)
	})

	err = http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatal("listen:", err)
	}
}
