package repofs

import (
	"time"

	"github.com/kopia/kopia/fs"
	"github.com/kopia/kopia/repo"
	"github.com/kopia/kopia/snapshot"
)

type sourceDirectories struct {
	parent          fs.Directory
	repo            *repo.Repository
	snapshotManager *snapshot.Manager
	userHost        string
}

func (s *sourceDirectories) Parent() fs.Directory {
	return s.parent
}

func (s *sourceDirectories) Metadata() *fs.EntryMetadata {
	return &fs.EntryMetadata{
		Name:        s.userHost,
		Permissions: 0555,
		Type:        fs.EntryTypeDirectory,
		ModTime:     time.Now(),
	}
}

func (s *sourceDirectories) Readdir() (fs.Entries, error) {
	sources, err := s.snapshotManager.ListSources()
	if err != nil {
		return nil, err
	}

	var result fs.Entries

	for _, src := range sources {
		if src.UserName+"@"+src.Host != s.userHost {
			continue
		}

		result = append(result, &sourceSnapshots{s, s.repo, s.snapshotManager, src})
	}

	return result, nil
}
