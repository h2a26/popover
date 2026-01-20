package customer

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"

	"mpay/postgres"
	"mpay/ui/components/table"
	"mpay/ui/handlers/render"
	"mpay/ui/handlers/url"
	"mpay/ui/pages/customer/consumer"
	"mpay/ui/utils"
)

type Handler struct {
	db  *postgres.Connection
	log *zap.Logger
}

func NewHandler(db *postgres.Connection, log *zap.Logger) *Handler {
	return &Handler{
		db:  db,
		log: log,
	}
}

func (h *Handler) GetListPage(c echo.Context) error {
	props, err := buildConsumersListProps(c, h.db)
	if err != nil {
		return render.Render(c, http.StatusOK, consumer.ListPage(props))
	}

	return render.Render(c, http.StatusOK, consumer.ListPage(props))
}

func (h *Handler) GetConsumerNewPage(c echo.Context) error {
	props := consumer.NewPageProps{
		CurrentPath: url.Customer_ListPage,
		Form:        consumer.FormProps{},
	}

	return render.Render(c, http.StatusOK, consumer.NewPage(props))
}

func (h *Handler) PostConsumerNewPage(c echo.Context) error {
	formProps := consumer.FormProps{
		PhoneNumber: c.FormValue("phone_number"),
	}

	formProps.ShowEntry = true

	newPageProps := consumer.NewPageProps{
		CurrentPath: c.Path(),
		Form:        formProps,
	}

	if utils.IsConfirmed(c) && formProps.ShowEntry {

		return render.RedirectPage(c, url.Customer_ConsumersList)
	}

	if utils.IsHTMX(c) {
		return render.RenderSwap(c, consumer.EntryFormContainer(newPageProps))
	}

	return render.Render(c, http.StatusOK, consumer.NewPage(newPageProps))
}

func (h *Handler) GetConsumersList(c echo.Context) error {
	props, err := buildConsumersListProps(c, h.db)
	if err != nil {
		return render.Render(c, http.StatusOK, consumer.ConsumersListContainer(props))
	}

	return render.Render(c, http.StatusOK, consumer.ConsumersListContainer(props))
}

func buildConsumersListProps(c echo.Context, db *postgres.Connection) (consumer.ConsumersListProps, error) {

	total, err := selectNonDraftConsumersCount(db)
	if err != nil {
		return consumer.ConsumersListProps{}, err
	}

	page, offset, limit := utils.ParseListingPageParams(c, total)

	data, err := selectAllNonDraftConsumers(db, offset, limit)
	if err != nil {
		return consumer.ConsumersListProps{}, err
	}

	pagination := table.PaginationFrom(uint(page), total)
	pageItems := table.BuildWindowedPagination(pagination)

	return consumer.ConsumersListProps{
		CurrentPath: url.Customer_ListPage,
		Data:        data,
		Pagination:  pagination,
		PageItems:   pageItems,
	}, nil
}

func (h *Handler) GetMerchantsList(c echo.Context) error {
	props, err := buildMerchantsListProps(c, h.db)
	if err != nil {
		return render.Render(c, http.StatusOK, consumer.MerchantsListContainer(props))
	}

	return render.Render(c, http.StatusOK, consumer.MerchantsListContainer(props))
}

func buildMerchantsListProps(c echo.Context, db *postgres.Connection) (consumer.MerchantsListProps, error) {

	total, err := selectNonDraftMerchantsCount(db)
	if err != nil {
		return consumer.MerchantsListProps{}, err
	}

	page, offset, limit := utils.ParseListingPageParams(c, total)

	data, err := selectAllNonDraftMerchants(db, offset, limit)
	if err != nil {
		return consumer.MerchantsListProps{}, err
	}

	pagination := table.PaginationFrom(uint(page), total)
	pageItems := table.BuildWindowedPagination(pagination)

	return consumer.MerchantsListProps{
		CurrentPath: url.Customer_ListPage,
		Data:        data,
		Pagination:  pagination,
		PageItems:   pageItems,
	}, nil
}

func (h *Handler) GetAgentsList(c echo.Context) error {
	props, err := buildAgentsListProps(c, h.db)
	if err != nil {
		return render.Render(c, http.StatusOK, consumer.AgentsListContainer(props))
	}

	return render.Render(c, http.StatusOK, consumer.AgentsListContainer(props))
}

func buildAgentsListProps(c echo.Context, db *postgres.Connection) (consumer.AgentsListProps, error) {

	total, err := selectNonDraftMerchantsCount(db)
	if err != nil {
		return consumer.AgentsListProps{}, err
	}

	page, offset, limit := utils.ParseListingPageParams(c, total)

	data, err := selectAllNonDraftAgents(db, offset, limit)
	if err != nil {
		return consumer.AgentsListProps{}, err
	}

	pagination := table.PaginationFrom(uint(page), total)
	pageItems := table.BuildWindowedPagination(pagination)

	return consumer.AgentsListProps{
		CurrentPath: url.Customer_ListPage,
		Data:        data,
		Pagination:  pagination,
		PageItems:   pageItems,
	}, nil
}
