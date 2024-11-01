package chat

type Chat struct {
	ID      int64 `json:id`
	Creator int64 `json:creator`
	Guest   int64 `json:guest`
}
