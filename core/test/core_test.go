package test

import (
	"database/sql"

	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
	"github.com/snple/beacon/core"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/sqlitedialect"
)

var _ = ginkgo.Describe("Test core API", ginkgo.Label("library"), func() {
	var db *bun.DB

	ginkgo.BeforeEach(func() {
		sqlite, err := sql.Open("sqlite3", ":memory:")
		gomega.Expect(err).ToNot(gomega.HaveOccurred())

		db = bun.NewDB(sqlite, sqlitedialect.New())

		err = core.CreateSchema(db)
		gomega.Expect(err).ToNot(gomega.HaveOccurred())
	})

	ginkgo.Context("init, start and stop core", func() {
		var cs *core.CoreService

		ginkgo.It("init", func(ctx ginkgo.SpecContext) {
			var err error
			cs, err = core.Core(db)
			gomega.Expect(err).ToNot(gomega.HaveOccurred())
		})

		ginkgo.It("start", func(ctx ginkgo.SpecContext) {
			cs.Start()
		})

		ginkgo.It("Stop", func(ctx ginkgo.SpecContext) {
			c := make(chan bool)
			go func() {
				cs.Stop()
				close(c)
			}()
			gomega.Eventually(c, "1s").Should(gomega.BeClosed())

		})
	})
})
