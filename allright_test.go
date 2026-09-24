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
			var s string = Allright(func() (string, bool) { return "ok", true }())
			Expect(s).To(Equal("ok"))
		})).To(Succeed())

		Expect(InterceptGomegaFailure(func() {
			var (
				s string
				i int
			)
			s, i = Allright2R(func() (string, int, bool) { return "ok", 42, true }())
			Expect(s).To(Equal("ok"))
			Expect(i).To(Equal(42))
		})).To(Succeed())

		Expect(InterceptGomegaFailure(func() {
			var (
				s string
				i int
				f float32
			)
			s, i, f = Allright3R(func() (string, int, float32, bool) { return "ok", 42, 666.6, true }())
			Expect(s).To(Equal("ok"))
			Expect(i).To(Equal(42))
			Expect(f).To(Equal(float32(666.6)))
		})).To(Succeed())
	})

	It("fails when it's not okay", func() {
		Expect(InterceptGomegaFailure(func() {
			_ = Allright(func() (string, bool) { return "", false }())
		})).To(HaveOccurred())

		Expect(InterceptGomegaFailure(func() {
			_, _ = Allright2R(func() (string, int, bool) { return "", 42, false }())
		})).To(HaveOccurred())

		Expect(InterceptGomegaFailure(func() {
			_, _, _ = Allright3R(func() (string, int, float32, bool) { return "", 42, 666.6, false }())
		})).To(HaveOccurred())
	})

})
