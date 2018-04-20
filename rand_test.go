package rdf2go_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/kiivihal/rdf2go"
)

var _ = Describe("Rand", func() {

	Describe("when calling RandStringBytesMask", func() {

		Context("a single time", func() {

			It("it should have a fixed length", func() {
				t := RandStringBytesMask(10)
				Expect(t).ToNot(BeEmpty())
				Expect(t).To(HaveLen(10))
			})
		})

		Context("multiple times", func() {
			It("should be unique", func() {

				t1 := RandStringBytesMask(10)
				Expect(t1).ToNot(BeEmpty())
				Expect(t1).To(HaveLen(10))

				t2 := RandStringBytesMask(10)
				Expect(t2).ToNot(BeEmpty())
				Expect(t2).To(HaveLen(10))
				Expect(t1).ToNot(Equal(t2))

			})
		})
	})
})
