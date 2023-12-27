package plugins

import (
	"github.com/atsuiest/gapigate/model"
	"github.com/atsuiest/gapigate/tools"
	"github.com/gofiber/fiber/v2"
)

type JwtPlugin struct {
	Secret     string
	Validation *model.Validation
}

func NewJwtPlugin(secret *model.Secret, validation *model.Validation) Plugin {
	plugin := JwtPlugin{
		Secret:     "",
		Validation: validation,
	}
	return &plugin
}

func (jwtp *JwtPlugin) Process(ctx *fiber.Ctx) error {
	return nil
}

// Verify Signature
func (jwtp *JwtPlugin) VerifySignature() {
	// Verify signature
}

// Verify Claims
func (jwtp *JwtPlugin) VerifyClaims(ctx *fiber.Ctx) {
	// Verify signature
	validate := tools.ValidateClaims(ctx.GetReqHeaders()["Authorization"][0], "system", "public")
	if validate {
		ctx.Next()
	} else {
		res := model.Response{
			Code:    "E403",
			Error:   "Access Denied",
			Message: "The provided access token has no access to requested path",
		}
		ctx.Status(fiber.StatusForbidden).JSON(&res)
	}
}
