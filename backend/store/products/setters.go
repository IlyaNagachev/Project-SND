package products

import "snd/backend/store/models"

func(s *Storage) New(product *models.Product) (err error){
	return s.session.Query(s.Tables.ProductTable.Insert()).BindStruct(product).Exec()

}
