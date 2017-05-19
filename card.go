// Copyright 2016 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package gogs

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type Card struct {
	ID        int64     `json:"ID"`
	List      int64     `json:"List"`
	Index     int64     `json:"Index"`
	Body      string    `json:"Body"`
	Assignee  *User     `json:"Assignee"`
	State     int		`json:"State"`
	Priority  int 		`json:"Priority"`
	Duration  int64  	`json:"Duration"`
	Time 	  int64 	`json:"Time"`
	Activated int64		`json:"Activated"`
	CreateD   int64 	`json:"Created"`
}

func (c *Client) GetCard(owner, repo string, id int64) (*Card, error) {
	card := new(Card)
	return card, c.getParsedResponse("GET", fmt.Sprintf("/repos/%s/%s/card/%d", owner, repo, id), nil, nil, card)
}

type CreateCardOption struct {
	List 	  int64   `json:"List"`
	Body      string  `json:"Body" binding:"Required"`
	Index  	  int64   `json:"Index"`
}

func (c *Client) CreateCard(owner, repo string, opt CreateCardOption) (*Card, error) {
	body, err := json.Marshal(&opt)
	if err != nil {
		return nil, err
	}
	card := new(Card)
	return card, c.getParsedResponse("POST", fmt.Sprintf("/repos/%s/%s/card", owner, repo), jsonHeader, bytes.NewReader(body), card)
}

type EditCardOption struct {
	Body  	  string 	`json:"Body"`
	Assignee  string 	`json:"Assignee"`
	Priority  int 		`json:"Priority"`
}

func (c *Client) EditCard(owner, repo string, list int64, index int64, opt EditCardOption) (*Card, error) {
	body, err := json.Marshal(&opt)
	if err != nil {
		return nil, err
	}
	card := new(Card)
	return card, c.getParsedResponse("PATCH", fmt.Sprintf("/repos/%s/%s/card/%d/%d", owner, repo, list, index), jsonHeader, bytes.NewReader(body), card)
}

type EditCardDurationOption struct {
	Duration  int64 	`json:"Duration"`
}

type EditCardIndexOption struct {
	Index  	  int64 	`json:"Index"`
}

type TransferCardOption struct {
	List  	  int64 	`json:"List"`
}

func (c *Client) DeleteCard(owner, repo string, list int64, index int64) error {
	_, err := c.getResponse("DELETE", fmt.Sprintf("/repos/%s/%s/card/%d/%d", owner, repo, list, index), nil, nil)
	return err
}
