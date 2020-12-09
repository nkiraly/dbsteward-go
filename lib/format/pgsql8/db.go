package pgsql8

var GlobalDb *Db = NewDB()

type Db struct {
}

// TODO(go,nth) should this be in lib?
type DbResult interface {
	RowCount() int
	Next() bool
	FetchRowStringMap() map[string]string // TODO(go,pgsql8) error handling
}

func NewDB() *Db {
	return &Db{}
}

func (self *Db) Connect(host string, port uint, name, user, pass string) {
	// TODO(go,pgsql)
}

func (self *Db) Query(sql string) DbResult {
	return nil
}