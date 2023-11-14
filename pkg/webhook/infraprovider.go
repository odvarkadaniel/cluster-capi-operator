package webhook

import (
	"context"
	"errors"
	"fmt"

	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/cluster-api-operator/api/v1alpha2"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"

	configv1 "github.com/openshift/api/config/v1"
)

type InfrastructureProviderWebhook struct {
	Platform configv1.PlatformType
}

func (r *InfrastructureProviderWebhook) SetupWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).
		WithValidator(r).
		For(&v1alpha2.InfrastructureProvider{}).
		Complete()
}

var _ webhook.CustomValidator = &InfrastructureProviderWebhook{}

// ValidateCreate implements webhook.Validator so a webhook will be registered for the type
func (r *InfrastructureProviderWebhook) ValidateCreate(ctx context.Context, obj runtime.Object) (admission.Warnings, error) {
	infraProvider, ok := obj.(*v1alpha2.InfrastructureProvider)
	if !ok {
		panic("expected to get an of object of type v1alpha2.InfrastructureProvider")
	}

	switch r.Platform {
	case configv1.AWSPlatformType:
		if infraProvider.Name != "aws" {
			return nil, fmt.Errorf("incorrect infra provider name for AWS platform: %s", infraProvider.Name)
		}
	case configv1.AzurePlatformType:
		if infraProvider.Name != "azure" {
			return nil, fmt.Errorf("incorrect infra provider name for Azure platform: %s", infraProvider.Name)
		}
	case configv1.GCPPlatformType:
		if infraProvider.Name != "gcp" {
			return nil, fmt.Errorf("incorrect infra provider name for GCP platform: %s", infraProvider.Name)
		}
	case configv1.PowerVSPlatformType:
		// for Power VS the upstream cluster api provider name is ibmcloud
		// https://github.com/kubernetes-sigs/cluster-api/blob/main/cmd/clusterctl/client/config/providers_client.go#L218-L222
		if infraProvider.Name != "ibmcloud" {
			return nil, fmt.Errorf("incorrect infra provider name for PowerVS platform: %s", infraProvider.Name)
		}
	case configv1.VSpherePlatformType:
		if infraProvider.Name != "vsphere" {
			return nil, fmt.Errorf("incorrect infra provider name for VSphere platform: %s", infraProvider.Name)
		}
	default:
		return nil, errors.New("platform not supported, skipping infra cluster controller setup")
	}

	return nil, nil
}

// ValidateUpdate implements webhook.Validator so a webhook will be registered for the type
func (r *InfrastructureProviderWebhook) ValidateUpdate(ctx context.Context, oldObj, newObj runtime.Object) (admission.Warnings, error) {
	_, ok := oldObj.(*v1alpha2.InfrastructureProvider)
	if !ok {
		panic("expected to get an of object of type v1alpha2.InfrastructureProvider")
	}
	newInfraProvider, ok := newObj.(*v1alpha2.InfrastructureProvider)
	if !ok {
		panic("expected to get an of object of type v1alpha2.InfrastructureProvider")
	}

	switch r.Platform {
	case configv1.AWSPlatformType:
		if newInfraProvider.Name != "aws" {
			return nil, fmt.Errorf("incorrect infra provider name for AWS platform: %s", newInfraProvider.Name)
		}
	case configv1.AzurePlatformType:
		if newInfraProvider.Name != "azure" {
			return nil, fmt.Errorf("incorrect infra provider name for Azure platform: %s", newInfraProvider.Name)
		}
	case configv1.GCPPlatformType:
		if newInfraProvider.Name != "gcp" {
			return nil, fmt.Errorf("incorrect infra provider name for GCP platform: %s", newInfraProvider.Name)
		}
	case configv1.PowerVSPlatformType:
		// for Power VS the upstream cluster api provider name is ibmcloud
		// https://github.com/kubernetes-sigs/cluster-api/blob/main/cmd/clusterctl/client/config/providers_client.go#L218-L222
		if newInfraProvider.Name != "ibmcloud" {
			return nil, fmt.Errorf("incorrect infra provider name for PowerVS platform: %s", newInfraProvider.Name)
		}
	case configv1.VSpherePlatformType:
		if newInfraProvider.Name != "vsphere" {
			return nil, fmt.Errorf("incorrect infra provider name for VSphere platform: %s", newInfraProvider.Name)
		}
	default:
		return nil, errors.New("platform not supported, skipping infra cluster controller setup")
	}

	return nil, nil
}

// ValidateDelete implements webhook.Validator so a webhook will be registered for the type
func (r *InfrastructureProviderWebhook) ValidateDelete(_ context.Context, obj runtime.Object) (admission.Warnings, error) {
	return nil, errors.New("deletion of infrastructure provider is not allowed")
}
