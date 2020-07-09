package vo

type PagingRequest struct {
	PageNum  int `binding:"required"`
	PageSize int `binding:"required"`
}
