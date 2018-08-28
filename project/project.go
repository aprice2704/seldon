package project

import (
	"encoding/gob"

	"github.com/aprice2704/seldon/store"
	"github.com/aprice2704/seldon/wbs"
)

// Project is the top level object representing an entire seldon project
type Project struct {
	ID        int
	Name      string
	Evergreen bool
	WBS       wbs.WBS
}

// Get for Storable
func (p Project) Get(st store.Storage, key string) (err error) {
	buf := st.Get(key)
	//var buf bytes.Buffer
	dec := gob.NewDecoder(&buf)
	err = dec.Decode(&p)
	return err
}

// Put for Storable
func (p Project) Put(st store.Storage) (value Storable, err error) {

}

func NewProject(n string) {
	return &Project{Name: n}
}

// // SaveProject stores an entire project in a BoltDB
// func (b *BoltStore) SaveProject(proj project.Project) (err error) {
//
// 	b.Initialize() // just in case
//
// 	bu := tx.Bucket([]byte(projectsBucket))
//
// 	if proj.ID = 0 {
// 		proj.ID, _ := bu.NextSequence()
// 	}
//
// 		err = b.db.Update(func(tx *bolt.Tx) error {
// 			var gobuf byte.Buffer
// 			enc := gob.NewEncoder(&gobuf)
// 			err := enc.Encode(proj)
// 			return bu.Put(itob(id), gobuf.Bytes())
// 		})
//
// 	return err
// }

// GetProjects gets the names of the projects in this store
// func (b BoltStore) GetProjects() (projects []string, err error) {
// 	ps := make([]string, 0)
// 	err = b.db.View(func(tx *bolt.Tx) error {
// 		bu := b.db.Bucket(projectsBucketName)
// 		c := bu.Cursor()
// 		for k, v := c.First(); k != nil; k, v = c.Next() {
// 			fmt.Printf("key=%s, value=%s\n", k, v)
// 			ps = append(ps, string(v))
// 		}
// 		return nil
// 	})
// 	if err != nil {
// 		return nil, err
// 	}
// 	return ps, nil
// }
