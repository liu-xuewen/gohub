package seeders

import (
	"fmt"
	"github.com/liu-xuewen/gohub/database/factories"
	"github.com/liu-xuewen/gohub/pkg/console"
	"github.com/liu-xuewen/gohub/pkg/logger"
	"github.com/liu-xuewen/gohub/pkg/seed"

	"gorm.io/gorm"
)

func init() {

	seed.Add("SeedLinksTable", func(db *gorm.DB) {

		links := factories.MakeLinks(5)

		result := db.Table("links").Create(&links)

		if err := result.Error; err != nil {
			logger.LogIf(err)
			return
		}

		console.Success(fmt.Sprintf("Table [%v] %v rows seeded", result.Statement.Table, result.RowsAffected))
	})
}
