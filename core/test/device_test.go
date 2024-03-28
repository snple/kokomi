package test

import (
	"database/sql"

	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
	"github.com/snple/kokomi/consts"
	"github.com/snple/kokomi/core"
	"github.com/snple/kokomi/pb"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/sqlitedialect"
)

var _ = ginkgo.Describe("Test core device API", ginkgo.Label("library"), func() {
	var db *bun.DB
	var cs *core.CoreService

	ginkgo.BeforeEach(func() {
		sqlite, err := sql.Open("sqlite3", ":memory:")
		gomega.Expect(err).ToNot(gomega.HaveOccurred())

		db = bun.NewDB(sqlite, sqlitedialect.New())

		// db.AddQueryHook(bundebug.NewQueryHook(bundebug.WithVerbose(true)))

		err = core.CreateSchema(db)
		gomega.Expect(err).ToNot(gomega.HaveOccurred())

		cs, err = core.Core(db)
		gomega.Expect(err).ToNot(gomega.HaveOccurred())

		cs.Start()
	})

	ginkgo.AfterEach(func() {
		c := make(chan bool)
		go func() {
			cs.Stop()
			close(c)
		}()
		gomega.Eventually(c, "1s").Should(gomega.BeClosed())
	})

	ginkgo.Context("device CRUD", func() {
		ginkgo.It("create", func(ctx ginkgo.SpecContext) {
			{
				request := &pb.Device{
					Name:   "test_device1",
					Desc:   "test",
					Secret: "123456",
					Status: consts.ON,
				}

				reply, err := cs.GetDevice().Create(ctx, request)

				gomega.Expect(err).ToNot(gomega.HaveOccurred())
				_ = reply
			}

			{
				request := &pb.Name{Name: "test_device1"}

				reply, err := cs.GetDevice().Name(ctx, request)

				gomega.Expect(err).ToNot(gomega.HaveOccurred())
				_ = reply
			}
		})

		ginkgo.It("destory", func(ctx ginkgo.SpecContext) {
			{
				request := &pb.Device{
					Name:   "test_device1",
					Desc:   "test",
					Secret: "123456",
					Status: consts.ON,
				}

				reply, err := cs.GetDevice().Create(ctx, request)

				gomega.Expect(err).ToNot(gomega.HaveOccurred())
				_ = reply
			}

			{
				request := &pb.Name{Name: "test_device1"}

				reply, err := cs.GetDevice().Name(ctx, request)

				gomega.Expect(err).ToNot(gomega.HaveOccurred())
				_ = reply

				{
					reply, err := cs.GetSync().GetDeviceUpdated(ctx, &pb.Id{Id: reply.GetId()})
					gomega.Expect(err).ToNot(gomega.HaveOccurred())
					gomega.Expect(reply.Updated > 0).To(gomega.Equal(true))
				}

				{
					_, err = cs.GetDevice().Destory(ctx, &pb.Id{Id: reply.GetId()})
					gomega.Expect(err).ToNot(gomega.HaveOccurred())
				}

				{
					reply, err := cs.GetSync().GetDeviceUpdated(ctx, &pb.Id{Id: reply.GetId()})
					gomega.Expect(err).ToNot(gomega.HaveOccurred())
					gomega.Expect(reply.Updated < 0).To(gomega.Equal(true))
				}
			}
		})
	})
})
