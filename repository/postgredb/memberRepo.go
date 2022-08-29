package postgredb

import (
	models_rep "GoWeb/models/repository"
	rep_interface "GoWeb/repository/interface"
	"context"
	"database/sql"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/xorcare/pointer"
)

type MemberRepo struct {
	mutex sync.Mutex
	db    *gorm.DB
}

func NewMemberRep(db *gorm.DB) rep_interface.IMemberRep {
	return &MemberRepo{
		db: db,
	}
}

// Close attaches the provider and close the connection
func (rep *MemberRepo) Close() {
	rep.db.Close()
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
	query := `SELECT account,password,permission,name,email,phone,address,create_at,update_at FROM member` +
		` WHERE is_alive = true AND account =$1 OR phone = $2 ORDER BY id`
	row := db.QueryRowContext(ctx, query, *account, *phone)
	if err := row.Scan(&result.Account, &result.Password, &result.Permission, &result.Name, &result.Email, &result.Phone, &result.Address, &result.CreateAt, &result.UpdateAt); err != nil {
		return nil, err
	}
	return result, nil
}
func (rep *MemberRepo) FindAll(ctx context.Context) (*[]models_rep.Member, error) {
	rep.mutex.Lock()
	defer rep.mutex.Unlock()
	db := rep.db.DB()
	result := []models_rep.Member{}
	query := `SELECT account,password,permission,name,email,phone,address,create_at,update_at
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
		if err := rows.Scan(&model.Account, &model.Password, &model.Permission, &model.Name, &model.Email, &model.Phone, &model.Address, &model.CreateAt, &model.UpdateAt); err != nil {
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
	query := `UPDATE member SET name =$2, email =$3,phone=$4, address=$5, update_at=$6 WHERE account = $1`
	_, err := db.ExecContext(ctx, query, param.Account, param.Name, param.Email, param.Phone, param.Address, time.Now().Format(time.RFC3339))
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
