package postgredb_test

import (
	"context"
	"log"
	"regexp"
	"testing"
	"time"

	models_rep "GoWeb/models/repository"
	"GoWeb/repository/postgredb"
	"GoWeb/utils"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/go-test/deep"
	"github.com/golang/mock/gomock"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/xorcare/pointer"
)

func TestMemberRepo_Insert(t *testing.T) {
	t.Parallel()
	mockCtl := gomock.NewController(t)
	defer mockCtl.Finish()
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

func TestMemberRepo_Find(t *testing.T) {
	t.Parallel()
	mockCtl := gomock.NewController(t)
	defer mockCtl.Finish()
	db, mock := newMemberMock()
	repository := postgredb.NewMemberRepo(db)
	defer func() {
		repository.Close()
	}()
	want := createRandomMemberAccount()
	query := `SELECT account,password,permission,name,email,phone,address,create_at,update_at FROM member` +
		` WHERE is_alive = true AND account =$1 OR phone = $2 ORDER BY id`
	mock.ExpectQuery(regexp.QuoteMeta(query)).
		WithArgs(want.Account, want.Phone).
		WillReturnRows(sqlmock.NewRows([]string{"account,", "password", "permission", "name", "email", "phone", "address", "create_at", "update_at"}).
			AddRow(want.Account, want.Password, want.Permission, want.Name, want.Email, want.Phone, want.Address, want.CreateAt, want.UpdateAt))
	got, err := repository.Find(context.Background(), &want.Account, want.Phone)
	require.NoError(t, err)
	require.Nil(t, deep.Equal(want, got))
}

func TestMemberRepo_FindAll(t *testing.T) {
	t.Parallel()
	mockCtl := gomock.NewController(t)
	defer mockCtl.Finish()
	db, mock := newMemberMock()
	repository := postgredb.NewMemberRepo(db)
	defer func() {
		repository.Close()
	}()
	want := []models_rep.Member{}
	want = append(want, *createRandomMemberAccount())
	want = append(want, *createRandomMemberAccount())
	query := `SELECT account,password,permission,name,email,phone,address,create_at,update_at
	FROM member
	WHERE is_alive = true`
	mock.ExpectQuery(regexp.QuoteMeta(query)).
		WillReturnRows(sqlmock.NewRows([]string{"account,", "password", "permission", "name", "email", "phone", "address", "create_at", "update_at"}).
			AddRow(want[0].Account, want[0].Password, want[0].Permission, want[0].Name, want[0].Email, want[0].Phone, want[0].Address, want[0].CreateAt, want[0].UpdateAt).
			AddRow(want[1].Account, want[1].Password, want[1].Permission, want[1].Name, want[1].Email, want[1].Phone, want[1].Address, want[1].CreateAt, want[1].UpdateAt))
	got, err := repository.FindAll(context.Background())
	require.NoError(t, err)
	require.Nil(t, deep.Equal(&want, got))
}

func TestMemberRepo_Updates(t *testing.T) {
	t.Parallel()
	mockCtl := gomock.NewController(t)
	defer mockCtl.Finish()
	db, mock := newMemberMock()
	repository := postgredb.NewMemberRepo(db)
	defer func() {
		repository.Close()
	}()
	want := createRandomUpdateMemberAccount()
	mock.ExpectExec(regexp.QuoteMeta(`UPDATE member SET name =$2, email =$3,phone=$4, address=$5, update_at=$6 WHERE account = $1`)).
		WithArgs(want.Account, want.Name, want.Email, want.Phone, want.Address, time.Now().Format(time.RFC3339)).
		WillReturnResult(sqlmock.NewResult(1, 1))
	err := repository.Updates(context.Background(), want)
	require.NoError(t, err)
}

func TestMemberRepo_Disable(t *testing.T) {
	t.Parallel()
	mockCtl := gomock.NewController(t)
	defer mockCtl.Finish()
	db, mock := newMemberMock()
	repository := postgredb.NewMemberRepo(db)
	defer func() {
		repository.Close()
	}()
	want := createRandomUpdateMemberAccount()
	mock.ExpectExec(regexp.QuoteMeta(`UPDATE member SET is_alive = false WHERE account = $1`)).
		WithArgs(want.Account).
		WillReturnResult(sqlmock.NewResult(1, 1))
	err := repository.Disable(context.Background(), want)
	require.NoError(t, err)
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
		Permission: 2,
		Name:       utils.Rand(10, utils.RAND_KIND_ALL),
		Email:      pointer.String(utils.RandomEmail()),
		Phone:      pointer.String("09" + utils.RandomNum(8)),
		Address:    pointer.String(utils.Rand(15, utils.RAND_KIND_ALL)),
		CreateAt:   pointer.Time(time.Now()),
		UpdateAt:   pointer.Time(time.Now()),
	}
}
func createRandomUpdateMemberAccount() *models_rep.UpdateMember {
	return &models_rep.UpdateMember{
		Account: "test_" + utils.Rand(10, utils.RAND_KIND_ALL),
		Name:    pointer.String(utils.Rand(10, utils.RAND_KIND_ALL)),
		Email:   pointer.String(utils.RandomEmail()),
		Phone:   pointer.String("09" + utils.RandomNum(8)),
		Address: pointer.String(utils.Rand(15, utils.RAND_KIND_ALL)),
	}
}

// endregion
