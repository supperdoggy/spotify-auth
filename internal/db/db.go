package db

import (
	"errors"
	"github.com/supperdoggy/spotify-web-project/spotify-auth/internal/utils"
	globalStructs "github.com/supperdoggy/spotify-web-project/spotify-globalStructs"
	"go.uber.org/zap"
	"gopkg.in/mgo.v2"
	"gopkg.in/night-codes/types.v1"
	"math/rand"
	"sync"
	"time"
)

type obj map[string]interface{}

type tokenCache struct {
	m map[string]globalStructs.Token
	mut sync.Mutex
}

type IDB interface {
	NewToken(userID string) (string, error)
	CheckToken(token string) (bool, string)
	NewCreds(creds globalStructs.Creds) error
	GetCredsByEmail(email string) (result globalStructs.Creds, err error)
}

type DB struct {
	Session *mgo.Session

	CredsCollection *mgo.Collection
	cache tokenCache
	logger *zap.Logger
}

const ValidFor = 30

func NewDB(dbName string, logger *zap.Logger) (IDB, error) {
	sess, err := mgo.Dial("")
	if err != nil {
		return nil, err
	}

	return &DB{
		Session:         sess,
		CredsCollection: sess.DB(dbName).C("Creds"),
		cache:           tokenCache{
			m: map[string]globalStructs.Token{},
		},
		logger: logger,
	}, nil
}

func (d *DB) NewToken(userID string) (string, error) {
	if userID == "" {
		return "", errors.New("id cannot be empty")
	}
	token := utils.GenerateToken()
	d.cache.mut.Lock()
	d.cache.m[token] = globalStructs.Token{
		UserID:   userID,
		TokenStr: token,
		// valid for 30 days
		Expire:   time.Now().Add(ValidFor*24*time.Hour),
	}
	d.cache.mut.Unlock()
	return token, nil
}

func (d *DB) CheckToken(token string) (bool, string) {
	d.cache.mut.Lock()
	v, ok := d.cache.m[token]
	d.cache.mut.Unlock()
	if !ok {
		return false, ""
	}

	if time.Now().After(v.Expire) {
		d.cache.mut.Lock()
		delete(d.cache.m, token)
		d.cache.mut.Unlock()

		return false, ""
	}
	return true, v.UserID
}

func (d *DB) NewCreds(creds globalStructs.Creds) error {
	// check if exists
	var prev globalStructs.Creds
	err := d.CredsCollection.Find(obj{"email": creds.Email}).One(&prev)
	if err != nil && err != mgo.ErrNotFound {
		return err
	} else if err == nil {
		return errors.New("email already exists")
	}

	rand.Seed(time.Now().UnixNano())
	creds.ID = types.String(rand.Int())
	return d.CredsCollection.Insert(creds)
}

func (d *DB) GetCredsByEmail(email string) (result globalStructs.Creds, err error) {
	err = d.CredsCollection.Find(obj{"email": email}).One(&result)
	return
}
