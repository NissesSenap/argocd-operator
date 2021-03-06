// +build !ignore_autogenerated

// Copyright 2020 ArgoCD Operator Developers
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// 	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by openapi-gen. DO NOT EDIT.

// This file was autogenerated by openapi-gen. Do not edit it manually!

package v1alpha1

import (
	spec "github.com/go-openapi/spec"
	common "k8s.io/kube-openapi/pkg/common"
)

func GetOpenAPIDefinitions(ref common.ReferenceCallback) map[string]common.OpenAPIDefinition {
	return map[string]common.OpenAPIDefinition{
		"./pkg/apis/argoproj/v1alpha1.ArgoCD":             schema_pkg_apis_argoproj_v1alpha1_ArgoCD(ref),
		"./pkg/apis/argoproj/v1alpha1.ArgoCDExport":       schema_pkg_apis_argoproj_v1alpha1_ArgoCDExport(ref),
		"./pkg/apis/argoproj/v1alpha1.ArgoCDExportSpec":   schema_pkg_apis_argoproj_v1alpha1_ArgoCDExportSpec(ref),
		"./pkg/apis/argoproj/v1alpha1.ArgoCDExportStatus": schema_pkg_apis_argoproj_v1alpha1_ArgoCDExportStatus(ref),
		"./pkg/apis/argoproj/v1alpha1.ArgoCDSpec":         schema_pkg_apis_argoproj_v1alpha1_ArgoCDSpec(ref),
		"./pkg/apis/argoproj/v1alpha1.ArgoCDStatus":       schema_pkg_apis_argoproj_v1alpha1_ArgoCDStatus(ref),
	}
}

func schema_pkg_apis_argoproj_v1alpha1_ArgoCD(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "ArgoCD is the Schema for the argocds API",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("./pkg/apis/argoproj/v1alpha1.ArgoCDSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("./pkg/apis/argoproj/v1alpha1.ArgoCDStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"./pkg/apis/argoproj/v1alpha1.ArgoCDSpec", "./pkg/apis/argoproj/v1alpha1.ArgoCDStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_argoproj_v1alpha1_ArgoCDExport(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "ArgoCDExport is the Schema for the argocdexports API",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("./pkg/apis/argoproj/v1alpha1.ArgoCDExportSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("./pkg/apis/argoproj/v1alpha1.ArgoCDExportStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"./pkg/apis/argoproj/v1alpha1.ArgoCDExportSpec", "./pkg/apis/argoproj/v1alpha1.ArgoCDExportStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_argoproj_v1alpha1_ArgoCDExportSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "ArgoCDExportSpec defines the desired state of ArgoCDExport",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"argocd": {
						SchemaProps: spec.SchemaProps{
							Description: "Argocd is the name of the ArgoCD instance to export.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"image": {
						SchemaProps: spec.SchemaProps{
							Description: "Image is the container image to use for the export Job.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"schedule": {
						SchemaProps: spec.SchemaProps{
							Description: "Schedule in Cron format, see https://en.wikipedia.org/wiki/Cron.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"storage": {
						SchemaProps: spec.SchemaProps{
							Description: "Storage defines the storage configuration options.",
							Ref:         ref("./pkg/apis/argoproj/v1alpha1.ArgoCDExportStorageSpec"),
						},
					},
					"version": {
						SchemaProps: spec.SchemaProps{
							Description: "Version is the tag/digest to use for the export Job container image.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
				},
				Required: []string{"argocd"},
			},
		},
		Dependencies: []string{
			"./pkg/apis/argoproj/v1alpha1.ArgoCDExportStorageSpec"},
	}
}

func schema_pkg_apis_argoproj_v1alpha1_ArgoCDExportStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "ArgoCDExportStatus defines the observed state of ArgoCDExport",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"phase": {
						SchemaProps: spec.SchemaProps{
							Description: "Phase is a simple, high-level summary of where the ArgoCDExport is in its lifecycle. There are five possible phase values: Pending: The ArgoCDExport has been accepted by the Kubernetes system, but one or more of the required resources have not been created. Running: All of the containers for the ArgoCDExport are still running, or in the process of starting or restarting. Succeeded: All containers for the ArgoCDExport have terminated in success, and will not be restarted. Failed: At least one container has terminated in failure, either exited with non-zero status or was terminated by the system. Unknown: For some reason the state of the ArgoCDExport could not be obtained.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
				},
				Required: []string{"phase"},
			},
		},
	}
}

func schema_pkg_apis_argoproj_v1alpha1_ArgoCDSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "ArgoCDSpec defines the desired state of ArgoCD",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"applicationInstanceLabelKey": {
						SchemaProps: spec.SchemaProps{
							Description: "ApplicationInstanceLabelKey is the key name where Argo CD injects the app name as a tracking label.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"configManagementPlugins": {
						SchemaProps: spec.SchemaProps{
							Description: "ConfigManagementPlugins is used to specify additional config management plugins.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"controller": {
						SchemaProps: spec.SchemaProps{
							Description: "Controller defines the Application Controller options for ArgoCD.",
							Ref:         ref("./pkg/apis/argoproj/v1alpha1.ArgoCDApplicationControllerSpec"),
						},
					},
					"dex": {
						SchemaProps: spec.SchemaProps{
							Description: "Dex defines the Dex server options for ArgoCD.",
							Ref:         ref("./pkg/apis/argoproj/v1alpha1.ArgoCDDexSpec"),
						},
					},
					"gaTrackingID": {
						SchemaProps: spec.SchemaProps{
							Description: "GATrackingID is the google analytics tracking ID to use.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"gaAnonymizeUsers": {
						SchemaProps: spec.SchemaProps{
							Description: "GAAnonymizeUsers toggles user IDs being hashed before sending to google analytics.",
							Type:        []string{"boolean"},
							Format:      "",
						},
					},
					"grafana": {
						SchemaProps: spec.SchemaProps{
							Description: "Grafana defines the Grafana server options for ArgoCD.",
							Ref:         ref("./pkg/apis/argoproj/v1alpha1.ArgoCDGrafanaSpec"),
						},
					},
					"ha": {
						SchemaProps: spec.SchemaProps{
							Description: "HA options for High Availability support for the Redis component.",
							Ref:         ref("./pkg/apis/argoproj/v1alpha1.ArgoCDHASpec"),
						},
					},
					"helpChatURL": {
						SchemaProps: spec.SchemaProps{
							Description: "HelpChatURL is the URL for getting chat help, this will typically be your Slack channel for support.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"helpChatText": {
						SchemaProps: spec.SchemaProps{
							Description: "HelpChatText is the text for getting chat help, defaults to \"Chat now!\"",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"image": {
						SchemaProps: spec.SchemaProps{
							Description: "Image is the ArgoCD container image for all ArgoCD components.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"import": {
						SchemaProps: spec.SchemaProps{
							Description: "Import is the import/restore options for ArgoCD.",
							Ref:         ref("./pkg/apis/argoproj/v1alpha1.ArgoCDImportSpec"),
						},
					},
					"initialRepositories": {
						SchemaProps: spec.SchemaProps{
							Description: "InitialRepositories to configure Argo CD with upon creation of the cluster.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"initialSSHKnownHosts": {
						SchemaProps: spec.SchemaProps{
							Description: "InitialSSHKnownHosts defines the SSH known hosts data upon creation of the cluster for connecting Git repositories via SSH.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"kustomizeBuildOptions": {
						SchemaProps: spec.SchemaProps{
							Description: "KustomizeBuildOptions is used to specify build options/parameters to use with `kustomize build`.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"oidcConfig": {
						SchemaProps: spec.SchemaProps{
							Description: "OIDCConfig is the OIDC configuration as an alternative to dex.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"prometheus": {
						SchemaProps: spec.SchemaProps{
							Description: "Prometheus defines the Prometheus server options for ArgoCD.",
							Ref:         ref("./pkg/apis/argoproj/v1alpha1.ArgoCDPrometheusSpec"),
						},
					},
					"rbac": {
						SchemaProps: spec.SchemaProps{
							Description: "RBAC defines the RBAC configuration for Argo CD.",
							Ref:         ref("./pkg/apis/argoproj/v1alpha1.ArgoCDRBACSpec"),
						},
					},
					"redis": {
						SchemaProps: spec.SchemaProps{
							Description: "Redis defines the Redis server options for ArgoCD.",
							Ref:         ref("./pkg/apis/argoproj/v1alpha1.ArgoCDRedisSpec"),
						},
					},
					"repo": {
						SchemaProps: spec.SchemaProps{
							Description: "Repo defines the repo server options for Argo CD.",
							Ref:         ref("./pkg/apis/argoproj/v1alpha1.ArgoCDRepoSpec"),
						},
					},
					"resourceCustomizations": {
						SchemaProps: spec.SchemaProps{
							Description: "ResourceCustomizations customizes resource behavior. Keys are in the form: group/Kind.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"resourceExclusions": {
						SchemaProps: spec.SchemaProps{
							Description: "ResourceExclusions is used to completely ignore entire classes of resource group/kinds.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"server": {
						SchemaProps: spec.SchemaProps{
							Description: "Server defines the options for the ArgoCD Server component.",
							Ref:         ref("./pkg/apis/argoproj/v1alpha1.ArgoCDServerSpec"),
						},
					},
					"statusBadgeEnabled": {
						SchemaProps: spec.SchemaProps{
							Description: "StatusBadgeEnabled toggles application status badge feature.",
							Type:        []string{"boolean"},
							Format:      "",
						},
					},
					"tls": {
						SchemaProps: spec.SchemaProps{
							Description: "TLS defines the TLS options for ArgoCD.",
							Ref:         ref("./pkg/apis/argoproj/v1alpha1.ArgoCDTLSSpec"),
						},
					},
					"usersAnonymousEnabled": {
						SchemaProps: spec.SchemaProps{
							Description: "UsersAnonymousEnabled toggles anonymous user access. The anonymous users get default role permissions specified argocd-rbac-cm.",
							Type:        []string{"boolean"},
							Format:      "",
						},
					},
					"version": {
						SchemaProps: spec.SchemaProps{
							Description: "Version is the tag to use with the ArgoCD container image for all ArgoCD components.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
				},
			},
		},
		Dependencies: []string{
			"./pkg/apis/argoproj/v1alpha1.ArgoCDApplicationControllerSpec", "./pkg/apis/argoproj/v1alpha1.ArgoCDDexSpec", "./pkg/apis/argoproj/v1alpha1.ArgoCDGrafanaSpec", "./pkg/apis/argoproj/v1alpha1.ArgoCDHASpec", "./pkg/apis/argoproj/v1alpha1.ArgoCDImportSpec", "./pkg/apis/argoproj/v1alpha1.ArgoCDPrometheusSpec", "./pkg/apis/argoproj/v1alpha1.ArgoCDRBACSpec", "./pkg/apis/argoproj/v1alpha1.ArgoCDRedisSpec", "./pkg/apis/argoproj/v1alpha1.ArgoCDRepoSpec", "./pkg/apis/argoproj/v1alpha1.ArgoCDServerSpec", "./pkg/apis/argoproj/v1alpha1.ArgoCDTLSSpec"},
	}
}

func schema_pkg_apis_argoproj_v1alpha1_ArgoCDStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "ArgoCDStatus defines the observed state of ArgoCD",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"applicationController": {
						SchemaProps: spec.SchemaProps{
							Description: "ApplicationController is a simple, high-level summary of where the Argo CD application controller component is in its lifecycle. There are five possible ApplicationController values: Pending: The Argo CD application controller component has been accepted by the Kubernetes system, but one or more of the required resources have not been created. Running: All of the required Pods for the Argo CD application controller component are in a Ready state. Failed: At least one of the  Argo CD application controller component Pods had a failure. Unknown: For some reason the state of the Argo CD application controller component could not be obtained.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"dex": {
						SchemaProps: spec.SchemaProps{
							Description: "Dex is a simple, high-level summary of where the Argo CD Dex component is in its lifecycle. There are five possible dex values: Pending: The Argo CD Dex component has been accepted by the Kubernetes system, but one or more of the required resources have not been created. Running: All of the required Pods for the Argo CD Dex component are in a Ready state. Failed: At least one of the  Argo CD Dex component Pods had a failure. Unknown: For some reason the state of the Argo CD Dex component could not be obtained.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"phase": {
						SchemaProps: spec.SchemaProps{
							Description: "Phase is a simple, high-level summary of where the ArgoCD is in its lifecycle. There are five possible phase values: Pending: The ArgoCD has been accepted by the Kubernetes system, but one or more of the required resources have not been created. Available: All of the resources for the ArgoCD are ready. Failed: At least one resource has experienced a failure. Unknown: For some reason the state of the ArgoCD phase could not be obtained.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"redis": {
						SchemaProps: spec.SchemaProps{
							Description: "Redis is a simple, high-level summary of where the Argo CD Redis component is in its lifecycle. There are five possible redis values: Pending: The Argo CD Redis component has been accepted by the Kubernetes system, but one or more of the required resources have not been created. Running: All of the required Pods for the Argo CD Redis component are in a Ready state. Failed: At least one of the  Argo CD Redis component Pods had a failure. Unknown: For some reason the state of the Argo CD Redis component could not be obtained.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"repo": {
						SchemaProps: spec.SchemaProps{
							Description: "Repo is a simple, high-level summary of where the Argo CD Repo component is in its lifecycle. There are five possible repo values: Pending: The Argo CD Repo component has been accepted by the Kubernetes system, but one or more of the required resources have not been created. Running: All of the required Pods for the Argo CD Repo component are in a Ready state. Failed: At least one of the  Argo CD Repo component Pods had a failure. Unknown: For some reason the state of the Argo CD Repo component could not be obtained.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"server": {
						SchemaProps: spec.SchemaProps{
							Description: "Server is a simple, high-level summary of where the Argo CD server component is in its lifecycle. There are five possible server values: Pending: The Argo CD server component has been accepted by the Kubernetes system, but one or more of the required resources have not been created. Running: All of the required Pods for the Argo CD server component are in a Ready state. Failed: At least one of the  Argo CD server component Pods had a failure. Unknown: For some reason the state of the Argo CD server component could not be obtained.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
				},
				Required: []string{"applicationController", "dex", "phase", "redis", "repo", "server"},
			},
		},
	}
}
