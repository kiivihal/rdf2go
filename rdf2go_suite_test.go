package rdf2go_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestRdf2go(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Rdf2go Suite")
}
