package main

import (
	"encoding/csv"
	"fmt"
	"github.com/gocolly/colly"
	"os"
)

type Bloger struct {
	Rank      string
	Name      string
	Category  string
	Followers string
	Country   string
	Authentic string
	Avg       string
}

func main() {

	scrapURL := "https://hypeauditor.com/top-instagram-all-russia/"

	// Создание нового коллектора
	c := colly.NewCollector(
		colly.AllowedDomains("www.hypeauditor.com", "hypeauditor.com"),
	)

	c.OnRequest(func(r *colly.Request) {
		fmt.Printf("Visiting %s\n", r.URL)
	})

	c.OnError(func(r *colly.Response, e error) {
		fmt.Printf("Error while scraping: %s\n", e.Error())
	})
	// Срез для хранения данных
	var blogers []Bloger

	// Обработка каждой строки таблицы на странице
	c.OnHTML("div.row div.row__top", func(e *colly.HTMLElement) {
		fmt.Println('a')
		rank := e.ChildText(".row-cell.rank span")
		name := e.ChildText(".contributor__title")
		category := e.ChildText(".row-cell.category .topic .ellipsis")
		followers := e.ChildText(".row-cell.subscribers")
		country := e.ChildText(".row-cell.audience")
		authentic := e.ChildText(".row-cell.authentic")
		avg := e.ChildText(".row-cell.engagement")

		bloger := Bloger{
			Rank:      rank,
			Name:      name,
			Category:  category,
			Followers: followers,
			Country:   country,
			Authentic: authentic,
			Avg:       avg,
		}

		blogers = append(blogers, bloger)
		fmt.Println(blogers)
	})

	// Посещение страницы и скрапинг данны
	c.Visit(scrapURL)

	// Создание CSV файла
	file, err := os.Create("instagram_top_russia.csv")
	if err != nil {
		fmt.Println("Ошибка при создании файла:", err)
		return
	}
	defer file.Close()

	// Создание записи в CSV файле
	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Запись данных в CSV файл
	headers := []string{
		"rank",
		"name",
		"category",
		"followers",
		"country",
		"authentic",
		"avg",
	}

	if err := writer.Write(headers); err != nil {
		return
	}

	for _, bloger := range blogers {
		record := []string{
			bloger.Rank,
			bloger.Name,
			bloger.Category,
			bloger.Followers,
			bloger.Country,
			bloger.Authentic,
			bloger.Avg,
		}

		if err := writer.Write(record); err != nil {
			return
		}
	}

	fmt.Println("Данные успешно записаны в instagram_top_russia.csv")
}
