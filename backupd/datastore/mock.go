package datastore

import (
	"errors"

	"github.com/raphaelm/backupd/backupd/model"
)

type mockStore struct {
	idCounter int64
	remotes   map[int64]model.Remote
}

func MockStore() *mockStore {
	s := mockStore{}
	s.remotes = make(map[int64]model.Remote)
	return &s
}

func (s *mockStore) SaveRemote(remote *model.Remote) (created bool, err error) {
	if remote.ID == 0 {
		s.idCounter++
		remote.ID = s.idCounter
		created = true
	}
	s.remotes[remote.ID] = *remote
	return created, nil
}

func (s *mockStore) DeleteRemote(remote *model.Remote) (deleted bool, err error) {
	if _, ok := s.remotes[remote.ID]; ok {
		delete(s.remotes, remote.ID)
		deleted = true
	}
	remote.ID = 0
	return deleted, nil
}

func (s *mockStore) Remote(id int64) (remote model.Remote, err error) {
	if remote, ok := s.remotes[id]; ok {
		return remote, nil
	}
	return model.Remote{}, errors.New("Object not found")
}

func (s *mockStore) Remotes() (remotes []model.Remote, err error) {
	v := make([]model.Remote, 0, len(remotes))
	for _, r := range s.remotes {
		v = append(v, r)
	}
	return v, nil
}
