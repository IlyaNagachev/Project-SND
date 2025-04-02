package products

import (
	"encoding/json"
	"fmt"
	"github.com/scylladb/gocqlx/v3/qb"
	"snd/backend/store/models"
)

func(s *Storage) Get(uuid string) (product models.Product, err error){
	err = s.session.Query(qb.Select(models.ProductTable).Where(qb.Eq("uuid")).ToCql()).BindMap(qb.M{"uuid":uuid}).GetRelease(&product)
	return
}

func(s *Storage) GetAll() (product []models.Product, err error){
	var m []map[string]interface{}
	if m, err = s.session.Session.Query(fmt.Sprintf("SELECT * FROM %s", models.ProductTable)).Iter().SliceMap(); err != nil{
		return
	}
	var bytes []byte

	if bytes, err = json.Marshal(m); err != nil{
		return
	}

	if err = json.Unmarshal(bytes, &product); err != nil{
		return
	}
	return
}
