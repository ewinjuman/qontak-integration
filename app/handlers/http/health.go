package http

import (
	"github.com/dimiro1/health"
	"github.com/dimiro1/health/db"
	"github.com/dimiro1/health/redis"
	"github.com/gofiber/fiber/v2"
	FiberUtils "github.com/gofiber/fiber/v2/utils"
	"github.com/pbnjay/memory"
	"qontak_integration/pkg/base"
	"qontak_integration/pkg/configs"
	"qontak_integration/platform/database"
)

func HealthCheck(c *fiber.Ctx) error {
	ctx := base.NewContext(c)

	dbConn, _ := database.MysqlConnection(ctx.Session)
	dba, _ := dbConn.DB()
	mysql := db.NewMySQLChecker(dba)
	handler := health.NewHandler()
	handler.AddChecker("database", mysql)

	handler.AddChecker("redis", redis.NewChecker("tcp", configs.Config.Redis.Address))

	resp := handler.CompositeChecker.Check()
	resp.AddInfo("freeMemory", FiberUtils.ByteSize(memory.FreeMemory()))

	// Return status 200 OK.
	return ctx.JSON(resp)
}
