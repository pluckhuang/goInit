package repository

import (
	"context"

	"gitee.com/geekbang/basic-go/webook/internal/domain"
	"gitee.com/geekbang/basic-go/webook/internal/repository/dao"
)

var (
	ErrUserDuplicateEmail = dao.ErrUserDuplicateEmail
	ErrUserNotFound       = dao.ErrUserNotFound
)

type UserRepository struct {
	dao *dao.UserDAO
}

func NewUserRepository(dao *dao.UserDAO) *UserRepository {
	return &UserRepository{
		dao: dao,
	}
}

func (r *UserRepository) FindByEmail(ctx context.Context, email string) (domain.User, error) {
	// SELECT * FROM `users` WHERE `email`=?
	u, err := r.dao.FindByEmail(ctx, email)
	if err != nil {
		return domain.User{}, err
	}
	return domain.User{
		Id:       u.Id,
		Email:    u.Email,
		Password: u.Password,
	}, nil
}

func (r *UserRepository) Create(ctx context.Context, u domain.User) error {
	return r.dao.Insert(ctx, dao.User{
		Email:    u.Email,
		Password: u.Password,
	})
}

func (r *UserRepository) FindById(ctx context.Context, userId int64) (domain.User, error) {
	u, err := r.dao.FindByID(ctx, userId)
	if err != nil {
		return domain.User{}, err
	}
	return domain.User{
		Id:          u.Id,
		Email:       u.Email,
		Description: u.Description,
		Birth:       u.Birth,
		NickName:    u.NickName,
	}, nil
}

func (r *UserRepository) Update(ctx context.Context, u domain.User) error {
	return r.dao.Update(ctx, dao.User{
		Id:          u.Id,
		Description: u.Description,
		Birth:       u.Birth,
		NickName:    u.NickName,
	})
}
