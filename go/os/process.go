/*
   Copyright 2014 Outbrain Inc.

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/

package os

import (
	"bytes"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
	"syscall"

	"github.com/outbrain/golib/log"
	"github.com/outbrain/orchestrator/go/config"
)

// CommandRun executes some text as a command. This is assumed to be
// text that will be run by a shell so we need to write out the
// command to a temporary file and then ask the shell to execute
// it, after which the temporary file is removed.
func CommandRun(commandText string, arguments ...string) error {
	// show the actual command we have been asked to run
	log.Infof("CommandRun(%v,%+v)", commandText, arguments)

	cmd, shellScript, err := generateShellScript(commandText, arguments...)
	defer os.Remove(shellScript)
	if err != nil {
		return log.Errore(err)
	}

	cmdOutput := &bytes.Buffer{}
	cmdError := &bytes.Buffer{}
	cmd.Stdout = cmdOutput
	cmd.Stderr = cmdError
	var waitStatus syscall.WaitStatus

	log.Infof("CommandRun/running: %s", strings.Join(cmd.Args, " "))
	err = cmd.Run()
	logOutput("stdout", cmdOutput.Bytes())
	logOutput("stderr", cmdError.Bytes())
	if err != nil {
		// Did the command fail because of an unsuccessful exit code
		if exitError, ok := err.(*exec.ExitError); ok {
			waitStatus = exitError.Sys().(syscall.WaitStatus)
			log.Errorf("CommandRun: failed. exit status %d", waitStatus.ExitStatus())
		}

		return log.Errore(err)
	}

	// Command was successful
	waitStatus = cmd.ProcessState.Sys().(syscall.WaitStatus)
	log.Infof("CommandRun successful. exit status %d", waitStatus.ExitStatus())

	return nil
}

// generateShellScript generates a temporary shell script based on
// the given command to be executed, writes the command to a temporary
// file and returns the exec.Command which can be executed together
// with the script name that was created.
func generateShellScript(commandText string, arguments ...string) (*exec.Cmd, string, error) {
	shell := config.Config.ProcessesShellCommand

	commandBytes := []byte(commandText)
	tmpFile, err := ioutil.TempFile("", "orchestrator-process-cmd-")
	if err != nil {
		return nil, "", log.Errorf("generateShellScript() failed to create TempFile: %v", err.Error())
	}
	// write commandText to temporary file
	ioutil.WriteFile(tmpFile.Name(), commandBytes, 0640)
	shellArguments := append([]string{}, tmpFile.Name())
	shellArguments = append(shellArguments, arguments...)

	cmd := exec.Command(shell, shellArguments...)

	return cmd, tmpFile.Name(), nil
}

// log the output from the command. Provide an indication of what's being logged (e.g. stdout, stderr) as a name
func logOutput(name string, output []byte) {
	if len(output) == 0 {
		return
	}
	log.Infof("CommandRun/%s: %s\n", name, string(output))
}
