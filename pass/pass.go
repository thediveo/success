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
	. "github.com/onsi/ginkgo/v2" //nolint // that's fine
	. "github.com/onsi/gomega"    //nolint // that's fine
)

// Passing takes a return value together with an error return value, returning a
// value on which its Bar method has to be called next, with a spec-supplied
// description that is shown in case the error value is not nil and thus spec
// failed.
//
//	bar := Passing(Foo(42)).Bar("sadly, your foo failed you")
func Passing[R any](r R, err error) result[R] {
	return result[R]{err: err, r1: r}
}

type result[R any] struct {
	err error
	r1  R
}

// Bar annotates the Passing assertion by passing either a (format) string (and
// optional inputs to format) or a “func() string”. In case the Passing
// assertion fails this assertion together with its optional inputs will be
// printed alongside the standard failure message.
func (r result[R]) Bar(annotation any, args ...any) R {
	GinkgoHelper()
	Expect(r.err).NotTo(HaveOccurred(), append([]any{annotation}, args...)...)
	return r.r1
}

// Passing2 takes two return values together with an error return value,
// returning a value on which its Bar method has to be called next, with a
// spec-supplied description that is shown in case the error value is not nil
// and thus spec failed.
//
//	bar := Passing2(Foo(42, "abc")).Bar("sadly, your foo failed you")
func Passing2[R1, R2 any](r1 R1, r2 R2, err error) result2[R1, R2] {
	return result2[R1, R2]{err: err, r1: r1, r2: r2}
}

type result2[R1, R2 any] struct {
	err error
	r1  R1
	r2  R2
}

// Bar annotates the Passing assertion by passing either a (format) string (and
// optional inputs to format) or a “func() string”. In case the Passing
// assertion fails this assertion together with its optional inputs will be
// printed alongside the standard failure message.
func (r result2[R1, R2]) Bar(annotation any, args ...any) (R1, R2) {
	GinkgoHelper()
	Expect(r.err).NotTo(HaveOccurred(), append([]any{annotation}, args...)...)
	return r.r1, r.r2
}

// Passing3 takes three return values together with an error return value,
// returning a value on which its Bar method has to be called next, with a
// spec-supplied description that is shown in case the error value is not nil
// and thus spec failed.
//
//	bar := Passing3(Foo(42, "abc", 3.14)).Bar("sadly, your foo failed you")
func Passing3[R1, R2, R3 any](r1 R1, r2 R2, r3 R3, err error) result3[R1, R2, R3] {
	return result3[R1, R2, R3]{err: err, r1: r1, r2: r2, r3: r3}
}

type result3[R1, R2, R3 any] struct {
	err error
	r1  R1
	r2  R2
	r3  R3
}

// Bar annotates the Passing assertion by passing either a (format) string (and
// optional inputs to format) or a “func() string”. In case the Passing
// assertion fails this assertion together with its optional inputs will be
// printed alongside the standard failure message.
func (r result3[R1, R2, R3]) Bar(annotation any, args ...any) (R1, R2, R3) {
	GinkgoHelper()
	Expect(r.err).NotTo(HaveOccurred(), append([]any{annotation}, args...)...)
	return r.r1, r.r2, r.r3
}
