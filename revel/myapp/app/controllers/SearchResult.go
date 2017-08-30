package controllers

import (
	dbConnector "myapp/app/dbConnector"
	models "myapp/app/models"

	"github.com/revel/revel"
)

func (c App) SearchResult(name string) revel.Result {

	//入力値をセット
	inputData := new(models.SearchInput)
	inputData.SetInput(name)

	//社員情報を検索
	results, err := dbConnector.Search(name)
	if err != nil {
		c.RenderError(err)
	}

	return c.Render(inputData, results)
}
