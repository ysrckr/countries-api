package handlers

import (
	"net/http"
	"strings"

	"github.com/gofiber/fiber/v3"
	"github.com/ysrckr/countries-api/internals/db"
	"github.com/ysrckr/countries-api/internals/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetCountriesHandler(c fiber.Ctx) error {

	countries := []models.Country{}
	queries, ok := getQueries(c)
	if ok {
		return withQueries(c, queries, &countries)
	}

	cursor, err := db.DB.QueryAll(c.Context(), "countries", bson.D{}, nil)
	if err != nil {
		return c.JSON(fiber.Map{
			"message": err,
		})
	}

	defer cursor.Close(c.Context())

	if err = cursor.All(c.Context(), &countries); err != nil {
		return c.JSON(fiber.Map{
			"message": err,
		})
	}

	return c.JSON(countries)
}

func withQueries(c fiber.Ctx, queries map[string]string, countries *[]models.Country) error {
	var fields []string
	var translations []string
	optionsList := bson.D{bson.E{Key: "name.common", Value: 1}, bson.E{Key: "cca2", Value: 1}}
	filters := bson.D{}

	fieldsQuery := queries["fields"]
	if fieldsQuery != "" {
		fields = strings.Split(fieldsQuery, ",")
	}

	translationsQuery := queries["translations"]
	if translationsQuery != "" {
		translations = strings.Split(translationsQuery, ",")
	}

	for _, field := range fields {
		optionsList = append(optionsList, bson.E{Key: field, Value: 1})
	}

	for _, translation := range translations {
		optionsList = append(optionsList, bson.E{Key: "translations." + translation, Value: 1})
	}

	opts := options.Find().SetProjection(optionsList)

	cursor, err := db.DB.QueryAll(c.Context(), "countries", filters, opts)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": err,
		})
	}

	defer cursor.Close(c.Context())

	if err = cursor.All(c.Context(), countries); err != nil {
		return c.JSON(fiber.Map{
			"message": err,
		})
	}

	return c.JSON(*countries)
}

func getQueries(c fiber.Ctx) (map[string]string, bool) {
	queries := c.Queries()
	length := len(queries)
	if length == 0 {
		return queries, false
	}

	return queries, true
}
