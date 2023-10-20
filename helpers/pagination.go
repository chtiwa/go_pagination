package helpers

type PaginationData struct {
	NextPage     int
	PreviousPage int
	CurrentPage  int
	TotalPages   int
	TwoAfter     int
	TwoBelow     int
	ThreeAfter   int
	BaseURL      string
}

func GetPaginationData(page int, totalPages int, baseURL string) PaginationData {
	return PaginationData{
		NextPage:     page + 1,
		PreviousPage: page - 1,
		CurrentPage:  page,
		TotalPages:   totalPages,
		TwoAfter:     page + 2,
		TwoBelow:     page - 2,
		ThreeAfter:   page + 3,
		BaseURL:      baseURL,
	}
}
