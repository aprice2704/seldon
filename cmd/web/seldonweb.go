package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/aprice2704/seldon/wbs"
	"github.com/gorilla/mux"
	"github.com/timshannon/bolthold"
)

func simpleLog(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.RequestURI)
		next.ServeHTTP(w, r)
	})
}

type simp struct {
	Name    string
	Created time.Time
}

const (
	version    = "0.9.0"
	localstore = "mylocal.seldon"
	port       = 8080
)

func main() {

	fmt.Printf("Seldon Web UI Server %s\n", version)

	// Some test data
	pieces := []wbs.Piece{ // Just a temp structure to read in values
		{Name: "Aldrin Base", Serial: 1000, ID: "A001", ParentID: ""},
		{Name: "Eagle Landing Pad", Serial: 1100, ID: "A002", ParentID: "A001"},
		{Name: "Foundation", Serial: 1200, ID: "A003", ParentID: "A002"},
		{Name: "Perimeter berm", Serial: 1300, ID: "A004", ParentID: "A002"},
		{Name: "Main Building", Serial: 1050, ID: "A005", ParentID: "A001"},
	}
	wbs := wbs.NewWBS("A001", pieces)

	fmt.Printf("WBS:\n%s", wbs.String())

	fmt.Printf("Opening DB %s\n", localstore)
	store, err := bolthold.Open(localstore, 0666, nil)
	if err != nil {
		panic(fmt.Sprintf("Cannot open mylocal db :( %v", err.Error()))
	}

	err = store.Upsert("MyAwesomeWBS", wbs.Pieces)
	if err != nil {
		panic(fmt.Sprintf("Couldn't upsert %s to %s, err %s", "MyAwesomeWBS", localstore, err.Error()))
	}

	Router := mux.NewRouter()
	Router.HandleFunc("/", rootHandler)

	web := &http.Server{
		Handler:      Router,
		Addr:         ":" + strconv.Itoa(port),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	fmt.Printf("Listening on %s\n", web.Addr)
	web.ListenAndServe()

}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Agrippina is the world's cutest kitty!\n")
}
