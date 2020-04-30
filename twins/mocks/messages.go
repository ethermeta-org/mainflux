// Copyright (c) Mainflux
// SPDX-License-Identifier: Apache-2.0

package mocks

import (
	"github.com/mainflux/mainflux/errors"
	"github.com/mainflux/mainflux/messaging"
)

var _ messaging.Publisher = (*mockBroker)(nil)

type mockBroker struct {
	subscriptions map[string]string
}

// New returns mock message publisher.
func New(sub map[string]string) messaging.Publisher {
	return &mockBroker{
		subscriptions: sub,
	}
}

func (mb mockBroker) Publish(topic string, msg messaging.Message) error {
	if len(msg.Payload) == 0 {
		return errors.New("failed to publish")
	}
	return nil
}