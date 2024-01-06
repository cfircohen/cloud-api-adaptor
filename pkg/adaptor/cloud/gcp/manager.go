// (C) Copyright Confidential Containers Contributors
// SPDX-License-Identifier: Apache-2.0

package gcp

import (
	"flag"

	"github.com/confidential-containers/cloud-api-adaptor/pkg/adaptor/cloud"
)

var gcpcfg Config

type Manager struct{}

func (_ *Manager) ParseCmd(flags *flag.FlagSet) {

	flags.StringVar(&gcpcfg.GcpCredentials, "gcp-credentials", "", "Google Application Credentials, defaults to `GCP_CREDENTIALS`")
	flags.StringVar(&gcpcfg.ProjectId, "gcp-project-id", "", "GCP Project ID")
	flags.StringVar(&gcpcfg.Zone, "gcp-zone", "", "Zone")
	flags.StringVar(&gcpcfg.ImageName, "gcp-image-name", "", "Pod VM image name")
	flags.StringVar(&gcpcfg.MachineType, "gcp-machine-type", "e2-medium", "Pod VM instance type")
	flags.StringVar(&gcpcfg.Network, "gcp-network", "", "Network ID to be used for the Pod VMs")
}

func (_ *Manager) LoadEnv() {
	cloud.DefaultToEnv(&gcpcfg.GcpCredentials, "GCP_CREDENTIALS", "")
}

func (_ *Manager) NewProvider() (cloud.Provider, error) {
	return NewProvider(&gcpcfg)
}
