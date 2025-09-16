package service_test

import (
	"context"
	"math/rand"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/sirupsen/logrus"
	"github.com/zenoss/zenkit/v5"
)

func TestService(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "{{Name | title}} Service Suite", AroundNode(func(ctx context.Context) context.Context {
		log := logrus.NewEntry(logrus.New())
		log.Logger.Out = GinkgoWriter
		log.Logger.Level = logrus.DebugLevel
		ctx = zenkit.LoggerToContext(ctx, log)
		return ctx
	}))
}
