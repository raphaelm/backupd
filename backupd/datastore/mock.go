package datastore

import (
	"errors"

	"github.com/raphaelm/backupd/backupd/model"
)

type mockStore struct {
	idCounter int64
	remotes   map[int64]model.Remote
	jobs      map[int64]model.Job
	backups   map[int64]model.Backup
}

func MockStore() *mockStore {
	s := mockStore{}
	s.remotes = make(map[int64]model.Remote)
	s.jobs = make(map[int64]model.Job)
	s.backups = make(map[int64]model.Backup)
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
	jobs, err := s.JobsForRemote(remote)
	if err != nil {
		return false, err
	}
	for _, j := range jobs {
		s.DeleteJob(&j)
	}
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

func (s *mockStore) SaveJob(job *model.Job) (created bool, err error) {
	if _, ok := s.remotes[job.RemoteID]; !ok {
		return false, errors.New("Remote not found")
	}
	if job.ID == 0 {
		s.idCounter++
		job.ID = s.idCounter
		created = true
	}
	s.jobs[job.ID] = *job
	return created, nil
}

func (s *mockStore) DeleteJob(job *model.Job) (deleted bool, err error) {
	backups, err := s.BackupsForJob(job)
	if err != nil {
		return false, err
	}
	for _, b := range backups {
		s.DeleteBackup(&b)
	}
	if _, ok := s.jobs[job.ID]; ok {
		delete(s.jobs, job.ID)
		deleted = true
	}
	job.ID = 0
	return deleted, nil
}

func (s *mockStore) Job(id int64) (job model.Job, err error) {
	if job, ok := s.jobs[id]; ok {
		return job, nil
	}
	return model.Job{}, errors.New("Object not found")
}

func (s *mockStore) Jobs() (jobs []model.Job, err error) {
	v := make([]model.Job, 0, len(jobs))
	for _, j := range s.jobs {
		v = append(v, j)
	}
	return v, nil
}

func (s *mockStore) JobsForRemote(remote *model.Remote) (jobs []model.Job, err error) {
	v := make([]model.Job, 0, len(jobs))
	for _, j := range s.jobs {
		if j.RemoteID == remote.ID {
			v = append(v, j)
		}
	}
	return v, nil
}

func (s *mockStore) SaveBackup(backup *model.Backup) (created bool, err error) {
	if _, ok := s.jobs[backup.JobID]; !ok {
		return false, errors.New("Job not found")
	}
	if backup.ID == 0 {
		s.idCounter++
		backup.ID = s.idCounter
		created = true
	}
	s.backups[backup.ID] = *backup
	return created, nil
}

func (s *mockStore) DeleteBackup(backup *model.Backup) (deleted bool, err error) {
	if _, ok := s.backups[backup.ID]; ok {
		delete(s.backups, backup.ID)
		deleted = true
	}
	backup.ID = 0
	return deleted, nil
}

func (s *mockStore) Backup(id int64) (backup model.Backup, err error) {
	if backup, ok := s.backups[id]; ok {
		return backup, nil
	}
	return model.Backup{}, errors.New("Object not found")
}

func (s *mockStore) Backups() (backups []model.Backup, err error) {
	v := make([]model.Backup, 0, len(backups))
	for _, j := range s.backups {
		v = append(v, j)
	}
	return v, nil
}

func (s *mockStore) BackupsForJob(job *model.Job) (backups []model.Backup, err error) {
	v := make([]model.Backup, 0, len(backups))
	for _, b := range s.backups {
		if b.JobID == job.ID {
			v = append(v, b)
		}
	}
	return v, nil
}
