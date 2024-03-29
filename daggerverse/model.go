package daggerverse

import (
	"time"
)

type RepositoryInfo struct {
	BrowseURL string    `json:"browse_url"`
	CreatedAt time.Time `json:"created_at"`
	GitCommit string    `json:"git_commit"`
	GitURL    string    `json:"git_url"`
	IndexedAt time.Time `json:"indexed_at"`
	Path      string    `json:"path"`
	Release   string    `json:"release"`
	Subpath   string    `json:"subpath"`
	Version   string    `json:"version"`
}

type CustomizedRepositoryInfo struct {
	RepositoryInfo
	Author         string        `json:"author"`
	AuthorRepoURL  string        `json:"author_url"`
	Age            time.Duration `json:"age"`
	ProjectDir     string        `json:"project_dir"`
	ModInstallPath string        `json:"mod_install_path"`
}

type (
	CustomizedRepositoryInfos     []*CustomizedRepositoryInfo
	CustomizedRepositoryInfoSlice []CustomizedRepositoryInfo
)

func (cr CustomizedRepositoryInfos) AddRepo(repo *CustomizedRepositoryInfo) CustomizedRepositoryInfos {
	return append(cr, repo)
}
