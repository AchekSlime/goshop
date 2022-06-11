package entities

type Product struct {
	Id         int    `json:"id" db:"id"`
	Title      string `json:"title" db:"title"`
	Price      int    `json:"price" db:"price"`
	HolderName string `json:"holder_name" db:"holder_name"`
	CategoryId int    `json:"category_id" db:"category_id"`
}

type Category struct {
	Id   int    `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
}
