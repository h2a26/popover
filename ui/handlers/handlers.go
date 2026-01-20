package handlers

import (
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"

	"mpay/postgres"
	"mpay/ui/handlers/customer"
	"mpay/ui/handlers/login"
	"mpay/ui/handlers/url"
)

func RegisterRoutes(e *echo.Echo, db *postgres.Connection, log *zap.Logger) {
	{
		e.GET(url.Root, login.Login)
		e.POST(url.Login, login.PostLogin)
	}

	{
		handler := customer.NewHandler(db, log)
		e.GET(url.Customer_ListPage, handler.GetListPage)
		e.GET(url.Customer_ConsumersList, handler.GetConsumersList)
		e.GET(url.Customer_MerchantsList, handler.GetMerchantsList)
		e.GET(url.Customer_AgentsList, handler.GetAgentsList)
		e.GET(url.Customer_ConsumerNewPage, handler.GetConsumerNewPage)
		e.POST(url.Customer_ConsumerNewPage, handler.PostConsumerNewPage)
	}
}
