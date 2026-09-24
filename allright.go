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
	. "github.com/onsi/ginkgo/v2" //nolint // that's fine
	. "github.com/onsi/gomega"    //nolint // that's fine
)

// Allright takes a return value together with an “ok” boolean return value,
// asserting that it's okay, and then returning only the value.
func Allright[R any](r R, ok bool) R {
	GinkgoHelper()
	Expect(ok).To(BeTrue(), "this is not allright")
	return r
}

// Allright2R takes two return values together with an “ok” boolean return
// value, asserting that it's okay, and then returning the two return values.
func Allright2R[R1, R2 any](r1 R1, r2 R2, ok bool) (R1, R2) {
	GinkgoHelper()
	Expect(ok).To(BeTrue(), "this is not allright")
	return r1, r2
}

// Allright3R takes three return values together with an “ok” boolean return
// value, asserting that it's okay, and then returning the three return values.
func Allright3R[R1, R2, R3 any](r1 R1, r2 R2, r3 R3, ok bool) (R1, R2, R3) {
	GinkgoHelper()
	Expect(ok).To(BeTrue(), "this is not allright")
	return r1, r2, r3
}
