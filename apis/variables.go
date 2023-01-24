/*
Copyright AppsCode Inc. and Contributors

Licensed under the AppsCode Free Trial License 1.0.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    https://github.com/appscode/licenses/raw/1.0.0/AppsCode-Free-Trial-1.0.0.md

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package apis

const (
	ImageRegistry = "imageRegistry"
	ImageRepo     = "imageRepo"
	ImageTag      = "imageTag"

	InvokerKind = "invokerKind"
	InvokerName = "invokerName"

	SnapshotName = "snapshotName"

	Namespace      = "namespace"
	BackupSession  = "backupSession"
	RestoreSession = "restoreSession"

	RepoName      = "repoName"
	RepoNamespace = "repoNamespace"

	TargetName      = "targetName"
	TargetKind      = "targetKind"
	TargetNamespace = "targetNamespace"
	TargetMountPath = "targetMountPath"
	TargetPaths     = "targetPaths"

	// default true
	// false when TmpDir.DisableCaching is true in backupConfig/restoreSession
	EnableCache    = "enableCache"
	InterimDataDir = "interimDataDir"

	KubeStashVolumePrefix = "/kubestash-volumes"
	DefaultMountPath      = KubeStashVolumePrefix + "/data"
	InterimDataDirPath    = KubeStashVolumePrefix + "/interim-data"
	TmpDirMountPath       = KubeStashVolumePrefix + "/tmp"

	LicenseApiService = "licenseApiService"
)