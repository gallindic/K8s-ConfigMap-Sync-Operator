package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ConfigMapSyncSpec defines the desired state of ConfigMapSync
type ConfigMapSyncSpec struct {
	SourceNamespace      string `json:"sourceNamespace"`
	DestinationNamespace string `json:"destinationNamespace"`
	ConfigMapName        string `json:"configMapName"`
}

// ConfigMapSyncStatus defines the observed state of ConfigMapSync
type ConfigMapSyncStatus struct {
	LastSyncTime metav1.Time `json:"lastSyncTime"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// ConfigMapSync is the Schema for the configmapsyncs API
type ConfigMapSync struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ConfigMapSyncSpec   `json:"spec,omitempty"`
	Status            ConfigMapSyncStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true
// ConfigMapSyncList contains a list of ConfigMapSync
type ConfigMapSyncList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ConfigMapSync `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ConfigMapSync{}, &ConfigMapSyncList{})
}
