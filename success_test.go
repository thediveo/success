// Copyright 2023 Harald Albrecht.
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
	"errors"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestSuccess(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "success package")
}

var _ = Describe("ensuring success", func() {

	It("bails out where there is no success", func() {
		msgs := InterceptGomegaFailures(func() {
			_ = Successful(func() (int, error) { return 0, errors.New("D'OH!") }())
		})
		Expect(msgs).To(ConsistOf(MatchRegexp(`(?s)Unexpected error:.*D'OH!`)))

		msgs = InterceptGomegaFailures(func() {
			_, _ = Successful2R(func() (int, string, error) { return 0, "uhoh", errors.New("D'OH!") }())
		})
		Expect(msgs).To(ConsistOf(MatchRegexp(`(?s)Unexpected error:.*D'OH!`)))

		msgs = InterceptGomegaFailures(func() {
			_, _, _ = Successful3R(func() (int, string, bool, error) { return 0, "", false, errors.New("D'OH!") }())
		})
		Expect(msgs).To(ConsistOf(MatchRegexp(`(?s)Unexpected error:.*D'OH!`)))
	})

	It("succeeds", func() {
		answer := Successful(func() (int, error) { return 42, nil }())
		Expect(answer).To(Equal(42))

		answer, answer2 := Successful2R(func() (int, string, error) { return 42, "foo", nil }())
		Expect(answer).To(Equal(42))
		Expect(answer2).To(Equal("foo"))

		answer, answer2, answer3 := Successful3R(func() (int, string, bool, error) { return 42, "foo", true, nil }())
		Expect(answer).To(Equal(42))
		Expect(answer2).To(Equal("foo"))
		Expect(answer3).To(BeTrue())
	})

})
