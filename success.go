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
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

// Successful takes a return value together with an error return value,
// returning only the value (without the error return value) and at the same
// time asserting that the error return value is nil.
func Successful[R any](r R, err error) R {
	GinkgoHelper()
	Expect(err).NotTo(HaveOccurred())
	return r
}

// Successful2R takes two return values together with an error return value,
// returning only both value (without the error return value) and at the same
// time asserting that the error return value is nil.
func Successful2R[R1 any, R2 any](r1 R1, r2 R2, err error) (R1, R2) {
	GinkgoHelper()
	Expect(err).NotTo(HaveOccurred())
	return r1, r2
}

// Successful2R takes two return values together with an error return value,
// returning only both value (without the error return value) and at the same
// time asserting that the error return value is nil.
func Successful3R[R1 any, R2 any, R3 any](r1 R1, r2 R2, r3 R3, err error) (R1, R2, R3) {
	GinkgoHelper()
	Expect(err).NotTo(HaveOccurred())
	return r1, r2, r3
}
