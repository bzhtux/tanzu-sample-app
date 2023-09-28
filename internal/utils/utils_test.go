package utils_test

import (
	"os"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/bzhtux/tsa/internal/utils"
)

var _ = Describe("Utils", func() {
	utils.DefaultDBFile = "/tmp/default.db"

	Describe("Get DB Config from Env var", func() {
		It("Should return the same FAKE_DB_FILE as provided", func() {
			var FAKE_DB_FILE = "/tmp/fake.db"
			os.Create(FAKE_DB_FILE)
			os.Setenv("DB_FILE", FAKE_DB_FILE)
			defer os.Setenv("DB_FILE", "")
			defer os.Remove(FAKE_DB_FILE)
			Expect(utils.GetConfig()).To(Equal(FAKE_DB_FILE))
		})
		It("Should return default value (DefaultDBFile) as no Env var is specified", func() {
			Expect(utils.GetConfig()).To(Equal(utils.DefaultDBFile))
		})
	})
	Describe("Connect to DB FILE", func() {
		It("Should return a valid DB connection to to SQLITE DB FILE", func() {
			os.Create("/tmp/sqlite.db")
			os.Setenv("DB_FILE", "/tmp/sqlite.db")
			cfg := utils.GetConfig()
			dbConn := utils.ConnectDB(cfg)
			defer os.Remove("/tmp/sqlite.db")
			defer os.Setenv("DB_FILE", "")
			Expect(dbConn.Error).To(BeNil())
		})
		It("Should throw an error as the SQLITE DB does not exist", func() {
			var DB_FILE = "/fake/sqlite.db"
			utils.DefaultDBFile = "/fake/sqlite.db"
			os.Setenv("DB_FILE", DB_FILE)
			cfg := utils.GetConfig()
			defer os.Setenv("DB_FILE", "")
			dbConn := utils.ConnectDB(cfg)
			Expect(dbConn).To(BeNil())
		})
	})
})
