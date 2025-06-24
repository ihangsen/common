package dto

type SearchDto struct {
	Q string `json:"q" binding:"max=10"`
	SearchPageDto
}
