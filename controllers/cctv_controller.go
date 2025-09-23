package controllers

import (
	"backend-ml-cctv-golang/dto"
	"backend-ml-cctv-golang/entity"
	"backend-ml-cctv-golang/repository"
	"errors"
	"fmt"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func StoreCCTVData(ctx *fiber.Ctx) error {
	input := new(dto.RequestCCTV)

	if err := ctx.BodyParser(input); err != nil {
		return ctx.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"http_code": fiber.StatusUnprocessableEntity,
			"errors":    err.Error(),
		})
	}

	validate := validator.New()

	if err := validate.Struct(input); err != nil {
		var errorsMap map[string]interface{}
		var validationErrors validator.ValidationErrors
		if errors.As(err, &validationErrors) {
			errorsMap = make(map[string]interface{})
			for _, fieldError := range validationErrors {
				fieldName := fieldError.Field()
				tagName := fieldError.Tag()
				if fieldName != "" {
					switch tagName {
					case "required":
						errorsMap[fieldName] = map[string]string{"error": fmt.Sprintf("%s Mohon Diisi!", fieldName)}
					case "datetime":
						errorsMap[fieldName] = map[string]string{"error": fmt.Sprintf("%s Format Tanggal Tidak Sesuai!", fieldName)}
					}
				}
			}
		}

		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"http_code": fiber.StatusBadRequest,
			"errors":    errorsMap,
		})
	}

	cctvInput := entity.CCTV{
		NamaCCTV: input.NamaCCTV,
		Objek:    *input.Objek,
	}

	cctvOutput, err := repository.SaveCCTVData(cctvInput)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"http_code": fiber.StatusBadRequest,
			"errors":    err.Error(),
		})
	}

	var result dto.ResponseCCTV
	result.NamaCCTV = cctvOutput.NamaCCTV
	result.Objek = cctvOutput.Objek
	result.Waktu = cctvOutput.CreatedAt.Format(time.RFC3339)

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"http_code": fiber.StatusCreated,
		"message":   "Berhasil Memabuat Data CCTV",
		"data":      result,
	})
}

func GetLastCCTVData(ctx *fiber.Ctx) error {
	cctvs, err := repository.GetLatestCCTVData()

	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"http_code": fiber.StatusNotFound,
			"errors":    err.Error(),
		})
	}

	var result []dto.ResponseCCTV
	for _, cctv := range cctvs {
		result = append(result, dto.ResponseCCTV{
			NamaCCTV: cctv.NamaCCTV,
			Objek:    cctv.Objek,
			Waktu:    cctv.CreatedAt.Format(time.RFC3339),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"http_code": fiber.StatusOK,
		"message":   "Ok Zone",
		"data":      result,
	})

}
