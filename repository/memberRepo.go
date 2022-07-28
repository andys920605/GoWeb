package repository

import (
	models_rep "GoWeb/models/repository"
	"context"
	"database/sql"
	"fmt"
	"log"
	"sync"

	"github.com/jinzhu/gorm"
)

type IMemberRepo interface {
	Insert(context.Context, *models_rep.Member) error
	Find() bool
	FindAll(context.Context) (*[]models_rep.Member, error)
	Updates() bool
	Disable() bool
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
func (rep *MemberRepo) Find() bool {
	return false
}
func (rep *MemberRepo) FindAll(ctx context.Context) (*[]models_rep.Member, error) {
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
func (rep *MemberRepo) Updates() bool {
	return false
}
func (rep *MemberRepo) Disable() bool {
	return false
}
