package sql

import (
	"Multimedia_Processing_Pipeline/constant"
	"errors"
	"github.com/syndtr/goleveldb/leveldb"
	"testing"
)

// go test -v -run TestCreate

func TestCreate(t *testing.T) {

}

func TestAddOne(t *testing.T) {

}

func TestGetLevelDB(t *testing.T) {
	p := &constant.Param{
		Root:     "C:\\Users\\zen\\Github\\Multimedia_Processing_Pipeline\\sql",
		Language: "English",
		Pattern:  "mp4",
		Model:    "base",
		Location: "C:\\Users\\zen\\Github\\Multimedia_Processing_Pipeline\\sql",
		Proxy:    "192.168.1.20:8889",
	}
	SetLevelDB(p)
	err := GetLevelDB().Put([]byte("key"), []byte("value"), nil)
	if err != nil {
		t.Log(err)
	}
	get, err := GetLevelDB().Get([]byte("key"), nil)
	if err != nil {
		t.Log(err)
	}
	t.Log(string(get))
	no, err := GetLevelDB().Get([]byte("key2"), nil)
	if err != nil {
		//err = "leveldb: not found"
		notFound := errors.New("leveldb: not found")
		errors.Is(err, leveldb.ErrNotFound)
		if err.Error() == notFound.Error() {
			t.Log("错误相等")
		} else {
			t.Log(err)
		}

	}
	t.Log(string(no), err)
}
func TestGetALL(t *testing.T) {
	p := &constant.Param{
		Root:     "C:\\Users\\zen\\Github\\Multimedia_Processing_Pipeline\\sql",
		Language: "English",
		Pattern:  "mp4",
		Model:    "base",
		Location: "C:\\Users\\zen\\Github\\Multimedia_Processing_Pipeline\\sql",
		Proxy:    "192.168.1.20:8889",
	}
	SetLevelDB(p)
	iter := GetLevelDB().NewIterator(nil, nil)
	for iter.Next() {
		// Remember that the contents of the returned slice should not be modified, and
		// only valid until the next call to Next.
		key := iter.Key()
		value := iter.Value()
		t.Log(key, value)
	}
	iter.Release()
	err := iter.Error()
	if err != nil {
		t.Log(err)
	}
}
