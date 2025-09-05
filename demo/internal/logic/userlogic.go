package logic

import (
	"context"
	"database/sql"
	"errors"

	"demo/internal/model"
	"demo/internal/svc"
	"demo/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// UserLogic 封装用户相关业务逻辑
// - CreateUser: 校验与加密密码，写入数据库
// - GetUser/ListUsers: 读取用户信息（不返回密码）
// - UpdateUser: 可选更新密码与邮箱（nil 表示不修改）
// - DeleteUser: 根据 ID 删除用户
type UserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserLogic {
	return &UserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// CreateUser 创建用户
// - 使用 bcrypt 对明文密码进行哈希
// - Email 可为空，使用 sql.NullString 表达可空
// - 返回对外脱敏后的用户视图（不包含密码）
func (l *UserLogic) CreateUser(req *types.CreateUserRequest) (*types.User, error) {
	l.Infof("CreateUser request: username=%s email=%v", req.Username, req.Email)
	if req.Username == "" || req.Password == "" {
		return nil, errors.New("username and password are required")
	}
	hashed, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	var email sql.NullString
	if req.Email != nil {
		email = sql.NullString{String: *req.Email, Valid: true}
	}
	user := &model.Users{Username: req.Username, PasswordHash: string(hashed), Email: email}
	if err := l.svcCtx.DB.WithContext(l.ctx).Create(user).Error; err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return nil, errors.New("username or email already exists")
		}
		return nil, err
	}
	return &types.User{ID: user.ID, Username: user.Username, Email: req.Email}, nil
}

// GetUser 根据 ID 查询单个用户
func (l *UserLogic) GetUser(req *types.IdPath) (*types.User, error) {
	var user model.Users
	if err := l.svcCtx.DB.WithContext(l.ctx).First(&user, req.ID).Error; err != nil {
		return nil, err
	}
	var emailPtr *string
	if user.Email.Valid {
		emailPtr = &user.Email.String
	}
	return &types.User{ID: user.ID, Username: user.Username, Email: emailPtr}, nil
}

// ListUsers 返回用户列表（脱敏视图）
func (l *UserLogic) ListUsers() ([]*types.User, error) {
	var users []model.Users
	if err := l.svcCtx.DB.WithContext(l.ctx).Order("id desc").Find(&users).Error; err != nil {
		return nil, err
	}
	resp := make([]*types.User, 0, len(users))
	for _, u := range users {
		uu := u
		var emailPtr *string
		if uu.Email.Valid {
			emailPtr = &uu.Email.String
		}
		resp = append(resp, &types.User{ID: uu.ID, Username: uu.Username, Email: emailPtr})
	}
	return resp, nil
}

// UpdateUser 支持部分字段更新
// - Password/Email 为空指针时不修改
// - Password 采用 bcrypt 重新哈希
func (l *UserLogic) UpdateUser(req *types.UpdateUserRequest) (*types.User, error) {
	var user model.Users
	if err := l.svcCtx.DB.WithContext(l.ctx).First(&user, req.ID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	if req.Password != nil {
		hashed, err := bcrypt.GenerateFromPassword([]byte(*req.Password), bcrypt.DefaultCost)
		if err != nil {
			return nil, err
		}
		user.PasswordHash = string(hashed)
	}
	if req.Email != nil {
		user.Email = sql.NullString{String: *req.Email, Valid: true}
	}
	if err := l.svcCtx.DB.WithContext(l.ctx).Save(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return nil, errors.New("email already exists")
		}
		return nil, err
	}
	var emailPtr *string
	if user.Email.Valid {
		emailPtr = &user.Email.String
	}
	return &types.User{ID: user.ID, Username: user.Username, Email: emailPtr}, nil
}

// DeleteUser 根据 ID 删除
func (l *UserLogic) DeleteUser(req *types.IdPath) error {
	return l.svcCtx.DB.WithContext(l.ctx).Delete(&model.Users{}, req.ID).Error
}
