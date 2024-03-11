/*
Copyright 2016 The Kubernetes Authors.

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

// +k8s:deepcopy-gen=package

// This is a test package.
package generics

import (
	"k8s.io/gengo/examples/deepcopy-gen/output_tests/structs"
)

type TunusedGeneric[T any] struct {
	Inner1 structs.Inner
	Inner2 structs.Inner
}

type TunusedGenericMulti[T any, U any, V any] struct {
	Inner1 structs.Inner
	Inner2 structs.Inner
}
