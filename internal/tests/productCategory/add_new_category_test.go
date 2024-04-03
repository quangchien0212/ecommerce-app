package productCategory

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-faker/faker/v4"
	"github.com/labstack/echo/v4"
	"github.com/quangchien0212/ecommerce-app/internal/server"
	"github.com/quangchien0212/ecommerce-app/internal/tests"
	"github.com/stretchr/testify/assert"
)

func TestAddNewCategory(t *testing.T) {
	testDB, _ := tests.Setup()
	defer tests.Teardown(testDB)

	type FCategory struct {
		Name        string `json:"name" faker:"word,unique"`
		Description string `json:"description" faker:"word,unique"`
	}

	t.Run("should return 201 created for a new category", func(t *testing.T) {
		var customCategory FCategory
		faker.FakeData(&customCategory)
		out, _ := json.Marshal(customCategory)

		e := echo.New()

		req := httptest.NewRequest(
			http.MethodPost,
			"/category",
			strings.NewReader(string(out)),
		)
		rec := httptest.NewRecorder()

		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		c := e.NewContext(req, rec)
		sv := server.NewServer(testDB)

		if assert.NoError(t, sv.AddCategory(c)) {
			assert.Equal(t, http.StatusCreated, rec.Code)
		}
	})
}
