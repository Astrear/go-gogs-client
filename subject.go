// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package gogs

// User represents a API user.
type Subject struct {
	ID      int64  `json:"id"`
	Name  	string `json:"name"`
}
