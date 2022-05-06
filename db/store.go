package db

type Store struct {
	db string
}

func NewStore(db string) Store {
	return Store{
		db: db,
	}
}
