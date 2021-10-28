/*
Copyright AppsCode Inc. and Contributors

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

// Code generated by Kubeform. DO NOT EDIT.

package v1alpha1

import (
	base "kubeform.dev/apimachinery/api/v1alpha1"

	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	kmapi "kmodules.xyz/client-go/api/v1"
	"sigs.k8s.io/cli-utils/pkg/kstatus/status"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Phase",type=string,JSONPath=`.status.phase`

type Window struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              WindowSpec   `json:"spec,omitempty"`
	Status            WindowStatus `json:"status,omitempty"`
}

type WindowSpec struct {
	State *WindowSpecResource `json:"state,omitempty" tf:"-"`

	Resource WindowSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type WindowSpecResource struct {
	ID string `json:"id,omitempty" tf:"id,omitempty"`

	EndTimeInSeconds *int64 `json:"endTimeInSeconds" tf:"end_time_in_seconds"`
	// +optional
	HostTagGroupHostNamesGroupAnded *bool   `json:"hostTagGroupHostNamesGroupAnded,omitempty" tf:"host_tag_group_host_names_group_anded"`
	Reason                          *string `json:"reason" tf:"reason"`
	// +optional
	RelevantCustomerTags []string `json:"relevantCustomerTags,omitempty" tf:"relevant_customer_tags"`
	// +optional
	RelevantHostNames []string `json:"relevantHostNames,omitempty" tf:"relevant_host_names"`
	// +optional
	RelevantHostTags []string `json:"relevantHostTags,omitempty" tf:"relevant_host_tags"`
	// +optional
	RelevantHostTagsAnded *bool   `json:"relevantHostTagsAnded,omitempty" tf:"relevant_host_tags_anded"`
	StartTimeInSeconds    *int64  `json:"startTimeInSeconds" tf:"start_time_in_seconds"`
	Title                 *string `json:"title" tf:"title"`
}

type WindowStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Phase status.Status `json:"phase,omitempty"`
	// +optional
	Conditions []kmapi.Condition `json:"conditions,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// WindowList is a list of Windows
type WindowList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Window CRD objects
	Items []Window `json:"items,omitempty"`
}
