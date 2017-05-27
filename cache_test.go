// Copyright 2017 The Xorm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package xorm

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCacheFind(t *testing.T) {
	assert.NoError(t, prepareEngine())

	type MailBox struct {
		Id       int64
		Username string
		Password string
	}

	cacher := NewLRUCacher2(NewMemoryStore(), time.Hour, 10000)
	testEngine.SetDefaultCacher(cacher)

	assert.NoError(t, testEngine.Sync2(new(MailBox)))

	var inserts = []*MailBox{
		{
			Username: "user1",
			Password: "pass1",
		},
		{
			Username: "user2",
			Password: "pass2",
		},
	}
	_, err := testEngine.Insert(inserts[0], inserts[1])
	assert.NoError(t, err)

	var boxes []MailBox
	assert.NoError(t, testEngine.Find(&boxes))
	assert.EqualValues(t, 2, len(boxes))
	for i, box := range boxes {
		assert.Equal(t, inserts[i].Id, box.Id)
		assert.Equal(t, inserts[i].Username, box.Username)
		assert.Equal(t, inserts[i].Password, box.Password)
	}

	boxes = make([]MailBox, 0, 2)
	assert.NoError(t, testEngine.Find(&boxes))
	assert.EqualValues(t, 2, len(boxes))
	for i, box := range boxes {
		assert.Equal(t, inserts[i].Id, box.Id)
		assert.Equal(t, inserts[i].Username, box.Username)
		assert.Equal(t, inserts[i].Password, box.Password)
	}

	testEngine.SetDefaultCacher(nil)
}

func TestCacheFind2(t *testing.T) {
	assert.NoError(t, prepareEngine())

	type MailBox2 struct {
		Id       uint64
		Username string
		Password string
	}

	cacher := NewLRUCacher2(NewMemoryStore(), time.Hour, 10000)
	testEngine.SetDefaultCacher(cacher)

	assert.NoError(t, testEngine.Sync2(new(MailBox2)))

	var inserts = []*MailBox2{
		{
			Username: "user1",
			Password: "pass1",
		},
		{
			Username: "user2",
			Password: "pass2",
		},
	}
	_, err := testEngine.Insert(inserts[0], inserts[1])
	assert.NoError(t, err)

	var boxes []MailBox2
	assert.NoError(t, testEngine.Find(&boxes))
	assert.EqualValues(t, 2, len(boxes))
	for i, box := range boxes {
		assert.Equal(t, inserts[i].Id, box.Id)
		assert.Equal(t, inserts[i].Username, box.Username)
		assert.Equal(t, inserts[i].Password, box.Password)
	}

	boxes = make([]MailBox2, 0, 2)
	assert.NoError(t, testEngine.Find(&boxes))
	assert.EqualValues(t, 2, len(boxes))
	for i, box := range boxes {
		assert.Equal(t, inserts[i].Id, box.Id)
		assert.Equal(t, inserts[i].Username, box.Username)
		assert.Equal(t, inserts[i].Password, box.Password)
	}

	testEngine.SetDefaultCacher(nil)
}
