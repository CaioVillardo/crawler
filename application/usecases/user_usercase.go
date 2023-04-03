package usecases

type UserUseCase struct {
	UserRepository repositories.UserRepository
}

func (u *UserUseCase) Create(user *domain.User) (*domain.User, error) {
	user, err := u.UserRepository.Insert(user)

	if err != nil {
		log.Fatalf("Error to persist new User %v", err)
	}
	return user, err
}
