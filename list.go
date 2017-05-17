// Copyright 2016 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package gogs

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type List struct {
	ID          int64  `json:"ID"`
	Title       string `json:"Title"`
	Index 		int    `json:"Index"`
}

func (c *Client) GetList(owner, repo string, id int64) (*List, error) {
	list := new(List)
	return list, c.getParsedResponse("GET", fmt.Sprintf("/repos/%s/%s/list/%d", owner, repo, id), nil, nil, list)
}

type CreateListOption struct {
	Title   string `json:"Title"`
	Index	int    `json:"Index"`
}

func (c *Client) CreateList(owner, repo string, opt CreateListOption) (*List, error) {
	body, err := json.Marshal(&opt)
	if err != nil {
		return nil, err
	}
	list := new(List)
	return list, c.getParsedResponse("POST", fmt.Sprintf("/repos/%s/%s/list", owner, repo), jsonHeader, bytes.NewReader(body), list)
}

type EditListOption struct {
	Title   string `json:"Title"`
	Index	int    `json:"Index"`
}

func (c *Client) EditList(owner, repo string, id int64, opt EditListOption) (*List, error) {
	body, err := json.Marshal(&opt)
	if err != nil {
		return nil, err
	}
	list := new(List)
	return list, c.getParsedResponse("PATCH", fmt.Sprintf("/repos/%s/%s/list/%d", owner, repo, id), jsonHeader, bytes.NewReader(body), list)
}

func (c *Client) DeleteList(owner, id int64, repo string) error {
	_, err := c.getResponse("DELETE", fmt.Sprintf("/repos/%s/%s/list/%d", owner, repo, id), nil, nil)
	return err
}
