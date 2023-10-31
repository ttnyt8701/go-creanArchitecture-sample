package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// ミドルウェア

// ログ
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		start := time.Now()

		log.Printf("Started %s %s",r.Method,r.URL.Path)

		next.ServeHTTP(w,r)

		log.Printf("completed %s in %v", r.URL.Path, time.Since(start))
	})
}

func mainHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, World!")
}


func main(){
	http.Handle("/", loggingMiddleware(http.HandlerFunc(mainHandler)))

	log.Fatal(http.ListenAndServe(":8089", nil))

}