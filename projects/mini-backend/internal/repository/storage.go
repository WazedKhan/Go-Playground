package repository

type StorageSelector interface {
	Get(key string) (*string, error)
	Save(key, value string) (bool, error)
}

type jsonStorage struct{}

func NewJsonStore() StorageSelector {
	return &jsonStorage{}
}

type sqliteStorage struct{}

func NewSqliteStorage() StorageSelector {
	return &sqliteStorage{}
}

type postgresStorage struct {}

func NewPostgresStorage() StorageSelector {
	return &postgresStorage{}
}
