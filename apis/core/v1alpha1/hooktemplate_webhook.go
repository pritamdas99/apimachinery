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
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
)

// log is for logging in this package.
var hooktemplatelog = logf.Log.WithName("hooktemplate-resource")

func (r *HookTemplate) SetupWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).
		For(r).
		Complete()
}

// TODO(user): EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!

//+kubebuilder:webhook:path=/mutate-core-kubestash-com-v1alpha1-hooktemplate,mutating=true,failurePolicy=fail,sideEffects=None,groups=core.kubestash.com,resources=hooktemplates,verbs=create;update,versions=v1alpha1,name=mhooktemplate.kb.io,admissionReviewVersions=v1

var _ webhook.Defaulter = &HookTemplate{}

// Default implements webhook.Defaulter so a webhook will be registered for the type
func (r *HookTemplate) Default() {
	hooktemplatelog.Info("default", "name", r.Name)

	// TODO(user): fill in your defaulting logic.
}

// TODO(user): change verbs to "verbs=create;update;delete" if you want to enable deletion validation.
//+kubebuilder:webhook:path=/validate-core-kubestash-com-v1alpha1-hooktemplate,mutating=false,failurePolicy=fail,sideEffects=None,groups=core.kubestash.com,resources=hooktemplates,verbs=create;update,versions=v1alpha1,name=vhooktemplate.kb.io,admissionReviewVersions=v1

var _ webhook.Validator = &HookTemplate{}

// ValidateCreate implements webhook.Validator so a webhook will be registered for the type
func (r *HookTemplate) ValidateCreate() error {
	hooktemplatelog.Info("validate create", "name", r.Name)

	// TODO(user): fill in your validation logic upon object creation.
	return nil
}

// ValidateUpdate implements webhook.Validator so a webhook will be registered for the type
func (r *HookTemplate) ValidateUpdate(old runtime.Object) error {
	hooktemplatelog.Info("validate update", "name", r.Name)

	// TODO(user): fill in your validation logic upon object update.
	return nil
}

// ValidateDelete implements webhook.Validator so a webhook will be registered for the type
func (r *HookTemplate) ValidateDelete() error {
	hooktemplatelog.Info("validate delete", "name", r.Name)

	// TODO(user): fill in your validation logic upon object deletion.
	return nil
}