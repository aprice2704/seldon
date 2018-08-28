package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/aprice2704/seldon/project"
	"github.com/aprice2704/seldon/store"
	"github.com/aprice2704/seldon/wbs"
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
	projname   = "Moonbase"
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

	fmt.Printf("Storing WBS:\n%s", mywbs.String())

	mystore, err := store.NewBolt(localstore)
	if err != nil {
		log.Fatalf("Trying to open %s, %s", localstore, err.Error())
	}
	defer mystore.Close()

	p := project.NewProject(projname)
	err := mystore.SaveProject(p)

	// var mygob bytes.Buffer
	// enc := gob.NewEncoder(&mygob)
	// //dec := gob.NewDecoder(&buf)

	// err := db.Update(func(tx *bolt.Tx) error {
	// 	for _, p := range mywbs.Pieces {
	// 		b, err := tx.CreateBucketIfNotExists([]byte(project))
	// 		if err != nil {
	// 			return fmt.Errorf("Couldn't create bucket: %s", err)
	// 		}
	// 		k := []byte(p.Key())
	// 		err = enc.Encode(p)
	// 		if err != nil {
	// 			return fmt.Errorf("Couldn't gob encode %s, err %s", &p, err.Error())
	// 		}
	// 		err = b.Put(k, mygob.Bytes())
	// 		if err != nil {
	// 			return fmt.Errorf("Couldn't save %s to %s, err %s", &p, localstore, err.Error())
	// 		}
	// 	}
	// 	return nil
	// })

	// Router := mux.NewRouter()
	// Router.HandleFunc("/", rootHandler)
	//
	// web := &http.Server{
	// 	Handler:      Router,
	// 	Addr:         ":" + strconv.Itoa(port),
	// 	WriteTimeout: 15 * time.Second,
	// 	ReadTimeout:  15 * time.Second,
	// }
	// fmt.Printf("Listening on %s\n", web.Addr)
	// web.ListenAndServe()

}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Agrippina is the world's cutest kitty!\n")
}
