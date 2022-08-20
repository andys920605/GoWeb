package postgredb_test

import (
	"context"
	"log"
	"regexp"
	"testing"

	models_rep "GoWeb/models/repository"
	"GoWeb/repository/postgredb"
	"GoWeb/utils"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/golang/mock/gomock"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
	"github.com/xorcare/pointer"
)

func TestMemberRepo_Insert(t *testing.T) {
	t.Parallel()
	mockCtl := gomock.NewController(t)
	defer mockCtl.Finish()
	//mockLogger := mock.NewMockIApiLogger(mockCtl)
	db, mock := newMemberMock()
	repository := postgredb.NewMemberRepo(db)
	defer func() {
		repository.Close()
	}()
	want := createRandomMemberAccount()
	query := `INSERT INTO member(account, password, permission, name, email, phone, address) VALUES ($1,$2,$3,$4,$5,$6,$7);`
	mock.ExpectExec(regexp.QuoteMeta(query)).
		WithArgs(want.Account, want.Password, want.Permission, want.Name, want.Email, want.Phone, want.Address).
		WillReturnResult(sqlmock.NewResult(1, 1))
	err := repository.Insert(context.Background(), want)
	assert.NoError(t, err)
}

// region private methods
func newMemberMock() (*gorm.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New() // mock sql.DB
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	gdb, err := gorm.Open("postgres", db) // open gorm db
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	gdb.LogMode(true)
	return gdb, mock
}

func createRandomMemberAccount() *models_rep.Member {
	return &models_rep.Member{
		Account:    "test_" + utils.Rand(10, utils.RAND_KIND_ALL),
		Password:   utils.Rand(10, utils.RAND_KIND_ALL),
		Permission: utils.GenerateRangeNum(2, 4),
		Name:       utils.Rand(10, utils.RAND_KIND_ALL),
		Email:      pointer.String(utils.RandomEmail()),
		Phone:      pointer.String("09" + utils.RandomNum(8)),
		Address:    pointer.String(utils.Rand(15, utils.RAND_KIND_ALL)),
	}
}

// endregion
