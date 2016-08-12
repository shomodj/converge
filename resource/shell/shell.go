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

package shell

import (
	"bytes"
	"fmt"
	"os/exec"
	"syscall"

	"github.com/asteris-llc/converge/resource"
)

// Shell task
type Shell struct {
	Interpreter string
	CheckStmt   string
	ApplyStmt   string
}

// Check system using CheckStmt
func (s *Shell) Check() (resource.TaskStatus, error) {
	out, code, err := s.exec(s.CheckStmt)
	return resource.NewStatus(out, code != 0, err)
}

// Apply ApplyStmt stanza to system
func (s *Shell) Apply() (err error) {
	out, code, err := s.exec(s.ApplyStmt)
	if code != 0 {
		return fmt.Errorf("exit code %d, output: %q", code, out)
	}

	return err
}
func (s *Shell) exec(script string) (out string, code uint32, err error) {
	command := exec.Command(s.Interpreter)
	stdin, err := command.StdinPipe()
	if err != nil {
		return "", 0, err
	}

	// TODO: does this create a race condition?
	var sink bytes.Buffer
	command.Stdout = &sink
	command.Stderr = &sink

	if err = command.Start(); err != nil {
		return "", 0, err
	}

	if _, err = stdin.Write([]byte(script)); err != nil {
		return "", 0, err
	}

	if err = stdin.Close(); err != nil {
		return "", 0, err
	}

	err = command.Wait()
	if _, ok := err.(*exec.ExitError); !ok && err != nil {
		return "", 0, err
	}

	switch result := command.ProcessState.Sys().(type) {
	case syscall.WaitStatus:
		code = uint32(result)
	default:
		panic(fmt.Sprintf("unknown type %+v", result))
	}

	return sink.String(), code, nil
}
