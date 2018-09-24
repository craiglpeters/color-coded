package main
 
import (
	"os"
	"fmt"
	"net/http"
	"rsc.io/quote"
)

func main() {
	c := os.Getenv("COLOR")
	if len(c) == 0{
		os.Setenv("COLOR", "#F1A94E") //Blue 44B3C2 and Yellow F1A94E. 
	} 

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "<html onclick=\"window.location.href = '/die'\" style='background:" + os.Getenv("COLOR") + "'> Requested: %s\n" + fmt.Println(quote.Hello()) + " </html>", r.URL.Path)
	})

	http.HandleFunc("/dashboard", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "<html> DASHBOARD Requested: %s\n </html>", r.URL.Path)
	})

	http.HandleFunc("/die", func(w http.ResponseWriter, r *http.Request) {
		die();
	})

	http.ListenAndServe(":8080", nil)
}

func die() {
	os.Exit(3)
}
