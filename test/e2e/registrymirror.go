//go:build e2e
// +build e2e

package e2e

import (
	"github.com/aws/eks-anywhere/test/framework"
)

func runRegistryMirrorConfigFlow(test *framework.ClusterE2ETest) {
	test.GenerateClusterConfig()
	test.DownloadArtifacts()
	test.ExtractDownloadedArtifacts()
	test.DownloadImages()
	test.ImportImages()
	test.CreateCluster()
	test.DeleteCluster()
}

func runTinkerbellRegistryMirrorFlow(test *framework.ClusterE2ETest) {
	test.GenerateClusterConfig()
	test.DownloadArtifacts()
	test.ExtractDownloadedArtifacts()
	test.DownloadImages()
	test.ImportImages()
	test.GenerateHardwareConfig()
	test.PowerOffHardware()
	test.CreateCluster(framework.WithForce())
	test.StopIfFailed()
	test.DeleteCluster()
	test.ValidateHardwareDecommissioned()
	test.CleanupDownloadedArtifactsAndImages()
}
