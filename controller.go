package watchr

import (
	"database/sql"
	"errors"

	"github.com/gchaincl/dotsql"
)

var (
	errNoRows = errors.New("ctl: no rows were affected")
)

// Controller operates on the database
type Controller struct {
	db  *sql.DB
	sql *dotsql.DotSql
}

// NewController opens the Sqlite database
func NewController() (*Controller, error) {
	db, err := sql.Open("sqlite3", "./watchr.db")
	if err != nil {
		return nil, err
	}

	dot, err := dotsql.LoadFromFile("queries.sql")

	return &Controller{
		db:  db,
		sql: dot,
	}, nil
}

// Close closes the database
func (c *Controller) Close() {
	c.db.Close()
}

// CreateUser creates a user
func (c *Controller) CreateUser(name, token string) error {
	res, err := c.sql.Exec(c.db, "create-user", name, token)
	if err != nil {
		return err
	}
	err = checkRowsAffected(res)
	if err != nil {
		return err
	}
	return nil
}

// ValidateUser will validate a token against a username
func (c *Controller) ValidateUser(name, token string) (bool, error) {
	row, err := c.sql.QueryRow(c.db, "validate-user", token, name)
	if err != nil {
		return false, err
	}

	var r int
	err = row.Scan(&r)
	if err != nil {
		return false, err
	}

	return r == 1, nil
}

// FindUserByToken finds a User from their token
func (c *Controller) FindUserByToken(token string) (*User, error) {
	row, err := c.sql.QueryRow(c.db, "find-user-by-token", token)
	if err != nil {
		return nil, err
	}

	u := new(User)
	err = row.Scan(&u.ID, &u.Name, &u.Email, &u.Level, &u.CreatedAt)
	if err != nil {
		return nil, err
	}
	return u, nil
}

// CreateRoom creates a room
func (c *Controller) CreateRoom(owner int, name string) error {
	res, err := c.sql.Exec(c.db, "create-room", owner, name)
	if err != nil {
		return err
	}
	err = checkRowsAffected(res)
	if err != nil {
		return err
	}
	return nil
}

// FindRoom finds a Room from its name
func (c *Controller) FindRoom(name string) (*Room, error) {
	row, err := c.sql.QueryRow(c.db, "find-room", name)
	if err != nil {
		return nil, err
	}

	r := new(Room)
	err = row.Scan(&r.OwnerID, &r.Name, &r.MediaType, &r.MediaSource, &r.CreatedAt, &r.ModifiedAt)
	if err != nil {
		return nil, err
	}
	return r, nil
}

// DeleteRoom removes a room forever
func (c *Controller) DeleteRoom(name string) error {
	res, err := c.sql.Exec(c.db, "delete-room", name)
	if err != nil {
		return err
	}
	err = checkRowsAffected(res)
	if err != nil {
		return err
	}
	return nil
}

func checkRowsAffected(res sql.Result) error {
	rows, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return errNoRows
	}
	return nil
}
