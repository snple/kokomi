package kokomi

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestKokomi(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Kokomi Suite")
}
