/*
Copyright 2020 The Crossplane Authors.

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

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

// A ConfigurationSource represents the source of a Playbook Configuration.
// +kubebuilder:validation:Enum=Remote;Inline
type ConfigurationSource string

// Module sources.
const (
	ConfigurationSourceRemote ConfigurationSource = "Remote"
	ConfigurationSourceInline ConfigurationSource = "Inline"
)

// PlaybookParameters are the configurable fields of a Playbook.
type PlaybookParameters struct {
	// The configuration of this playbook; i.e. the configuration containing its playbook.yml
	// file. When the playbook's ProviderCredentialssource is 'Remote' (the default) this can be
	// any address supported by Ansible.Builtin.git,
	// TODO support other remotes https://docs.ansible.com/ansible/latest/collections/ansible/builtin/index.html
	// When the playbook's source is 'Inline' the
	// content of a simple playbook.yml file may be written inline.
	Configuration string `json:"module"`

	// Source of configuration of this playbook.
	Source ConfigurationSource `json:"source"`
}

// PlaybookObservation are the observable fields of a Playbook.
type PlaybookObservation struct {
	// TODO(negz): Should we include outputs here? Or only in connection
	// details.
}

// A PlaybookSpec defines the desired state of a Playbook.
type PlaybookSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       PlaybookParameters `json:"forProvider"`
}

// A PlaybookStatus represents the observed state of a Workspace.
type PlaybookStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          PlaybookObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// A Playbook of Ansible Configuration.
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
type Playbook struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   PlaybookSpec   `json:"spec"`
	Status PlaybookStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PlaybookList contains a list of Playbook
type PlaybookList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Playbook `json:"items"`
}
