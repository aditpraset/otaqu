package controller

import (
	"fmt"
	"log"
	"net/http"
	"otaqu/database"
	"otaqu/model"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/gofiber/fiber/v2"
	uuid "github.com/satori/go.uuid"
	"github.com/shopspring/decimal"
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
	x := 0

	doc.Find("div.__deal_item").Children().Each(func(i int, sel *goquery.Selection) {
		name := sel.Find(".__deal_item_title").Text()
		if name != "" {
			replacer := strings.NewReplacer("\n      ", "", "\n    ", "")
			kupon.UUID = uuid.NewV4()
			kupon.Name = replacer.Replace(sel.Find(".__deal_item_title").Text())
			kupon.Description = replacer.Replace(sel.Find(".__deal_item_description").Text())

			doc.Find("div.__deal_item_buy").Children().Each(func(j int, d *goquery.Selection) {
				full := d.Find("i.fa-star").Size()
				half := d.Find("i.fa-star-half").Size()
				valFull := float32(full)
				valHalf := float32(half) * 0.5
				rating := valFull + valHalf
				kupon.Rating = decimal.NewFromFloat32(float32(rating))

			})

			link, _ := sel.Attr("href")
			err := database.DB.Create(&kupon).Error
			if err != nil {
				x--
			}
			if link != "" {
				dtl, err := http.Get(link)
				if err != nil {
					log.Fatal(err)
				}
				defer dtl.Body.Close()

				if dtl.StatusCode != 200 {
					log.Fatalf("status code error: %d %s", dtl.StatusCode, dtl.Status)
				}

				det, err := goquery.NewDocumentFromReader(dtl.Body)
				if err != nil {
					log.Fatal(err)
				}

				det.Find("div.__box.__box_shadow").Find("div.row").Children().Each(func(i int, d *goquery.Selection) {
					KuponPrice.UUID = uuid.NewV4()
					KuponPrice.KuponID = kupon.UUID
					KuponPrice.Name = d.Find("h3").Text()
					replacer = strings.NewReplacer("Rp", "", ".", "")

					price, _ := strconv.Atoi(replacer.Replace(d.Find("li.__deal_detail_package_list").Find("div.h4").Text()))

					if err != nil {
						KuponPrice.Price = 0
					}
					KuponPrice.Price = uint64(price)
					database.DB.Create(&KuponPrice)

				})
			}

			x++
		}

	})
	response.Success = true
	response.Message = fmt.Sprintf("%d%s", x, " Data Berhasil di Scraping")
	return ctx.JSON(response)

}

func ShowKupon(ctx *fiber.Ctx) error {
	var response model.Response
	kupons := []model.Kupon{}

	data := database.DB.Preload("KuponPrice").Find(&kupons)
	response.Success = true
	response.Message = "Proses Berhasil"
	response.Data = data

	return ctx.JSON(response)

}
