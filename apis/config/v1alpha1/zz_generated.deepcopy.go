//go:build !ignore_autogenerated
// +build !ignore_autogenerated

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

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GenericWebhookInfo) DeepCopyInto(out *GenericWebhookInfo) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GenericWebhookInfo.
func (in *GenericWebhookInfo) DeepCopy() *GenericWebhookInfo {
	if in == nil {
		return nil
	}
	out := new(GenericWebhookInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubeStashConfig) DeepCopyInto(out *KubeStashConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ControllerManagerConfigurationSpec.DeepCopyInto(&out.ControllerManagerConfigurationSpec)
	out.License = in.License
	out.WebhookInfo = in.WebhookInfo
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubeStashConfig.
func (in *KubeStashConfig) DeepCopy() *KubeStashConfig {
	if in == nil {
		return nil
	}
	out := new(KubeStashConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KubeStashConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LicenseOptions) DeepCopyInto(out *LicenseOptions) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LicenseOptions.
func (in *LicenseOptions) DeepCopy() *LicenseOptions {
	if in == nil {
		return nil
	}
	out := new(LicenseOptions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WebhookInfo) DeepCopyInto(out *WebhookInfo) {
	*out = *in
	out.Validating = in.Validating
	out.Mutating = in.Mutating
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WebhookInfo.
func (in *WebhookInfo) DeepCopy() *WebhookInfo {
	if in == nil {
		return nil
	}
	out := new(WebhookInfo)
	in.DeepCopyInto(out)
	return out
}