// Copyright 2020 Amazon.com, Inc. or its affiliates. All Rights Reserved.
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
//
// +build darwin

// Package interactivecommands implements session shell plugin with interactive commands.
package interactivecommands

import (
	"fmt"
	"strings"

	"github.com/aws/amazon-ssm-agent/agent/session/contracts"
)

// validateProperties validates whether the commands are not empty.
func (p *InteractiveCommandsPlugin) validateProperties(shellProps contracts.ShellProperties) error {
	if strings.TrimSpace(shellProps.MacOS.Commands) == "" {
		return fmt.Errorf("Commands cannot be empty for session type %s", p.name())
	}
	return nil
}