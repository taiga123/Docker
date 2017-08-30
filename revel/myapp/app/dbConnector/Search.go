package dbConnector

import (
	"context"
	"encoding/json"
	"io"
	"log"
	models "myapp/app/models"
	"os"

	elastic "gopkg.in/olivere/elastic.v5"
)

func Search(name string) ([]*models.ViewSearchResult, error) {

	//elasticSearch接続
	client, err := elastic.NewClient(elastic.SetURL(os.Getenv("ELASTICSEARCH_URL")), elastic.SetSniff(false))
	if err != nil {
		log.Printf("elasticSearch接続失敗")
		return nil, err
	}

	//クエリ発行
	var svc *elastic.ScrollService
	if len(name) != 0 {
		query := elastic.NewQueryStringQuery(`"` + name + `"`)
		query = query.DefaultField("Name")
		svc = client.Scroll().
			Index("employee_info").
			Type("info").
			Query(query).
			Sort("Number", true).
			Size(5000)
	} else {
		svc = client.Scroll().
			Index("employee_info").
			Type("info").
			Sort("Number", true).
			Size(5000)
	}

	//検索結果取得
	var results []*models.ViewSearchResult
	for {
		res, err := svc.Do(context.TODO())
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Printf("elasticSearch結果取得失敗 ： %s", err.Error())
			return nil, err
		}
		if res == nil {
			log.Printf("expectrd results != nil; got nil")
			return nil, err
		}

		for _, hit := range res.Hits.Hits {
			var s models.DBRowData
			err := json.Unmarshal(*hit.Source, &s)
			if err != nil {
				log.Printf("Unmarshal失敗")
				return nil, err
			}

			result := new(models.ViewSearchResult)
			result.SetView(s.Number, s.Name, s.Affiliation, s.Position, s.MailAddress, s.PhoneNumber, s.EntryMonth, s.Birthday)
			results = append(results, result)
		}
	}

	return results, nil
}
