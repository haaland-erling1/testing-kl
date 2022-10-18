package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	ct "operators.kloudlite.io/apis/common-types"
	"operators.kloudlite.io/lib/constants"
	rApi "operators.kloudlite.io/lib/operator"
)

type Expose struct {
	Enabled         bool   `json:"enabled"`
	Domain          string `json:"domain"`
	BasicAuthSecret string `json:"basicAuthSecret"`
}

// KibanaSpec defines the desired state of Kibana
type KibanaSpec struct {
	// +kubebuilder:default=1
	ReplicaCount int          `json:"replicaCount,omitempty"`
	Region       string       `json:"region"`
	Resources    ct.Resources `json:"resources"`
	ElasticUrl   string       `json:"elasticUrl"`
	Expose       Expose       `json:"expose"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:JSONPath=".status.isReady",name=Ready,type=boolean
// +kubebuilder:printcolumn:JSONPath=".metadata.creationTimestamp",name=Age,type=date

// Kibana is the Schema for the kibanas API
type Kibana struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   KibanaSpec  `json:"spec,omitempty"`
	Status rApi.Status `json:"status,omitempty"`
}

func (k *Kibana) GetStatus() *rApi.Status {
	return &k.Status
}

func (k *Kibana) GetEnsuredLabels() map[string]string {
	return map[string]string{
		constants.MsvcNameKey: k.Name,
	}
}

func (k *Kibana) GetEnsuredAnnotations() map[string]string {
	return map[string]string{}
}

// +kubebuilder:object:root=true

// KibanaList contains a list of Kibana
type KibanaList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Kibana `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Kibana{}, &KibanaList{})
}
