package main

import (
	"fmt"
	k8sType "k8s.io/api/apps/v1beta2"
	k8sCore "k8s.io/api/core/v1"
	k8sMeta "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
	"k8s.io/client-go/kubernetes"
	k8sClientApps "k8s.io/client-go/kubernetes/typed/apps/v1beta2"
	k8sClientCore "k8s.io/client-go/kubernetes/typed/core/v1"
)

type Deployer struct {
	deployments k8sClientApps.DeploymentInterface
	services    k8sClientCore.ServiceInterface
}

func NewDeployer(clientset *kubernetes.Clientset) *Deployer {
	namespace := "default"
	return &Deployer{
		deployments: clientset.AppsV1beta2().Deployments(namespace),
		services:    clientset.Core().Services(namespace),
	}
}

func (d *Deployer) Deploy(project string) error {
	if err := d.performDeploy(project, fmt.Sprintf("%s:v1", project)); err != nil {
		return err
	}

	return d.exposeService(project)
}

func (d *Deployer) performDeploy(project string, imageName string) error {
	replicas := int32(1)

	_, err := d.deployments.Create(&k8sType.Deployment{
		ObjectMeta: k8sMeta.ObjectMeta{
			Name:   project,
			Labels: map[string]string{"app": project},
		},
		Spec: k8sType.DeploymentSpec{
			Replicas: &replicas,
			Selector: &k8sMeta.LabelSelector{
				MatchLabels: map[string]string{"app": project},
			},
			Template: k8sCore.PodTemplateSpec{
				ObjectMeta: k8sMeta.ObjectMeta{
					Labels: map[string]string{"app": project},
				},
				Spec: k8sCore.PodSpec{
					Containers: []k8sCore.Container{
						k8sCore.Container{
							Name:  project,
							Image: imageName,
							Ports: []k8sCore.ContainerPort{
								k8sCore.ContainerPort{
									ContainerPort: int32(8080),
								},
							},
						},
					},
				},
			},
		},
	})

	return err
}

func (d *Deployer) exposeService(project string) error {
	_, err := d.services.Create(&k8sCore.Service{
		// TODO TypeMeta:
		ObjectMeta: k8sMeta.ObjectMeta{
			Name:   project,
			Labels: map[string]string{"app": project},
		},
		Spec: k8sCore.ServiceSpec{
			Ports: []k8sCore.ServicePort{
				k8sCore.ServicePort{
					Port: int32(8080),
					TargetPort: intstr.IntOrString{
						Type:   intstr.Int,
						IntVal: 8080,
					},
				},
			},
			Selector: map[string]string{"app": project},
			Type:     k8sCore.ServiceTypeLoadBalancer,
		},
	})

	return err
}
