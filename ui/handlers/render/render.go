package render

import (
	"net/http"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

func Render(ctx echo.Context, statusCode int, t templ.Component) error {
	buf := templ.GetBuffer()
	defer templ.ReleaseBuffer(buf)

	if err := t.Render(ctx.Request().Context(), buf); err != nil {
		return err
	}

	return ctx.HTML(statusCode, buf.String())
}

func RenderSwap(ctx echo.Context, t templ.Component) error {
	return Render(ctx, http.StatusOK, t)
}

func RedirectPage(ctx echo.Context, url string) error {
	ctx.Response().Header().Set("HX-Redirect", url)
	return ctx.NoContent(http.StatusOK)
}
