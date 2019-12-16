// Copyright 2018 Author: Ruslan Bikchentaev. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package telegrambot

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewTelegramBot(t *testing.T) {
	b, err := NewTelegramBot("cfg.yml")
	if !assert.Nil(t, err) {
		t.Fatal(err)
	}
	assert.NotEmpty(t, b.Token)

	err = b.GetUpdates()
	if !assert.Nil(t, err) {
		t.Fatal(err)
	}

	t.Log(b.props["result"])

	err = b.GetChat(b.ChatID)
	if !assert.Nil(t, err) {
		t.Fatal(err)
	}
	t.Log(b.props["result"])

	err = b.GetChatMemberCount(b.ChatID)
	if !assert.Nil(t, err) {
		t.Fatal(err)
	}

	c := b.props["result"].(float64)
	// for i := ; i < c; i++ {
	b.GetChatMember(b.ChatID, "91653754")
	b.SendMessage(fmt.Sprintf("user - %v", b.props["result"]), true)
	// }
	b.SendMessage(fmt.Sprintf("try to sum users of group - %f", c), true)
	t.Log(b.Response.String())
}
