package business

import "gym/pkg/logger"

type AuthService struct {
	Test int
}

func (a *AuthService) TestFunc() {
	a.Test = 1
	logger.Logger().Println("TestFunc")
}
