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
	"bufio"
	"fmt"
	"os"
	"strings"
)

func readFile() []string {
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

func getShell(userName string) string {
	//get the users shell from their username
	passwd := readFile()

	for _, line := range passwd {
		slice := strings.Split(line, ":")
		if slice[0] == userName {
			return slice[6] //return the shell from /etc/passwd. We don't need to sanity check this because if it's wrong they have bigger problems.
		}

	}
	//default to sh, widest compatability
	return "/bin/sh"
}
