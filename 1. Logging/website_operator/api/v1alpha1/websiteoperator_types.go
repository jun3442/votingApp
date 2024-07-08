/*
Copyright 2024.

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
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// WebsiteOperatorSpec defines the desired state of WebsiteOperator
type WebsiteOperatorSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Foo is an example field of WebsiteOperator. Edit websiteoperator_types.go to remove/update
	Foo string `json:"foo,omitempty"`
}

// WebsiteOperatorStatus defines the observed state of WebsiteOperator
type WebsiteOperatorStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// WebsiteOperator is the Schema for the websiteoperators API
type WebsiteOperator struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   WebsiteOperatorSpec   `json:"spec,omitempty"`
	Status WebsiteOperatorStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// WebsiteOperatorList contains a list of WebsiteOperator
type WebsiteOperatorList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []WebsiteOperator `json:"items"`
}

func init() {
	SchemeBuilder.Register(&WebsiteOperator{}, &WebsiteOperatorList{})
}
