package configmap

import (
	"context"
	"fmt"
	"pandax/base/global"

	v1 "k8s.io/api/core/v1"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

// ConfigMapDetail API resource provides mechanisms to inject containers with configuration data while keeping
// containers agnostic of Kubernetes
type ConfigMapDetail struct {
	// Extends list item structure.
	ConfigMap `json:",inline"`

	// Data contains the configuration data.
	// Each key must be a valid DNS_SUBDOMAIN with an optional leading dot.
	Data map[string]string `json:"data,omitempty"`
}

// GetConfigMapDetail returns detailed information about a config map
func GetConfigMapDetail(client kubernetes.Interface, namespace, name string) (*ConfigMapDetail, error) {
	global.Log.Info(fmt.Sprintf("Getting details of %s config map in %s namespace", name, namespace))

	rawConfigMap, err := client.CoreV1().ConfigMaps(namespace).Get(context.TODO(), name, metaV1.GetOptions{})

	if err != nil {
		return nil, err
	}

	return getConfigMapDetail(rawConfigMap), nil
}

func getConfigMapDetail(rawConfigMap *v1.ConfigMap) *ConfigMapDetail {
	return &ConfigMapDetail{
		ConfigMap: toConfigMap(rawConfigMap.ObjectMeta),
		Data:      rawConfigMap.Data,
	}
}
