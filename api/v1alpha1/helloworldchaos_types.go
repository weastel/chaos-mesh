package v1alpha1

import (
	"k8s.io/apimachinery/pkg/runtime"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +kubebuilder:object:root=true

// HelloWorldChaos is the Schema for the helloworldchaos API
type HelloWorldChaos struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
}

func (in *HelloWorldChaos) DeepCopyObject() runtime.Object {
	out := HelloWorldChaos{}
	out.TypeMeta = in.TypeMeta
	out.ObjectMeta = in.ObjectMeta

	return &out
}

// +kubebuilder:object:root=true

// HelloWorldChaosList contains a list of HelloWorldChaos
type HelloWorldChaosList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []HelloWorldChaos `json:"items"`
}

func (in *HelloWorldChaosList) DeepCopyObject() runtime.Object {
	out := HelloWorldChaosList{}
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	out.Items = in.Items

	return &out
}

func (in *HelloWorldChaosList) DeepCopyInto(out *HelloWorldChaosList) {
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	out.Items = in.Items
}

func init() {
	SchemeBuilder.Register(&HelloWorldChaos{}, &HelloWorldChaosList{})
}
