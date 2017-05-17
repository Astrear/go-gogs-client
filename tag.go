// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package gogs

// Tag represents a API user.
type Tag struct {
	ID        int64  `json:"id"`
	Etiqueta  string `json:"etiqueta"`
}