package test

import (
	"database/sql"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/snple/kokomi/consts"
	"github.com/snple/kokomi/core/core"
	"github.com/snple/kokomi/pb"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/sqlitedialect"
)

var _ = Describe("Test core device API", Label("library"), func() {
	var db *bun.DB
	var cs *core.CoreService

	BeforeEach(func() {
		sqlite, err := sql.Open("sqlite3", ":memory:")
		Expect(err).ToNot(HaveOccurred())

		db = bun.NewDB(sqlite, sqlitedialect.New())

		// db.AddQueryHook(bundebug.NewQueryHook(bundebug.WithVerbose(true)))

		err = core.CreateSchema(db)
		Expect(err).ToNot(HaveOccurred())

		cs, err = core.Core(db)
		Expect(err).ToNot(HaveOccurred())

		cs.Start()
	})

	BeforeEach(func() {
		c := make(chan bool)
		go func() {
			cs.Stop()
			close(c)
		}()
		Eventually(c, "1s").Should(BeClosed())
	})

	Context("device CRUD", func() {
		It("create", func(ctx SpecContext) {
			{
				request := &pb.Device{
					Name:   "test_device1",
					Desc:   "test",
					Secret: "123456",
					Status: consts.ON,
				}

				reply, err := cs.GetDevice().Create(ctx, request)

				Expect(err).ToNot(HaveOccurred())
				_ = reply
			}

			{
				request := &pb.Name{Name: "test_device1"}

				reply, err := cs.GetDevice().Name(ctx, request)

				Expect(err).ToNot(HaveOccurred())
				_ = reply
			}
		})

		It("destory", func(ctx SpecContext) {
			{
				request := &pb.Device{
					Name:   "test_device1",
					Desc:   "test",
					Secret: "123456",
					Status: consts.ON,
				}

				reply, err := cs.GetDevice().Create(ctx, request)

				Expect(err).ToNot(HaveOccurred())
				_ = reply
			}

			{
				request := &pb.Name{Name: "test_device1"}

				reply, err := cs.GetDevice().Name(ctx, request)

				Expect(err).ToNot(HaveOccurred())
				_ = reply

				{
					reply, err := cs.GetSync().GetDeviceUpdated(ctx, &pb.Id{Id: reply.GetId()})
					Expect(err).ToNot(HaveOccurred())
					Expect(reply.Updated > 0).To(Equal(true))
				}

				{
					_, err = cs.GetDevice().Destory(ctx, &pb.Id{Id: reply.GetId()})
					Expect(err).ToNot(HaveOccurred())
				}

				{
					reply, err := cs.GetSync().GetDeviceUpdated(ctx, &pb.Id{Id: reply.GetId()})
					Expect(err).ToNot(HaveOccurred())
					Expect(reply.Updated < 0).To(Equal(true))
				}
			}
		})
	})
})
