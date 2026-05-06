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

var _ = Describe("assignments", func() {

	type X struct{}
	type Y interface{ fubar() }
	type Z struct{ Y }

	It("returns correct type expectations", func() {
		m := assignableTo[Y]()
		ok, err := m.Match(42)
		Expect(err).NotTo(HaveOccurred())
		Expect(ok).To(BeFalse())
		Expect(m.FailureMessage(42)).To(ContainSubstring("be assignable to type success.Y"))
	})

	It("returns the assignable value", func() {
		Expect(InterceptGomegaFailure(func() {
			_ = AssignableTo[Y](&X{})
		})).To(MatchError(ContainSubstring("to be assignable to type success.Y")))
	})

	It("fails for un-assignable values", func() {
		var y Y
		Expect(InterceptGomegaFailure(func() {
			y = AssignableTo[Y](&Z{})
		})).To(Succeed())
		Expect(y).NotTo(BeNil())
	})

})
