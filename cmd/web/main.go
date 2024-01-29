package main

import (
	// "fmt"
	"flag"
	"log"
	"net/http"
	"os"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

const (
	p_infoLog string = "./tmp/info.log"
)

type application struct {
	infoLog *log.Logger
	errorLog *log.Logger
}

func main() {
	addr := flag.String("addr", ":4000", "HTTP network address")
	dsn := flag.String("dsn", "web:pass@/snippetbox?parseTime=true", "MySQL data source name")

	flag.Parse()

	f, err := os.OpenFile(p_infoLog, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil{
		log.Fatal(err)
	}
	defer f.Close()

	infoLog := log.New(f, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Llongfile)

	app := application{infoLog, errorLog}

	db, err := openDB(*dsn)
	if err != nil{
		errorLog.Fatal(err)
	}

	defer db.Close()
	
	srv := &http.Server{
		Addr: *addr,
		Handler: app.routes(),
		ErrorLog: errorLog,
	}

	infoLog.Println("Starting server on", *addr)
	err = srv.ListenAndServe()
	errorLog.Fatal(err)
}


func openDB(dsn string) (*sql.DB, error){
	db, err := sql.Open("mysql", dsn)
	if err != nil{
		return nil, err 
	}
	
	if err = db.Ping(); err != nil{
		return nil, err
	}
	return db, nil
}