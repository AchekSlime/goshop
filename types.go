package goshop

type Product struct {
	Id         int    `json:"id" db:"id"`
	Title      string `json:"title" db:"title"`
	Price      int    `json:"price" db:"price"`
	HolderName string `json:"holder_name" db:"holder_name"`
}
