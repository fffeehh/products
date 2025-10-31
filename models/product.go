package models
:wa
:wqa

type Product struct{
	ID int `json:"id"`
	Title string `json:"title"`
	NumberOfGrams float32 `json:"numberOfGram"`
}
