// user/service.go
package user

type UserService struct {
	repo UserRepository
}

func NewUserService(repo UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) IsEmailRegistered(email string) bool {
	user, err := s.repo.FindByEmail(email)
	return err == nil && user != nil
}
