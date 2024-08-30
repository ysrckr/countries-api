package handlers

import (
	"strings"

	"github.com/gofiber/fiber/v3"
	"github.com/ysrckr/countries-api/internals/db"
	"github.com/ysrckr/countries-api/internals/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetCountriesHandler(c fiber.Ctx) error {

	queries, ok := getQueries(c)
	if ok {
		return withQueries(c, queries)
	}

	countries := []models.Country{}

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

func withQueries(c fiber.Ctx, queries map[string]string) error {
	countries := []models.QueryCountry{}
	var fields []string
	var languages []string
	optionsList := bson.D{bson.E{Key: "name.common", Value: 1}}

	fieldsQuery := queries["fields"]
	if fieldsQuery != "" {
		fields = strings.Split(fieldsQuery, ",")
	}
	languagesQuery := queries["languages"]
	if languagesQuery != "" {
		languages = strings.Split(languagesQuery, ",")
	}

	for _, field := range fields {
		optionsList = append(optionsList, bson.E{Key: field, Value: 1})
	}

	for _, language := range languages {
		optionsList = append(optionsList, bson.E{Key: "translations." + language + ".common", Value: 1})
	}

	opts := options.Find().SetProjection(optionsList)

	cursor, err := db.DB.QueryAll(c.Context(), "countries", bson.D{}, opts)
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

func getQueries(c fiber.Ctx) (map[string]string, bool) {
	queries := c.Queries()
	length := len(queries)
	if length == 0 {
		return queries, false
	}

	return queries, true
}
