package handlers

import (
	"log"
	"strings"

	"github.com/gofiber/fiber/v3"
	"github.com/ysrckr/countries-api/internals/db"
	"github.com/ysrckr/countries-api/internals/models"
)

func GetCountriesHandler(c fiber.Ctx) error {
	countries := []models.Country{}
	queries, ok := getQueries(c)
	if ok {
		filtersQuery := queries["filters"]
		filters := strings.Split(filtersQuery, ",")
		log.Println(filters)

	}

	cursor, err := db.DB.QueryAll(c.Context(), "countries")
	if err = cursor.All(c.Context(), &countries); err != nil {
		return c.JSON(fiber.Map{
			"message": err,
		})
	}
	defer cursor.Close(c.Context())

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
