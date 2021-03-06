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

package mode

import (
	"os"

	"github.com/asteris-llc/converge/load/registry"
	"github.com/asteris-llc/converge/resource"
	"golang.org/x/net/context"
)

// Preparer for file Mode
//
// Mode monitors the mode of a file
type Preparer struct {
	// Destination specifies which file will be modified by this resource. The
	// file must exist on the system (for example, having been created with
	// `file.content`.)
	Destination string `hcl:"destination" required:"true" nonempty:"true"`

	// Mode is the mode of the file, specified in octal.
	Mode *uint32 `hcl:"mode" base:"8" required:"true"`
}

// Prepare this resource for use
func (p *Preparer) Prepare(ctx context.Context, render resource.Renderer) (resource.Task, error) {
	modeTask := &Mode{
		Destination: p.Destination,
		Mode:        os.FileMode(*p.Mode),
	}
	return modeTask, modeTask.Validate()
}

func init() {
	registry.Register("file.mode", (*Preparer)(nil), (*Mode)(nil))
}
