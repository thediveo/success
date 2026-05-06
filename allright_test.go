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
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("asserting everything's allright", func() {

	It("asserts it's okay", func() {
		Expect(InterceptGomegaFailure(func() {
			var s string = Allright(func() (string, bool) { return "", true }()) //nolint:staticcheck
			_ = s
		})).To(Succeed())
	})

	It("fails when it's not okay", func() {
		Expect(InterceptGomegaFailure(func() {
			var s string = Allright(func() (string, bool) { return "", false }()) //nolint:staticcheck
			_ = s
		})).To(HaveOccurred())
	})

})
