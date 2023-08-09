package test

import (
	"database/sql"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/snple/kokomi/core/core"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/sqlitedialect"

	_ "github.com/mattn/go-sqlite3"
)

var _ = Describe("Test core API", Label("library"), func() {
	var db *bun.DB

	BeforeEach(func() {
		sqlite, err := sql.Open("sqlite3", ":memory:")
		Expect(err).ToNot(HaveOccurred())

		db = bun.NewDB(sqlite, sqlitedialect.New())

		err = core.CreateSchema(db)
		Expect(err).ToNot(HaveOccurred())
	})

	Context("init, start and stop core", func() {
		var cs *core.CoreService

		It("init", func(ctx SpecContext) {
			var err error
			cs, err = core.Core(db)
			Expect(err).ToNot(HaveOccurred())
		})

		It("start", func(ctx SpecContext) {
			cs.Start()
		})

		It("Stop", func(ctx SpecContext) {
			c := make(chan bool)
			go func() {
				cs.Stop()
				close(c)
			}()
			Eventually(c, "1s").Should(BeClosed())

		})
	})
})
