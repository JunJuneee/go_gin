package service

type LoginService interface {
	Login(username, password string) bool
}

type loginService struct {
	authorizedUserName string
	authorizedPassword string
}

func NewLoginService() LoginService {
	return &loginService{
		"pragmatic",
		"reviews",
	}
}

func (s *loginService) Login(username, password string) bool {
	return s.authorizedUserName == username && s.authorizedPassword == password
}
