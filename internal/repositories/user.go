package repositories

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/et-nik/otus-highload/internal/domain"
	"github.com/pkg/errors"
)

const queryFind = `
	SELECT users.id, users.auth_token_hash, users.age, users.email, users.password, users.name, users.surname, users.sex, 
	       users.city, users.interests, group_concat(users_friends.target_id) as friend_ids FROM users
	LEFT JOIN users_friends ON users_friends.source_id = users.id
	GROUP BY users.id;
`

const queryFindByID = `
	SELECT users.id, users.auth_token_hash, users.age, users.email, users.password, users.name, users.surname, users.sex, 
	       users.city, users.interests, group_concat(users_friends.target_id) as friend_ids FROM users
		LEFT JOIN users_friends ON users_friends.source_id = users.id
		WHERE users.id = ?
		GROUP BY users.id;

`

const queryFindByEmail = `
	SELECT users.id, users.auth_token_hash, users.age, users.email, users.password, users.name, users.surname, users.sex, 
	       users.city, users.interests, group_concat(users_friends.target_id) as friend_ids FROM users
		LEFT JOIN users_friends ON users_friends.source_id = users.id
		WHERE users.email = ?
		GROUP BY users.id;
`

type UserRepository struct {
	connection *sql.DB
}

func NewUserRepository(connection *sql.DB) *UserRepository {
	return &UserRepository{connection: connection}
}

func (repository *UserRepository) Find(ctx context.Context) ([]*domain.User, error) {
	rows, err := repository.connection.QueryContext(ctx, queryFind)
	if err != nil {
		return nil, err
	}
	if rows.Err() != nil {
		return nil, rows.Err()
	}

	var users []*domain.User
	for rows.Next() {
		user, err := repository.scan(rows)
		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

// FindByEmail TODO: Replace to Find with criteria.
func (repository *UserRepository) FindByEmail(ctx context.Context, email string) (*domain.User, error) {
	row := repository.connection.QueryRowContext(ctx, queryFindByEmail, email)

	return repository.scan(row)
}

func (repository *UserRepository) FindByID(ctx context.Context, id int) (*domain.User, error) {
	row := repository.connection.QueryRowContext(ctx, queryFindByID, id)

	return repository.scan(row)
}

func (repository *UserRepository) Save(ctx context.Context, user *domain.User) error {
	if user.ID == 0 {
		return repository.insert(ctx, user)
	}

	return repository.update(ctx, user)
}

const queryUpdate = `
	UPDATE users SET age=?, auth_token_hash=?, email=?, password=?, name=?, surname=?, sex=?, city=?, interests=?
		WHERE id=?
`

func (repository *UserRepository) update(ctx context.Context, user *domain.User) error {
	stmt, err := repository.connection.Prepare(queryUpdate)
	if err != nil {
		return err
	}

	interests, err := json.Marshal(user.Interests)
	if err != nil {
		return err
	}

	_, err = stmt.ExecContext(
		ctx,
		user.Age,
		user.AuthTokenHash,
		user.Email,
		user.PasswordHash,
		user.Name,
		user.Surname,
		user.Sex,
		user.City,
		interests,
		user.ID,
	)
	if err != nil {
		return err
	}

	return repository.updateFriends(ctx, user)
}

const queryInsert = `
	INSERT INTO users(age, auth_token_hash, email, password, name, surname, sex, city, interests)
		VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?)
`

func (repository *UserRepository) insert(ctx context.Context, user *domain.User) error {
	stmt, err := repository.connection.Prepare(queryInsert)
	if err != nil {
		return err
	}

	interests, err := json.Marshal(user.Interests)
	if err != nil {
		return err
	}

	result, err := stmt.ExecContext(
		ctx,
		user.Age,
		user.AuthTokenHash,
		user.Email,
		user.PasswordHash,
		user.Name,
		user.Surname,
		user.Sex,
		user.City,
		interests,
	)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	user.ID = int(id)

	return repository.updateFriends(ctx, user)
}

const queryFriends = `
	REPLACE INTO users_friends(source_id, target_id) VALUES %s
`

func (repository *UserRepository) updateFriends(ctx context.Context, user *domain.User) error {
	var vals []interface{}
	valuesStr := ""

	for _, friendID := range user.Friends {
		valuesStr += "(?, ?),"
		vals = append(vals, user.ID, friendID)
	}

	valuesStr = valuesStr[0 : len(valuesStr)-1]
	sqlStr := fmt.Sprintf(queryFriends, valuesStr)

	stmt, err := repository.connection.Prepare(sqlStr)
	if err != nil {
		return err
	}

	_, err = stmt.ExecContext(ctx, vals...)
	if err != nil {
		return err
	}

	return nil
}

type scanner interface {
	Scan(dest ...interface{}) error
}

//nolint:funlen
func (repository *UserRepository) scan(row scanner) (*domain.User, error) {
	var userData struct {
		ID            int
		AuthTokenHash *string
		Email         string
		Password      string
		Name          string
		Surname       string
		Age           int
		Sex           string
		Interests     []byte
		City          string
		Friends       *string
	}
	err := row.Scan(
		&userData.ID,
		&userData.AuthTokenHash,
		&userData.Age,
		&userData.Email,
		&userData.Password,
		&userData.Name,
		&userData.Surname,
		&userData.Sex,
		&userData.City,
		&userData.Interests,
		&userData.Friends,
	)
	if err != nil && errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	var interests []string
	err = json.Unmarshal(userData.Interests, &interests)
	if err != nil {
		return nil, err
	}

	authTokenHash := ""
	if userData.AuthTokenHash != nil {
		authTokenHash = *userData.AuthTokenHash
	}

	var friends []int
	if userData.Friends != nil {
		friendsRawList := strings.Split(*userData.Friends, ",")
		for _, f := range friendsRawList {
			friendID, _ := strconv.Atoi(f)
			friends = append(friends, friendID)
		}
	}

	return &domain.User{
		ID:            userData.ID,
		Email:         userData.Email,
		AuthTokenHash: authTokenHash,
		PasswordHash:  userData.Password,
		Name:          userData.Name,
		Surname:       userData.Surname,
		Age:           userData.Age,
		Sex:           userData.Sex,
		Interests:     interests,
		City:          userData.City,
		Friends:       friends,
	}, err
}
