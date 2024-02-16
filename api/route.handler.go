package api

import (
	// "strings"

	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/atsuiest/gapigate/config"
	"github.com/atsuiest/gapigate/model"
	"github.com/atsuiest/gapigate/plugins"
	"github.com/gofiber/fiber/v2"
)

func RouteHandler(ctx *fiber.Ctx) error {
	backendKey := fmt.Sprintf("%s|%s", ctx.Method(), ctx.OriginalURL())
	data := findMap(backendKey)
	if data == nil {
		data = findMap(backendKey[:strings.LastIndex(backendKey, "/")] + "/::")
		if data == nil {
			return ctx.Status(fiber.StatusNotFound).JSON(RES404)
		}
	}
	backend := data.(model.Backend)
	if backend.Plugin.JwtEnabled {
		if !validateAccess(ctx, backend.Plugin.JwtName) {
			return ctx.Status(fiber.StatusForbidden).JSON(RES403)
		}
	}
	status, body := httpCall(ctx.Method(), parseParams(ctx, backend), ctx.Request().Body())
	return ctx.Status(status).Send(body)
}

func findMap(key string) interface{} {
	value, ok := config.BackendMap[key]
	if ok {
		return value
	}
	return nil
}

func validateAccess(ctx *fiber.Ctx, jwtName string) bool {
	validation, ok := config.ValidationsMap["jwt|"+jwtName]
	if !ok {
		return false
	}
	return plugins.CheckClaims(ctx, validation)
}

func parseParams(ctx *fiber.Ctx, backend model.Backend) string {
	target := backend.Target.URL
	if strings.Contains(backend.Pattern, "/:") {
		urlData := strings.Split(backend.Pattern, "/:")
		value := strings.Split(ctx.OriginalURL(), urlData[0])[1]
		handler := urlData[1]
		target = strings.Replace(target, "/:"+handler, value, 1)
	}
	return target
}

func httpCall(method string, url string, body []byte) (int, []byte) {
	req, err := http.NewRequest(method, url, bytes.NewBuffer(body))
	if err != nil {
		fmt.Printf("client: could not create request: %s\n", err)
		os.Exit(1)
	}
	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return 500, nil
	}

	defer res.Body.Close()

	fmt.Printf("client: got response!\n")
	fmt.Printf("client: status code: %d\n", res.StatusCode)

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		return 500, nil
	}
	return res.StatusCode, resBody
}
