/*
Copyright 2023.

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

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// MemcacheSpec defines the desired state of Memcache
type MemcacheSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Replicas is an example field of Memcache. Edit memcache_types.go to remove/update
	Replicas int32  `json:"replicas"`
	Image    string `json:"image"`
}

// MemcacheStatus defines the observed state of Memcache
type MemcacheStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	State   string `json:"state"`
	Message string `json:"message,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
//+kubebuilder:object:root=true
//+kubebuilder:subresource:status
//+kubebuilder:resource:scope=Namespaced,shortName=mc
//+kubebuilder:printcolumn:name=Replicas,type=string,JSONPath=".spec.replicas"
//+kubebuilder:printcolumn:name=Image,type=string,JSONPath=".spec.image"
//+kubebuilder:printcolumn:name=State,type=string,JSONPath=".status.state"
//+kubebuilder:printcolumn:name="Age",type=date,JSONPath=.metadata.creationTimestamp

// Memcache is the Schema for the memcaches API
type Memcache struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   MemcacheSpec   `json:"spec,omitempty"`
	Status MemcacheStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
//+kubebuilder:object:root=true

// MemcacheList contains a list of Memcache
type MemcacheList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Memcache `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Memcache{}, &MemcacheList{})
}
