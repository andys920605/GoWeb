package repository

import (
	models_rep "GoWeb/models/repository"
	"context"
	"database/sql"
	"log"

	"github.com/jinzhu/gorm"
)

//go:generate mockgen -destination=../test/mock/imember_mock_repository.go -package=mock GoWeb/repository IMember

type IMemberRepo interface {
	Insert() bool
	Find() bool
	FindAll(context.Context) (*[]models_rep.Member, error)
	Updates() bool
	Disable() bool
}

type MemberRepo struct {
	db *gorm.DB
}

func NewMemberRepo(db *gorm.DB) IMemberRepo {
	return &MemberRepo{
		db: db,
	}
}

func (rep *MemberRepo) Insert() bool {
	return true
}
func (rep *MemberRepo) Find() bool {
	return false
}
func (rep *MemberRepo) FindAll(ctx context.Context) (*[]models_rep.Member, error) {
	db := rep.db.DB()
	result := []models_rep.Member{}
	query := `SELECT account,password,permission,name,email,phone,address,create_at
	FROM test4.member
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
