/*
Copyright 2019 The Kubernetes Authors.

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

package main

import (
	"io"

	"github.com/spf13/cobra"
	"k8s.io/kops/cmd/kops/util"
	"k8s.io/kubectl/pkg/util/i18n"
	"k8s.io/kubectl/pkg/util/templates"
)

var (
	distrustLong = templates.LongDesc(i18n.T(`
	Distrust keypairs.
	`))

	distrustExample = templates.Examples(i18n.T(`
		# Distrust a keypair
        kops distrust keypair ca 6977545226837259959403993899

	`))

	distrustShort = i18n.T("Distrust keypairs.")
)

func NewCmdDistrust(f *util.Factory, out io.Writer) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "distrust",
		Short:   distrustShort,
		Long:    distrustLong,
		Example: distrustExample,
	}

	// create subcommands
	cmd.AddCommand(NewCmdDistrustKeypair(f, out))

	return cmd
}
