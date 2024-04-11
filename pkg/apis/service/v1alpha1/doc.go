// SPDX-FileCopyrightText: 2024 SAP SE or an SAP affiliate company and Gardener contributors
//
// SPDX-License-Identifier: Apache-2.0

// +k8s:deepcopy-gen=package
// +k8s:conversion-gen=github.com/gardener/gardener-extension-falco/pkg/apis/service
// +k8s:openapi-gen=true
// +k8s:defaulter-gen=TypeMeta

//go:generate gen-crd-api-reference-docs -api-dir . -config ../../../../hack/api-reference/service.json -template-dir $GARDENER_HACK_DIR/api-reference/template -out-file ../../../../hack/api-reference/service.md

// Package v1alpha1 contains the Falco extension.
// +groupName=falco.extensions.gardener.cloud
package v1alpha1 // import "github.com/gardener/gardener-extension-falco/pkg/apis/service/v1alpha1"
