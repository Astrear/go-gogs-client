// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package gogs

// User represents a API user.

type CourseInformation struct {
	Subject interface{} `json:"Subjects"`
	Group interface{} `json:"Groups"`
	Semester interface{} `json:"Semesters"`
	Course interface{} `json:"Course"`
}
