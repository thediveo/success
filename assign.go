// Copyright 2026 Harald Albrecht.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package success

import (
	"reflect"

	"github.com/onsi/gomega/gcustom"
	"github.com/onsi/gomega/types"

	. "github.com/onsi/ginkgo/v2" //nolint // that's fine
	. "github.com/onsi/gomega"    //nolint // that's fine
)

// AssignableTo asserts that v is assignable to the type T and then returns v
// cast to T; otherwise, it fails the current test.
func AssignableTo[T any](v any) T {
	GinkgoHelper()
	Expect(v).To(assignableTo[T]())
	return v.(T)
}

func assignableTo[T any]() types.GomegaMatcher {
	return gcustom.MakeMatcher(func(actual any) (bool, error) {
		_, ok := actual.(T)
		return ok, nil
	}).WithMessage("be assignable to type " + reflect.TypeFor[T]().String())
}
