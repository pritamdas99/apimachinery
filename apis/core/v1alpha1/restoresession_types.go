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

package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	kmapi "kmodules.xyz/client-go/api/v1"
)

const (
	ResourceKindRestoreSession     = "RestoreSession"
	ResourceSingularRestoreSession = "restoresession"
	ResourcePluralRestoreSession   = "restoresessions"
)

// +k8s:openapi-gen=true
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=restoresessions,singular=restoresession,shortName=restore,categories={kubestash,appscode,all}
// +kubebuilder:printcolumn:name="Repository",type="string",JSONPath=".spec.dataSource.repository"
// +kubebuilder:printcolumn:name="Failure-Policy",type="string",JSONPath=".spec.failurePolicy"
// +kubebuilder:printcolumn:name="Phase",type="string",JSONPath=".status.phase"
// +kubebuilder:printcolumn:name="Duration",type="string",JSONPath=".status.duration"
// +kubebuilder:printcolumn:name="Age",type="date",JSONPath=".metadata.creationTimestamp"

// RestoreSession represents one restore run for the targeted application
type RestoreSession struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   RestoreSessionSpec   `json:"spec,omitempty"`
	Status RestoreSessionStatus `json:"status,omitempty"`
}

// RestoreSessionSpec specifies the necessary configurations for restoring data into a target
type RestoreSessionSpec struct {
	// Target indicates the target application where the data will be restored.
	// The target must be in the same namespace as the RestoreSession CR.
	// +optional
	Target *kmapi.TypedObjectReference `json:"target,omitempty"`

	// DataSource specifies the information about the data that will be restored
	DataSource *RestoreDataSource `json:"dataSource,omitempty"`

	// Addon specifies addon configuration that will be used to restore the target.
	Addon *AddonInfo `json:"addon,omitempty"`

	// Hooks specifies the restore hooks that should be executed before and/or after the restore.
	// +optional
	Hooks *RestoreHooks `json:"hooks,omitempty"`

	// FailurePolicy specifies what to do if the restore fail.
	// Valid values are:
	// - "Fail": Stash should mark the restore as failed if any component fail to complete its restore. This is the default behavior.
	// - "Retry": Stash will retry to restore the failed component according to the `retryConfig`.
	// +kubebuilder:default=Fail
	// +optional
	FailurePolicy FailurePolicy `json:"failurePolicy,omitempty"`

	// RetryConfig specifies the behavior of retry in case of a restore failure.
	// +optional
	RetryConfig *RetryConfig `json:"retryConfig,omitempty"`

	// Timeout specifies a duration in seconds that KubeStash should wait for the session execution to be completed.
	// If the session execution does not finish within this time period, KubeStash will consider this session as failure.
	// Then, it will re-try according to the RetryConfig.
	// +optional
	Timeout *int32 `json:"timeout,omitempty"`
}

// RestoreDataSource specifies the information about the data that will be restored
type RestoreDataSource struct {
	// Repository points to the Repository name from which the data will be restored
	Repository string `json:"repository,omitempty"`

	// Snapshot specifies the Snapshot name that will be restored.
	// If you want to use Point-In-Time recovery option, don't specify this field. Specify `pitr` field instead.
	// +optional
	Snapshot string `json:"snapshot,omitempty"`

	// PITR stands for Point-In-Time Recovery. You can provide a target time instead of specifying a particular Snapshot.
	// Stash will automatically find the latest Snapshot that satisfies the targeted time and restore it.
	// +optional
	PITR PITR `json:"pitr,omitempty"`

	// Components specifies the components that will be restored. If you keep this field empty, then all
	// the components that were backed up in the desired Snapshot will be restored.
	// +optional
	Components []string `json:"components,omitempty"`
}

// PITR specifies the target time and behavior of Point-In-Time Recovery
type PITR struct {
	// TargetTime specifies the desired date and time at which you want to roll back your application data
	// +kubebuilder:validation:Format=date-time
	TargetTime string `json:"targetTime,omitempty"`

	// Exclusive specifies whether to exclude the Snapshot that falls in the exact time specified
	// in the `targetTime` field. By default, Stash will select the Snapshot that fall in the exact time.
	// +optional
	Exclusive bool `json:"exclusive,omitempty"`
}

// RestoreHooks specifies the hooks that will be executed before and/or after restore
type RestoreHooks struct {
	// PreRestore specifies a list of hooks that will be executed before restore
	// +optional
	PreRestore []HookInfo `json:"preRestore,omitempty"`

	// PostRestore specifies a list of hooks that will be executed after restore
	// +optional
	PostRestore []HookInfo `json:"postRestore,omitempty"`
}

// RestoreSessionStatus defines the observed state of RestoreSession
type RestoreSessionStatus struct {
	// Phase represents the current state of the restore process
	// +optional
	Phase RestorePhase `json:"phase,omitempty"`

	// Duration specify total time taken to complete the restore process
	// +optional
	Duration string `json:"duration,omitempty"`

	// Deadline specifies a timestamp till this session is valid. If the session does not complete within this deadline,
	// it will be considered as failed.
	Deadline string `json:"deadline,omitempty"`

	// Components represents the individual component restore status
	// +optional
	Components []ComponentRestoreStatus `json:"components,omitempty"`

	// Hooks represents the hook execution status
	// +optional
	Hooks []HookExecutionStatus `json:"hooks,omitempty"`

	// Backup represents the backup pause status. This section only present if the restore process
	// pauses any active backup before restoring
	// +optional
	Backup *BackupPausedStatus `json:"backup,omitempty"`

	// Conditions specifies a list of conditions related to this restore session
	// +optional
	Conditions []kmapi.Condition `json:"conditions,omitempty"`
}

// RestorePhase represents the current state of the restore process
// +kubebuilder:validation:Enum=Pending;Running;Failed;Skipped
type RestorePhase string

const (
	RestorePending RestorePhase = "Pending"
	RestoreRunning RestorePhase = "Running"
	RestoreFailed  RestorePhase = "Failed"
	RestoreSkipped RestorePhase = "Skipped"
)

// ComponentRestoreStatus represents the restore status of individual components
type ComponentRestoreStatus struct {
	// Name indicate to the name of the component
	Name string `json:"name,omitempty"`

	// Phase represents the restore phase of the component
	// +optional
	Phase RestorePhase `json:"phase,omitempty"`

	// Duration specify total time taken to complete the restore process for this component
	// +optional
	Duration string `json:"duration,omitempty"`

	// Error specifies the reason in case of restore failure for the component
	// +optional
	Error string `json:"error,omitempty"`
}

// BackupPausedStatus holds the respective backup invoker information if the restore
// process pauses any active backup.
type BackupPausedStatus struct {
	// Paused specifies whether the respective backup has been paused or not
	Paused bool `json:"paused,omitempty"`

	// Invoker refers to the respective backup invoker
	Invoker *core.TypedLocalObjectReference `json:"invoker,omitempty"`
}

//+kubebuilder:object:root=true

// RestoreSessionList contains a list of RestoreSession
type RestoreSessionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RestoreSession `json:"items"`
}

func init() {
	SchemeBuilder.Register(&RestoreSession{}, &RestoreSessionList{})
}