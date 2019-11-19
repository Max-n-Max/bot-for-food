package resources

type CandlesHistoryBody struct {
	Pair        string `json:"pair"`
	Resolution  string `json:"resolution"`
	Start       int64  `json:"start"`
	End         int64  `json:"end"`
	Limit       int    `json:"limit"`
	OldestFirst bool   `json:"OldestFirst"`
}
