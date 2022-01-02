package sqlstore_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/triumphpc/go-http-rest-api/internal/app/model"
	"github.com/triumphpc/go-http-rest-api/internal/app/storage/sqlstore"
	"testing"
)

func TestUserRepository_Create(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("users")

	s := sqlstore.New(db)
	u := model.TestUser(t)
	assert.NoError(t, s.User().Create(u))
	//assert.NotNil(t, u)
}

func TestUserRepository_FindByEmail(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("users")

	s := sqlstore.New(db)
	example := "user@example.org"
	_, err := s.User().FindByEmail(example)

	assert.Error(t, err)

	usr := model.TestUser(t)
	usr.Email = example

	err = s.User().Create(usr)

	assert.NoError(t, err)

	_, err = s.User().FindByEmail(example)
	assert.NoError(t, err)
	assert.NotNil(t, usr)
	assert.Equal(t, usr.Email, example)

}
