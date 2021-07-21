package domain

import "fmt"

type UserService struct {
	userRepository UserRepository
}

func NewUserService(userRepository UserRepository) *UserService {
	return &UserService{userRepository: userRepository}
}

func (s *UserService) CreateUser(cmd *CreateUserCommand) (*User, error) {
	id, err := s.userRepository.NextId()
	if err != nil {
		return nil, err
	}

	user := NewUser(id, cmd.Username, cmd.FirstName, cmd.LastName, cmd.Email, cmd.Phone)
	if err := s.userRepository.Save(user); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserService) GetUser(cmd *GetUserCommand) (*User, error) {
	user, err := s.userRepository.FindById(cmd.Id)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, NewEntityNotExists(fmt.Sprintf("user with id '%d' deos not exists", cmd.Id))
	}

	return user, nil
}

func (s *UserService) UpdateUser(cmd *UpdateUserCommand) (*User, error) {
	user, err := s.userRepository.FindById(cmd.Id)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, NewEntityNotExists(fmt.Sprintf("user with id '%d' deos not exists", cmd.Id))
	}

	if cmd.Email != "" {
		user.ChangeEmail(cmd.Email)
	}
	if cmd.Phone != "" {
		user.ChangePhone(cmd.Phone)
	}
	if cmd.LastName != "" {
		user.ChangeLastName(cmd.LastName)
	}
	if cmd.FirstName != "" {
		user.ChangeFirstName(cmd.FirstName)
	}

	if err := s.userRepository.Save(user); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserService) DeleteUser(cmd *DeleteUserCommand) error {
	user, err := s.userRepository.FindById(cmd.Id)
	if err != nil {
		return err
	}

	if user == nil {
		return NewEntityNotExists(fmt.Sprintf("user with id '%d' deos not exists", cmd.Id))
	}

	return s.userRepository.Delete(user)
}
