package test

import (
	"fmt"
	"github.com/brianvoe/gofakeit/v6"
	. "github.com/onsi/ginkgo/v2"
	"strings"
)

type ginkgoTWrapper struct {
	GinkgoTInterface
	defaultName string
}

type WrapGinkgoTOption func(*ginkgoTWrapper)

func WrapGinkgoT(t GinkgoTInterface, opts ...WrapGinkgoTOption) *ginkgoTWrapper {

	wrapper := &ginkgoTWrapper{
		GinkgoTInterface: t,
		defaultName:      fmt.Sprintf("test-%s", strings.ToLower(gofakeit.LetterN(6))),
	}

	for _, opt := range opts {
		opt(wrapper)
	}

	return wrapper
}
