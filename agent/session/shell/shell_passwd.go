// Copyright 2018 Amazon.com, Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may not
// use this file except in compliance with the License. A copy of the
// License is located at
//
// http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
// either express or implied. See the License for the specific language governing
// permissions and limitations under the License.

//go:build darwin || freebsd || linux || netbsd || openbsd
// +build darwin freebsd linux netbsd openbsd

// Package shell implements session shell plugin.

//fork by liwadman@. Parses /etc/passwd, and sets the users shell as the 
package shell

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"syscall"
	"time"
	"bufio"

	"github.com/aws/amazon-ssm-agent/agent/appconfig"
	agentContracts "github.com/aws/amazon-ssm-agent/agent/contracts"
	"github.com/aws/amazon-ssm-agent/agent/log"
	mgsConfig "github.com/aws/amazon-ssm-agent/agent/session/config"
	mgsContracts "github.com/aws/amazon-ssm-agent/agent/session/contracts"
	"github.com/aws/amazon-ssm-agent/agent/session/shell/constants"
	"github.com/aws/amazon-ssm-agent/agent/session/shell/execcmd"
	"github.com/aws/amazon-ssm-agent/agent/session/utility"
	"github.com/creack/pty"
	"github.com/google/shlex"
)


func readFile() ([]string, error) {
	//slurp up /etc/passwd
    file, err := os.Open("/etc/passwd")
    if err != nil {
       fmt.Println(err)
    }
    defer file.Close()
	fmt.Println(file)
	
	var lines []string
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        lines = append(lines, scanner.Text())
    }
	return lines //send lines back to caller
   // fmt.Println(lines[0])
	//fmt.Println(strings.Split(lines[0],":")) //split by colon
	//line := strings.Split(lines[0],":")[6] //the shell is space 6
	//fmt.Println(line)
	
}

func getShell(userName) ([]string, error) {
	//get the users shell from their username
    passwd := readFile("/etc/passwd")

	for _, line := range passwd {
		line = strings.Split(line, ":")
		if line[0] == userName {
			return line[6] //return the shell from /etc/passwd. We don't need to sanity check this because if it's wrong they have bigger problems.
		}

		

	}

}