// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package gogs

// Tag represents a API user.
type Notification struct {
	ID        		int64  `json:"id"`
	UserID    		int64  `json:"user_id"`
	Watched    		bool   `json:"watched"`
	Description		string `json:"description"`
	Type 			int    `json:"type"`
	CreatedUnix 	int64  `json:"created_unix"`
	TimeSince		string `json:"time_since"`
}