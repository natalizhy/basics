package teststore

import (
	"github.com/natalizhy/basics/http-rest-api/internal/app/model"
	"github.com/natalizhy/basics/http-rest-api/internal/app/store"
)

type Store struct {
	userRepository *UserRepository
}

func New() *Store {
}

func (s *Store) User() store.UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}

	s.userRepository = &UserRepository{
		store: s,
		users: make(map[int]*model.User),
	}

	return s.userRepository
}

//store.User().Create{}   вызывать юзера
