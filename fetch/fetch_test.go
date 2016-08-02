// Copyright © 2016 Asteris, LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package fetch_test

import (
	"testing"

	"golang.org/x/net/context"

	"github.com/asteris-llc/converge/fetch"
	"github.com/stretchr/testify/assert"
)

func TestAnyBadURL(t *testing.T) {
	t.Parallel()

	_, err := fetch.Any(context.Background(), ":://asdf")
	assert.Error(t, err)
}

func TestAnyUnimplementedProtocol(t *testing.T) {
	t.Parallel()

	_, err := fetch.Any(context.Background(), "nope://")
	if assert.Error(t, err) {
		assert.EqualError(t, err, `protocol "nope" is not implemented`)
	}
}
