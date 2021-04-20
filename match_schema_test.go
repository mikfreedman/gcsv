package gcsv_test

import (
	. "github.com/mikfreedman/gcsv"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestCustomMatcher(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Custom Matcher Suite")
}

var _ = Describe("RepresentSchema Object", func() {
	var (
		data   string
		actual []interface{}
	)

	BeforeEach(func() {
		actual = []interface{}{
			"category",
			"item",
			1.0,
			1,
			false,
		}

	})

	Context("containing the csv with a schema of actual", func() {
		BeforeEach(func() {
			data = "flowers,roses,4.0,1,false\nflowers,roses,4.0,1,false"
		})

		It("should succeed", func() {
			Ω(data).Should(RepresentSchema(actual))
		})
		Context("with headers", func() {
			BeforeEach(func() {
				data = "category,item,price,quantity,exported\nflowers,roses,4.0,1,false\nflowers,roses,4.0,1,false"
			})

			It("should succeed", func() {
				Ω(data).Should(RepresentSchema(actual, IgnoreHeaderRow()))
			})
		})
	})

	Context("containing the csv with a schema of something else", func() {
		BeforeEach(func() {
			data = "flowers,roses,4.0,1,false\nflowers,roses,4.0,1,bob"
		})

		It("should fail", func() {
			Ω(data).ShouldNot(RepresentSchema(actual))
		})
	})

	Context("when actual is not a string", func() {
		It("should error", func() {
			_, err := RepresentSchema(actual).Match(1)
			Ω(err).Should(HaveOccurred())
		})
	})

	Context("when actual is not parseable", func() {
		BeforeEach(func() {
			data = "roses,4.0,1,false\nflowers,roses,4.0,1,bob"
		})

		It("should error", func() {
			_, err := RepresentSchema(actual).Match(data)
			Ω(err).Should(HaveOccurred())
		})
	})
})
