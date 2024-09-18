package request

type Headers struct {
	Authorization string `header:"Authorization" binding:"required,min=20"`
}
