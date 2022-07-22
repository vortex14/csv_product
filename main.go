package main

import (
	"fmt"
	"strings"

	"github.com/brianvoe/gofakeit/v6"

	"github.com/labstack/echo/v4"

	"net/http"
)

func main() {
	e := echo.New()
	e.GET("/product", func(c echo.Context) error {

		pvalue := fmt.Sprintf(`code;code_for_site;title;group;is_service;articul;brand;country_manufacturer;oem_approved;oem;weight;volume;options;ngroup;barcode;tracking_type
%s;%d;"%s";;0;%d;THO;;48076-11222,48076-11222,48076-11222;48076-11222,48076-11222,48076-11222;0.5211222;0.0015360011222;{};;45601169711222;1`,
			strings.ToUpper(gofakeit.LetterN(9)),
			gofakeit.Number(111111111, 999999999),
			gofakeit.BeerName(),
			gofakeit.Number(00000000, 99999999),
		)

		return c.String(http.StatusOK, pvalue)
	})
	e.Logger.Fatal(e.Start(":12493"))
}
