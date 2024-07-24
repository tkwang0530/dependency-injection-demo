package db

import "sync"

type InMemoryDB struct {
	sync.Mutex
	todos  map[int]string
	nextID int
}

func NewInMemoryDB() *InMemoryDB {
	return &InMemoryDB{
		todos:  make(map[int]string),
		nextID: 1,
	}
}

func (db *InMemoryDB) Add(todo string) int {
	db.Lock()
	defer db.Unlock()
	id := db.nextID
	db.todos[id] = todo
	db.nextID++
	return id
}

func (db *InMemoryDB) GetAll() map[int]string {
	db.Lock()
	defer db.Unlock()
	return db.todos
}

func (db *InMemoryDB) Delete(id int) {
	db.Lock()
	defer db.Unlock()
	delete(db.todos, id)
}
