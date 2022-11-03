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
	"testing"

	"github.com/stretchr/testify/assert"
	v1 "k8s.io/api/core/v1"
	v12 "k8s.io/apimachinery/pkg/apis/meta/v1"
	kmapi "kmodules.xyz/client-go/api/v1"
)

func TestRestoreSessionPhaseBasedOnComponentsPhase(t *testing.T) {
	tests := []struct {
		name           string
		restoreSession *RestoreSession
		expectedPhase  RestorePhase
	}{
		{
			name: "RestoreSession should be Pending if all components are Pending",
			restoreSession: sampleRestoreSession(func(r *RestoreSession) {
				r.Status.Components = []ComponentRestoreStatus{
					{
						Name:  "manifest",
						Phase: RestorePending,
					},
					{
						Name:  "configserver",
						Phase: RestorePending,
					},
					{
						Name:  "shard-0",
						Phase: RestorePending,
					},
					{
						Name:  "shard-1",
						Phase: RestorePending,
					},
				}
			}),

			expectedPhase: RestorePending,
		},
		{
			name: "RestoreSession should be Running if any component is Running",
			restoreSession: sampleRestoreSession(func(r *RestoreSession) {
				r.Status.Components = []ComponentRestoreStatus{
					{
						Name:  "manifest",
						Phase: RestoreRunning,
					},
					{
						Name:  "configserver",
						Phase: RestorePending,
					},
					{
						Name:  "shard-0",
						Phase: RestorePending,
					},
					{
						Name:  "shard-1",
						Phase: RestorePending,
					},
				}
			}),

			expectedPhase: RestoreRunning,
		},
		{
			name: "RestoreSession should be Running if any component is not completed",
			restoreSession: sampleRestoreSession(func(r *RestoreSession) {
				r.Status.Components = []ComponentRestoreStatus{
					{
						Name:  "manifest",
						Phase: RestoreSucceeded,
					},
					{
						Name:  "configserver",
						Phase: RestoreFailed,
					},
					{
						Name:  "shard-0",
						Phase: RestoreFailed,
					},
					{
						Name:  "shard-1",
						Phase: RestorePending,
					},
				}
			}),

			expectedPhase: RestoreRunning,
		},
		{
			name: "RestoreSession should be Failed if any component Failed",
			restoreSession: sampleRestoreSession(func(r *RestoreSession) {
				setPostRestoreHooksExecutionSucceededConditionToTrue(r)
				r.Status.Components = []ComponentRestoreStatus{
					{
						Name:  "manifest",
						Phase: RestoreFailed,
					},
					{
						Name:  "configserver",
						Phase: RestoreFailed,
					},
					{
						Name:  "shard-0",
						Phase: RestoreSucceeded,
					},
					{
						Name:  "shard-1",
						Phase: RestoreSucceeded,
					},
				}
			}),

			expectedPhase: RestoreFailed,
		},
		{
			name: "RestoreSession should be Failed if all components Failed",
			restoreSession: sampleRestoreSession(func(r *RestoreSession) {
				setPostRestoreHooksExecutionSucceededConditionToTrue(r)
				r.Status.Components = []ComponentRestoreStatus{
					{
						Name:  "manifest",
						Phase: RestoreFailed,
					},
					{
						Name:  "configserver",
						Phase: RestoreFailed,
					},
					{
						Name:  "shard-0",
						Phase: RestoreFailed,
					},
					{
						Name:  "shard-1",
						Phase: RestoreFailed,
					},
				}
			}),

			expectedPhase: RestoreFailed,
		},
		{
			name: "RestoreSession should be Succeeded if all components Succeeded",
			restoreSession: sampleRestoreSession(func(r *RestoreSession) {
				setPostRestoreHooksExecutionSucceededConditionToTrue(r)
				r.Status.Components = []ComponentRestoreStatus{
					{
						Name:  "manifest",
						Phase: RestoreSucceeded,
					},
					{
						Name:  "configserver",
						Phase: RestoreSucceeded,
					},
					{
						Name:  "shard-0",
						Phase: RestoreSucceeded,
					},
					{
						Name:  "shard-1",
						Phase: RestoreSucceeded,
					},
				}
			}),

			expectedPhase: RestoreSucceeded,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expectedPhase, test.restoreSession.CalculatePhase())
		})
	}
}

func TestRestoreSessionPhaseIsFailedIfPreRestoreHooksExecutionSucceededConditionIsFalse(t *testing.T) {
	rs := sampleRestoreSession(func(r *RestoreSession) {
		r.Status.Conditions = append(r.Status.Conditions, kmapi.Condition{
			Type:   TypePreRestoreHooksExecutionSucceeded,
			Status: v1.ConditionFalse,
			Reason: ReasonFailedToExecutePreRestoreHooks,
		})
	})

	assert.Equal(t, RestoreFailed, rs.CalculatePhase())
}

func TestRestoreSessionPhaseIsFailedIfPostRestoreHooksExecutionSucceededConditionIsFalse(t *testing.T) {
	rs := sampleRestoreSession(func(r *RestoreSession) {
		r.Status.Conditions = append(r.Status.Conditions, kmapi.Condition{
			Type:   TypePostRestoreHooksExecutionSucceeded,
			Status: v1.ConditionFalse,
			Reason: ReasonFailedToExecutePostRestoreHooks,
		})
	})

	assert.Equal(t, RestoreFailed, rs.CalculatePhase())
}

func TestRestoreSessionPhaseIsFailedIfRestoreExecutorEnsuredConditionIsFalse(t *testing.T) {
	rs := sampleRestoreSession(func(r *RestoreSession) {
		r.Status.Conditions = append(r.Status.Conditions, kmapi.Condition{
			Type:   TypeRestoreExecutorEnsured,
			Status: v1.ConditionFalse,
			Reason: ReasonFailedToEnsureRestoreExecutor,
		})
	})

	assert.Equal(t, RestoreFailed, rs.CalculatePhase())
}

func TestRestoreSessionPhaseIsFailedIfDeadlineExceededConditionIsTrue(t *testing.T) {
	rs := sampleRestoreSession(func(r *RestoreSession) {
		r.Status.Conditions = append(r.Status.Conditions, kmapi.Condition{
			Type:   TypeDeadlineExceeded,
			Status: v1.ConditionTrue,
			Reason: ReasonFailedToCompleteWithinDeadline,
		})
	})

	assert.Equal(t, RestoreFailed, rs.CalculatePhase())
}

func TestRestoreSessionPhaseIsRunningIfPostRestoreHooksNotExecuted(test *testing.T) {
	rs := sampleRestoreSession(func(r *RestoreSession) {
		r.Status.Components = []ComponentRestoreStatus{
			{
				Name:  "manifest",
				Phase: RestoreSucceeded,
			},
			{
				Name:  "configserver",
				Phase: RestoreSucceeded,
			},
			{
				Name:  "shard-0",
				Phase: RestoreSucceeded,
			},
			{
				Name:  "shard-0",
				Phase: RestoreSucceeded,
			},
		}
	})
	assert.Equal(test, RestoreRunning, rs.CalculatePhase())
}

func sampleRestoreSession(transformFuncs ...func(*RestoreSession)) *RestoreSession {
	rs := &RestoreSession{
		ObjectMeta: v12.ObjectMeta{
			Name:      "sample-mysql-restore",
			Namespace: "demo",
		},
		Spec: RestoreSessionSpec{
			Target: &kmapi.TypedObjectReference{
				APIGroup: "appcatalog.appscode.com",
				Kind:     "AppBinding",
				Name:     "sample-mysql",
			},
			DataSource: &RestoreDataSource{
				Repository: "sample-mysql-backup-gcs-storage",
				Snapshot:   "sample-mysql-backup-1561974001",
			},
			Addon: &AddonInfo{
				Name: "stash-mysql-90.31",
				Tasks: []TaskReference{
					{
						Name: "ManifestRestore",
					},
					{
						Name: "LogicalBackupRestore",
					},
				},
			},
			Hooks: &RestoreHooks{
				PreRestore: []HookInfo{
					{
						Name: "cleanup-old-databases",
						HookTemplate: &kmapi.ObjectReference{
							Name:      "mysql-query-executor",
							Namespace: "demo",
						},
					},
				},
				PostRestore: []HookInfo{
					{
						Name: "run-migration",
						HookTemplate: &kmapi.ObjectReference{
							Name:      "mysql-query-executor",
							Namespace: "demo",
						},
					},
				},
			},
		},
	}

	for _, fn := range transformFuncs {
		fn(rs)
	}

	return rs
}

func setPostRestoreHooksExecutionSucceededConditionToTrue(rs *RestoreSession) {
	newCond := kmapi.Condition{
		Type:    TypePostRestoreHooksExecutionSucceeded,
		Status:  v1.ConditionTrue,
		Reason:  ReasonSuccessfullyExecutedPostRestoreHooks,
		Message: "Post-Restore Hooks have been executed successfully.",
	}
	rs.Status.Conditions = kmapi.SetCondition(rs.Status.Conditions, newCond)
}