// Copyright 2014 The go-gitlab AUTHORS. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gitlab

import (
	"bytes"
	"fmt"
	"net/url"
)

type ProjectFile struct {
	ID   *string `json:"id,omitempty"`
	Mode *string `json:"mode,omitempty"`
	Name *string `json:"name,omitempty"`
	Type *string `json:"type,omitempty"`
}

func (p ProjectFile) String() string {
	return Stringify(p)
}

// Get a list of repository files and directories in a project.
//
// TODO: GitLab API docs: http://doc.gitlab.com/ce/api/repositories.html#list-repository-tree
func (s *ProjectsService) Tree(projectID int, path string, ref_name string) (*[]ProjectFile, *Response, error) {
	data := new([]ProjectFile)

	params := url.Values{}
	params.Add("path", path)
	if ref_name != "" {
		params.Add("ref_name", ref_name)
	}
	u := fmt.Sprintf("projects/%v/repository/tree?%v", projectID, params.Encode())

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return data, nil, err
	}

	resp, err := s.client.Do(req, data)
	if err != nil {
		return data, resp, err
	}

	return data, resp, err
}

// Get the raw file contents for a blob by blob sha.
//
// TODO: GitLab API docs: http://doc.gitlab.com/ce/api/repositories.html#raw-blob-content

// Get a an archive of the repository
//
// TODO: GitLab API docs: http://doc.gitlab.com/ce/api/repositories.html#get-file-archive

// Get the raw file contents for a file by commit sha and path. Returns
// a *bytes.Buffer with the file data.
//
// GitLab API docs: http://doc.gitlab.com/ce/api/repositories.html#raw-file-content
func (s *ProjectsService) GetFileContents(projectID int, sha, filepath string) (*bytes.Buffer, *Response, error) {
	data := bytes.NewBuffer(nil)

	params := url.Values{}
	params.Add("filepath", filepath)
	u := fmt.Sprintf("projects/%v/repository/blobs/%v?%v", projectID, sha, params.Encode())

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return data, nil, err
	}

	resp, err := s.client.Do(req, data)
	if err != nil {
		return data, resp, err
	}

	return data, resp, err
}
