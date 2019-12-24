package model

import (
	"errors"
	"fmt"
	"time"
	"log"
)

type User struct {
	ID int
	Name string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone,omitempty"`
	Created time.Time `json:"created"`
}

var sampleUsers = []*User{
	&User{
		Name: "Joe",
		Email: "joec@gmail.com",
		Phone: "3476092075",
		Created: time.Now(),
	},
	&User{
		Name: "Mindy",
		Email: "mindyj@gmail.com",
		Phone: "3476092126",
		Created: time.Now(),
	},
	&User{
		Name: "Ava",
		Email: "ava.chasinga@gmail.com",
		Created: time.Now(),
	},
}

func createSampleUsers() {
	for _, user := range sampleUsers {
		if err := CreateUser(user); err != nil {
			log.Println(err)
			continue
		}
		
		log.Printf("user %s created\n", user.Email)
	}
}

func CreateUser(u *User) error {
	if userAlreadyExists(u) {
		return errors.New("User already existed")
	}

	_, err := db.Exec(`
INSERT INTO usr (name, email, phone, created)
VALUES ($1, $2, $3, $4)
;`, u.Name, u.Email, u.Phone, u.Created)
	if err != nil {
		return errors.New("Failed to create user")
	}
	return nil
}

func GetUser(u *User) (*User, error) {
	stmt := fmt.Sprintf(`
SELECT  id,
        name,
        email,
        phone,
        created
FROM    usr
WHERE   id = $1
OR      email = $2;`)
	if err := db.QueryRow(stmt, u.ID, u.Email).Scan(
		&u.ID, &u.Name, &u.Email, &u.Phone, &u.Created,
	); err != nil {
		return nil, err
	}
	return u, nil
}

func GetUsers(sorting string, ordering string, limit int) ([]*User, error) {
	users := []*User{}
	stmt := fmt.Sprintf(`
SELECT  usr.id,
        usr.name,
        usr.email,
        usr.phone,
        usr.created
FROM      usr
ORDER BY  %s %s
LIMIT     %d
;`, ordering, sorting, limit)
	rows, err := db.Query(stmt)
	if err != nil {
		return users, err
	}

	for rows.Next() {
		u := new(User)
		if err := rows.Scan(
			&u.ID,
			&u.Name,
			&u.Email,
			&u.Phone,
			&u.Created,
		); err != nil {
			return users, err
		}
		users = append(users, u)
	}

	if err := rows.Err(); err != nil {
		return users, err
	}

	return users, nil
}


func userAlreadyExists(u *User) bool {
	var exists bool
	err := db.QueryRow(`
SELECT EXISTS (
  SELECT  1
  FROM    usr
  WHERE   id = $1
  OR      email = $2
);`, u.ID, u.Email).Scan(&exists)
	if err != nil {
		log.Fatal(err)
	}
	return exists
}
