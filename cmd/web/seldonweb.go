package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/aprice2704/seldon/wbs"
	"github.com/boltdb/bolt"
	"github.com/gorilla/mux"
)

const (
	project = "Moonbase"
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
		{Name: "Aldrin Base", ID: "A001", ParentID: ""},
		{Name: "Eagle Landing Pad", ID: "A002", ParentID: "A001"},
		{Name: "Foundation", ID: "A003", ParentID: "A002"},
		{Name: "Perimeter berm", ID: "A004", ParentID: "A002"},
		{Name: "Main Building", ID: "A005", ParentID: "A001"},
	}
	mywbs := wbs.NewWBS("Aldrin", "A001", pieces)

	fmt.Printf("Storing WBS:\n%s", wbs.String())

	fmt.Printf("Opening DB %s\n", localstore)
	db, err := bolt.Open(localstore, 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		panic(fmt.Sprintf("Cannot open mylocal db :( %v", err.Error()))
	}
	defer db.Close()

	err := db.Update(func(tx *bolt.Tx) error {
		for _, p := range wbs.Pieces {
			if err != nil {
				panic(fmt.Sprintf("Couldn't save %s to %s, err %s", p, localstore, err.Error()))
			}
		}
		return nil
	})

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
