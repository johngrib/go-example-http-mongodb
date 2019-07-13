package db

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("DbUtil", func() {
	var db Util

	type Person struct {
		Name  string
		Hobby string
	}

	Describe("Ping", func() {
		Context("with simple call", func() {
			It("returns true", func() {
				Expect(db.Ping()).To(BeTrue())
			})
		})
	})
})
