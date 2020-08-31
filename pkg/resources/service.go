package resources

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"

	hawtiov1alpha1 "github.com/hawtio/hawtio-operator/pkg/apis/hawtio/v1alpha1"
)

//func NewServiceDefinitionForCR
func NewServiceDefinitionForCR(cr *hawtiov1alpha1.Hawtio) *corev1.Service {
	name := cr.Name

	return &corev1.Service{
		TypeMeta: metav1.TypeMeta{
			APIVersion: "v1",
			Kind:       "Service",
		},
		ObjectMeta: metav1.ObjectMeta{
			Annotations: map[string]string{
				"service.alpha.openshift.io/serving-cert-secret-name": name + "-tls-serving",
			},
			Labels: map[string]string{"app": "hawtio"},
			Name:   name,
		},
		Spec: corev1.ServiceSpec{
			Ports: []corev1.ServicePort{
				{
					Name:       name,
					Protocol:   "TCP",
					Port:       443,
					TargetPort: intstr.FromString(containerPortName),
				},
			},
			Selector:                 labelsForHawtio(name),
			SessionAffinity:          "None",
			PublishNotReadyAddresses: true,
		},
	}
}
