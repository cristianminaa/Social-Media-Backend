package database

import (
	"encoding/json"
	"os"
)

type Client struct {
	dbPath string
}

func NewClient(dbPath string) Client {
	return Client{
		dbPath: dbPath,
	}
}

func (c Client) updateDB(db databaseSchema) error {
	data, err := json.Marshal(db)
	if err != nil {
		return err
	}
	err = os.WriteFile(c.dbPath, data, 0600)
	if err != nil {
		return err
	}
	return nil
}

func (c Client) readDB() (databaseSchema, error) {
	data, err := os.ReadFile(c.dbPath)
	if err != nil {
		return databaseSchema{}, err
	}
	db := databaseSchema{}
	err = json.Unmarshal(data, &db)
	return db, err
}

func (c Client) createDB() error {
	data, err := json.Marshal(databaseSchema{
		Users: make(map[string]User),
		Posts: make(map[string]Post),
	})
	if err != nil {
		return err
	}
	err = os.WriteFile(c.dbPath, data, 0600)
	if err != nil {
		return err
	}
	return nil
}

type databaseSchema struct {
	Users map[string]User `json:"users"`
	Posts map[string]Post `json:"posts"`
}

func (c Client) EnsureDB() error {
	_, err := os.ReadFile(c.dbPath)
	if err != nil {
		return c.createDB()
	}
	return nil
}
