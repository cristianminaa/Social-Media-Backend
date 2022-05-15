package database

import (
	"encoding/json"
	"errors"
	"os"
	"time"
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

// User
type User struct {
	CreatedAt time.Time `json:"createdAt"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Name      string    `json:"name"`
	Age       int       `json:"age"`
}

func (c Client) CreateUser(email, password, name string, age int) (User, error) {
	db, err := c.readDB()
	if err != nil {
		return User{}, err
	}
	if _, ok := db.Users[email]; ok {
		return User{}, errors.New("User already exists")
	}
	user := User{
		CreatedAt: time.Now().UTC(),
		Email:     email,
		Password:  password,
		Name:      name,
		Age:       age,
	}
	db.Users[email] = user
	err = c.updateDB(db)
	if err != nil {
		return User{}, err
	}
	return user, err
}

func (c Client) UpdateUser(email, password, name string, age int) (User, error) {
	db, err := c.readDB()
	if err != nil {
		return User{}, err
	}
	if _, ok := db.Users[email]; !ok {
		return User{}, errors.New("User does not exist")
	}
	user := db.Users[email]
	user.Password = password
	user.Name = name
	user.Age = age
	db.Users[email] = user
	err = c.updateDB(db)
	if err != nil {
		return User{}, err
	}
	return user, err
}

func (c Client) GetUser(email string) (User, error) {
	db, err := c.readDB()
	if err != nil {
		return User{}, err
	}
	user, ok := db.Users[email]
	if !ok {
		return User{}, errors.New("User does not exist")
	}
	return user, nil
}

func (c Client) DeleteUser(email string) error {
	db, err := c.readDB()
	if err != nil {
		return err
	}
	delete(db.Users, email)
	err = c.updateDB(db)
	if err != nil {
		return err
	}
	return nil
}

// Post
type Post struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UserEmail string    `json:"userEmail"`
	Text      string    `json:"text"`
}
