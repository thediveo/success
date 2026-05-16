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

package pass

import (
	"errors"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("passing without errors", func() {

	It("fails with a spec-supplied description", func() {
		Expect(InterceptGomegaFailure(func() {
			_ = Passing(func() (int, error) { return 0, errors.New("JKL305") }()).Bar("not enough foo: %d", 42)
		})).To(MatchError(MatchRegexp(`not enough foo: 42\n.*\n.*\n.*JKL305`)))

		Expect(InterceptGomegaFailure(func() {
			_, _ = Passing2(func() (int, string, error) { return 0, "", errors.New("JKL305") }()).Bar("not enough foo: %d", 42)
		})).To(MatchError(MatchRegexp(`not enough foo: 42\n.*\n.*\n.*JKL305`)))

		Expect(InterceptGomegaFailure(func() {
			_, _, _ = Passing3(func() (int, string, float32, error) { return 0, "", 0, errors.New("JKL305") }()).Bar("not enough foo: %d", 42)
		})).To(MatchError(MatchRegexp(`not enough foo: 42\n.*\n.*\n.*JKL305`)))
	})

	It("succeeds", func() {
		bar := Passing(func() (int, error) { return 42, nil }()).Bar("not enough foo")
		Expect(bar).To(Equal(42))

		bar1, bar2 := Passing2(func() (int, string, error) { return 42, "abc", nil }()).Bar("not enough foo")
		Expect(bar1).To(Equal(42))
		Expect(bar2).To(Equal("abc"))

		bar1, bar2, bar3 := Passing3(func() (int, string, float32, error) { return 42, "abc", 3.14, nil }()).Bar("not enough foo")
		Expect(bar1).To(Equal(42))
		Expect(bar2).To(Equal("abc"))
		Expect(bar3).To(BeNumerically("~", float32(3.14)))
	})

})
