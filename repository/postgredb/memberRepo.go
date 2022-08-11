package repository

import (
	models_rep "GoWeb/models/repository"
	"context"
	"database/sql"
	"fmt"
	"log"
	"sync"

	"github.com/jinzhu/gorm"
	"github.com/xorcare/pointer"
)

type IMemberRepo interface {
	Insert(context.Context, *models_rep.Member) error
	Find(context.Context, *string, *string) (*models_rep.Member, error)
	FindAll(context.Context) (*[]models_rep.Member, error)
	Updates(context.Context, *models_rep.UpdateMember) error
	Disable(context.Context, *models_rep.UpdateMember) error
}
type MemberRepo struct {
	mutex sync.Mutex
	db    *gorm.DB
}

func NewMemberRepo(db *gorm.DB) IMemberRepo {
	return &MemberRepo{
		db: db,
	}
}

func (rep *MemberRepo) Insert(ctx context.Context, param *models_rep.Member) error {
	rep.mutex.Lock()
	defer rep.mutex.Unlock()
	db := rep.db.DB()
	query := `INSERT INTO member(account, password, permission, name, email, phone, address) VALUES ($1,$2,$3,$4,$5,$6,$7);`
	_, err := db.ExecContext(ctx, query, param.Account, param.Password, param.Permission, param.Name, param.Email, param.Phone, param.Address)
	if err != nil {
		return fmt.Errorf("%s", err.Error())
	}
	return nil
}
func (rep *MemberRepo) Find(ctx context.Context, account *string, phone *string) (*models_rep.Member, error) {
	rep.mutex.Lock()
	defer rep.mutex.Unlock()
	db := rep.db.DB()
	result := &models_rep.Member{}
	if account == nil {
		account = pointer.String("")
	}
	if phone == nil {
		phone = pointer.String("")
	}
	query := `SELECT account,password,permission,name,email,phone,address,create_at FROM member` +
		` WHERE is_alive = true AND account =$1 OR phone = $2 ORDER BY id`
	row := db.QueryRowContext(ctx, query, *account, *phone)
	if err := row.Scan(&result.Account, &result.Password, &result.Permission, &result.Name, &result.Email, &result.Phone, &result.Address, &result.CreateAt); err != nil {
		return nil, err
	}
	return result, nil
}
func (rep *MemberRepo) FindAll(ctx context.Context) (*[]models_rep.Member, error) {
	rep.mutex.Lock()
	defer rep.mutex.Unlock()
	db := rep.db.DB()
	result := []models_rep.Member{}
	query := `SELECT account,password,permission,name,email,phone,address,create_at
	FROM member
	WHERE is_alive = true`
	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			log.Println("No rows were returned!")
			return nil, err
		default:
			log.Printf("Unable to scan the row. %v", err)
		}
	}
	for rows.Next() {
		model := new(models_rep.Member)
		if err := rows.Scan(&model.Account, &model.Password, &model.Permission, &model.Name, &model.Email, &model.Phone, &model.Address, &model.CreateAt); err != nil {
			return nil, err
		}
		result = append(result, *model)
	}

	return &result, err
}
func (rep *MemberRepo) Updates(ctx context.Context, param *models_rep.UpdateMember) error {
	rep.mutex.Lock()
	defer rep.mutex.Unlock()
	db := rep.db.DB()
	query := `UPDATE member SET name =$2, email =$3,phone=$4, address=$5 WHERE account = $1`
	_, err := db.ExecContext(ctx, query, param.Account, param.Name, param.Email, param.Phone, param.Address)
	if err != nil {
		return fmt.Errorf("%s", err.Error())
	}
	return nil
}
func (rep *MemberRepo) Disable(ctx context.Context, param *models_rep.UpdateMember) error {
	rep.mutex.Lock()
	defer rep.mutex.Unlock()
	db := rep.db.DB()
	query := `UPDATE member SET is_alive = false WHERE account = $1`
	_, err := db.ExecContext(ctx, query, param.Account)
	return err
}
