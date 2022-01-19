/*
Copyright 2021 Red Hat, Inc.

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
package authorinomanifests

import (
	"embed"

	logf "sigs.k8s.io/controller-runtime/pkg/log"
)

// Content holds authorino operator manifests
//go:embed autogenerated
var content embed.FS

func OperatorContent() ([]byte, error) {
	logf.Log.V(1).Info("Resource file", "name", "autogenerated/authorino-operator.yaml")
	return content.ReadFile("autogenerated/authorino-operator.yaml")
}
