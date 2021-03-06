/*
Copyright 2019 the original author or authors.

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
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	"github.com/projectriff/system/pkg/apis"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

var (
	ProcessorLabelKey = GroupVersion.Group + "/processor"
)

var (
	_ apis.Resource = (*Processor)(nil)
)

// ProcessorSpec defines the desired state of Processor
type ProcessorSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Build resolves the image from a build resource. As the target build
	// produces new images, they will be automatically rolled out to the
	// processor.
	// +optional
	Build *Build `json:"build"`

	// Inputs references an ordered list of streams to bind as inputs
	Inputs []StreamBinding `json:"inputs"`
	// Outputs references an ordered list of streams to bind as outputs
	// +optional
	Outputs []StreamBinding `json:"outputs"`

	// Template pod
	// +optional
	Template *corev1.PodSpec `json:"template,omitempty"`
}

type Build struct {
	// ContainerRef references a container in this namespace.
	ContainerRef string `json:"containerRef,omitempty"`

	// FunctionRef references an application in this namespace.
	FunctionRef string `json:"functionRef,omitempty"`
}

type StreamBinding struct {
	// Stream name, from this namespace, to be bound to the processor
	Stream string `json:"stream"`

	// Alias exposes the stream under another name within the processor
	// +optional
	Alias string `json:"alias,omitempty"`
}

// ProcessorStatus defines the observed state of Processor
type ProcessorStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	apis.Status `json:",inline"`

	DeprecatedInputAddresses     []string `json:"inputAddresses,omitempty"`
	DeprecatedOutputAddresses    []string `json:"outputAddresses,omitempty"`
	DeprecatedOutputContentTypes []string `json:"outputContentTypes,omitempty"`
	DeploymentName               string   `json:"deploymentName,omitempty"`
	ScaledObjectName             string   `json:"scaledObjectName,omitempty"`
	LatestImage                  string   `json:"latestImage,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:categories="riff"
// +kubebuilder:printcolumn:name="Ready",type=string,JSONPath=`.status.conditions[?(@.type=="Ready")].status`
// +kubebuilder:printcolumn:name="Reason",type=string,JSONPath=`.status.conditions[?(@.type=="Ready")].reason`
// +genclient

// Processor is the Schema for the processors API
type Processor struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ProcessorSpec   `json:"spec,omitempty"`
	Status ProcessorStatus `json:"status,omitempty"`
}

func (*Processor) GetGroupVersionKind() schema.GroupVersionKind {
	return SchemeGroupVersion.WithKind("Processor")
}

func (p *Processor) GetStatus() apis.ResourceStatus {
	return &p.Status
}

// +kubebuilder:object:root=true

// ProcessorList contains a list of Processor
type ProcessorList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Processor `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Processor{}, &ProcessorList{})
}
