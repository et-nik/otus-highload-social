package di

import (
	"database/sql"

	"github.com/matthewhartstonge/argon2"

	"github.com/et-nik/otus-highload/internal/di/config"
	"github.com/et-nik/otus-highload/internal/domain"
	"github.com/et-nik/otus-highload/internal/repositories"
)

type Container struct {
	config     *config.Config
	connection *sql.DB

	argon *argon2.Config

	userRepository domain.UserRepository
}

func NewContainer(config *config.Config) *Container {
	return &Container{config: config}
}

func (c *Container) UserRepository() domain.UserRepository {
	if c.userRepository == nil {
		repository, err := c.createUserRepository()
		if err != nil {
			panic(err)
		}

		c.userRepository = repository
	}

	return c.userRepository
}

func (c *Container) Connection() *sql.DB {
	if c.connection == nil {
		connection, err := c.createConnection()
		if err != nil {
			panic(err)
		}

		c.connection = connection
	}

	return c.connection
}

func (c *Container) createConnection() (*sql.DB, error) {
	return sql.Open("mysql", c.config.Database)
}

func (c *Container) createUserRepository() (domain.UserRepository, error) {
	repository := repositories.NewUserRepository(c.Connection())

	return repository, nil
}

func (c *Container) Argon() *argon2.Config {
	if c.argon == nil {
		argon, err := c.createArgon()
		if err != nil {
			panic(err)
		}

		c.argon = argon
	}

	return c.argon
}

func (c *Container) createArgon() (*argon2.Config, error) {
	argon := argon2.DefaultConfig()

	return &argon, nil
}
