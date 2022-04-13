package controller

import (
	"log"
	"net/http"
	"otaqu/database"
	"otaqu/model"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/gofiber/fiber/v2"
	uuid "github.com/satori/go.uuid"
)

func GetKupon(ctx *fiber.Ctx) error {
	kupon := model.Kupon{}
	KuponPrice := model.KuponPrice{}
	var response model.Response

	res, err := http.Get("https://lakupon.com/category/terpopuler?utm_source=homepage&utm_medium=category_header&utm_campaign=terpopuler")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	rows := make([]model.GetKupon, 0)

	doc.Find(".__deal_item").Children().Each(func(i int, sel *goquery.Selection) {
		row := new(model.GetKupon)
		row.Name = sel.Find(".__deal_item_title").Text()
		row.Description = sel.Find(".__deal_item_description").Text()
		doc.Find(".__deal_item_rating").Children().Each(func(i int, sel *goquery.Selection) {
			row.Rating = sel.Find(".__deal_item_rating_rate").Text()
		})
		row.Price = sel.Find(".__deal_item_price_sell").Text()

		rows = append(rows, *row)
	})

	for i := 0; i < len(rows); i++ {
		if rows[i].Name != "" {
			kupon.UUID = uuid.NewV4()
			replacer := strings.NewReplacer("\n      ", "", "\n    ", "")

			kupon.Name = replacer.Replace(rows[i].Name)
			kupon.Description = replacer.Replace(rows[i].Description)
			value, err := strconv.ParseFloat(strings.Replace(rows[i].Rating, "/10", "", -1), 32)
			if err != nil {
				kupon.Rating = 0
			}
			kupon.Rating = float32(value)

			KuponPrice.UUID = uuid.NewV4()
			KuponPrice.KuponID = kupon.UUID
			replacer = strings.NewReplacer("Rp", "", ".", "")
			price, _ := strconv.Atoi(replacer.Replace(rows[i].Price))

			if err != nil {
				KuponPrice.Price = 0
			}
			KuponPrice.Price = uint64(price)

			database.DB.Create(&kupon)
			database.DB.Create(&KuponPrice)
		}
	}
	response.Success = true
	response.Message = "Proses Scrapping Berhasil"

	return ctx.JSON(response)

}

func ShowKupon(ctx *fiber.Ctx) error {
	var response model.Response
	kupons := []model.Kupon{}
	kupon := model.Kupon{}
	kuponPrice := model.KuponPrice{}

	rows, _ := database.DB.Table("kupons as k").
		Joins("Join kupon_prices as kp on k.uuid = kp.kupon_id").
		Select("k.*,kp.uuid,kp.kupon_id,kp.price,kp.created_at,kp.updated_at").Rows()

	for rows.Next() {
		err := rows.Scan(&kupon.UUID, &kupon.Name, &kupon.Description, &kupon.Rating, &kupon.CreatedAt, &kupon.UpdatedAt, &kuponPrice.UUID, &kuponPrice.KuponID, &kuponPrice.Price, &kuponPrice.CreatedAt, &kuponPrice.UpdatedAt)
		if err != nil {
			log.Panic(err)
		}
		kupon.KuponPrice = kuponPrice
		kupons = append(kupons, kupon)
	}
	response.Success = true
	response.Message = "Proses Berhasil"
	response.Data = kupons

	return ctx.JSON(response)

}
