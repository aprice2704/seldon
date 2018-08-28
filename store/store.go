package store

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"time"

	"github.com/boltdb/bolt"
)

// Storage is a thing capable of storing seldon projects
type Storage interface {
	Initialize()
	Close()
	SetBucket(bucket string)
	Get(key string) (buf bytes.Buffer)
	Put(key string, buf bytes.Buffer)
	//Encode(thing Storable)
	//Decoder(buf bytes.Buffer) (thing Storable)
}

// Storable is a thing that can put itself in Storage
type Storable interface {
	Get(store Storage, key string) (value Storable, err error)
	Put(store Storage) (value Storable, err error)
}

// Bolt store -------------------------

const (
	projectsBucketName = "Seldon-Projects" // the top level bucket where we store the names of the projects in this db
)

// BoltStore stores seldon projects in BoltDB
type BoltStore struct {
	filename      string
	db            *bolt.DB
	currentbucket *bolt.Bucket
	projsBucket   *bolt.Bucket
	//enc            *gob.Encoder
	//dec            *gob.Decoder
	//encbuf, decbuf *bytes.Buffer
}

// SetBucket set the current bucket (named for project in Bolt)
func (bs *BoltStore) SetBucket(bucket string) {
	bs.db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte(bucket))
		if err != nil {
			return fmt.Errorf("Could not create bucket %s: %s", bucket, err)
		}
		bs.currentbucket = b
		return nil
	})
}

// NewBolt makes a Bolt store and call Initialize to set up its insides
func NewBolt(filename string) (bs *BoltStore, err error) {
	bs = new(BoltStore)
	bs.filename = filename
	fmt.Printf("Opening DB %s\n", filename)
	db, err := bolt.Open(filename, 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		return nil, fmt.Errorf("Cannot open %s local db :( %v", filename, err.Error())
	}
	bs.filename = filename
	bs.db = db
	err = bs.Initialize()
	//bs.enc = gob.NewEncoder(bs.encbuf)
	//bs.dec = gob.NewDecoder(bs.decbuf)
	return bs, err
}

// Initialize just sets up internal db structures
func (bs *BoltStore) Initialize() (err error) {
	bs.db.Update(func(tx *bolt.Tx) error {
		bu, err2 := tx.CreateBucketIfNotExists([]byte(projectsBucketName))
		if err2 != nil {
			return fmt.Errorf("Create bucket: %s", err2)
		}
		bs.projsBucket = bu
		return nil
	})
	return nil
}

// Close a bolt store
func (bs BoltStore) Close() {
	bs.db.Close()
	bs.db = nil
	bs.currentbucket = nil
	bs.filename = ""
}

// itob returns an 8-byte big endian representation of v.
func itob(v int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}
