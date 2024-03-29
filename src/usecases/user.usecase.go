package usecases

import (
	"github.com/hm-mtmtmgs/mcdonalds-menu-gacha-backend/controllers/requests"
	"github.com/hm-mtmtmgs/mcdonalds-menu-gacha-backend/controllers/responses"
	"github.com/hm-mtmtmgs/mcdonalds-menu-gacha-backend/domains/entities"
	"github.com/hm-mtmtmgs/mcdonalds-menu-gacha-backend/domains/models"
	"github.com/hm-mtmtmgs/mcdonalds-menu-gacha-backend/domains/services"
	"github.com/hm-mtmtmgs/mcdonalds-menu-gacha-backend/domains/values"
	"github.com/hm-mtmtmgs/mcdonalds-menu-gacha-backend/repositories"
	"github.com/hm-mtmtmgs/mcdonalds-menu-gacha-backend/router/middlewares"
	"github.com/hm-mtmtmgs/mcdonalds-menu-gacha-backend/utils"
	"github.com/pkg/errors"
)

type IUserUsecase interface {
	SignUp(requests.SignUpRequest) error
	Login(requests.LoginRequest) (responses.TokenResponse, error)
	GetUser(uint) (responses.GetUserResponse, error)
}

type UserUsecase struct {
	baseRepository repositories.IBaseRepository
	userRepository repositories.IUserRepository
	userService    services.IUserService
}

func NewUserUsecase(
	baseRepository repositories.IBaseRepository,
	userRepository repositories.IUserRepository,
	userService services.IUserService,
) *UserUsecase {
	userUsecase := UserUsecase{
		baseRepository: baseRepository,
		userRepository: userRepository,
		userService:    userService,
	}
	utils.CheckDependencies(userUsecase)
	return &userUsecase
}

/*
サインアップ
*/
func (userUsecase *UserUsecase) SignUp(req requests.SignUpRequest) error {
	user, err := entities.NewUser(
		values.NewUserLastName(req.LastName),
		values.NewUserFirstName(req.FirstName),
		values.NewEmail(req.Email),
		values.NewUserPassword(req.Password),
	)
	if err != nil {
		return err
	}
	// メールアドレスの重複チェック
	if userUsecase.userService.IsUserEmailDuplication(user.Email) {
		return errors.Errorf("メールアドレスは既に使用されています")
	}
	err = userUsecase.userRepository.CreateUser(user)
	if err != nil {
		return errors.Errorf("Something went wrong")
	}
	return nil
}

/*
ログイン
*/
func (userUsecase *UserUsecase) Login(req requests.LoginRequest) (responses.TokenResponse, error) {
	var res responses.TokenResponse
	user, err := userUsecase.userRepository.GetUserByEmail(req.Email)
	if user == (models.User{}) {
		return res, errors.Errorf("Unauthorized")
	}
	if err != nil {
		return res, errors.Errorf("Something went wrong")
	}
	// パスワードチェック
	ok := utils.ComparePassword(user.Password, req.Password)
	if !ok {
		return res, errors.Errorf("Unauthorized")
	}
	// jwt生成
	token, err := middlewares.GenerateJwt(user.Id)
	if err != nil {
		return res, errors.Errorf("Something went wrong")
	}
	res = responses.NewTokenResponse(token)
	return res, nil
}

/*
ユーザ取得
*/
func (userUsecase *UserUsecase) GetUser(userId uint) (responses.GetUserResponse, error) {
	var res responses.GetUserResponse
	user, err := repositories.GetById[models.User](userUsecase.baseRepository.GetDB(), userId)
	if err != nil {
		return res, errors.Errorf("Unauthorized")
	}
	res = responses.NewGetUserResponse(user)
	return res, err
}
