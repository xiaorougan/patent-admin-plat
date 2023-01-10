package dto

type NoveltyReportReq struct {
	KeyWords  []string `json:"keyWords"`
	Title     string   `json:"title"`
	CL        string   `json:"CL"`
	Applicant string   `json:"applicant"`
	Org       string   `json:"org"`
}
