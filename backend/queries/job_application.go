package queries

// GetJobApplication presents the members of the GetJobApplication URL query.
type GetJobApplication struct {
	Labels []string `form:"labels" json:"labels"`
}

// PostJobApplication presents the members of POST application request.
type PostJobApplication struct {
	JobId     string   `form:"job" json:"job"`
	User      string   `form:"user" json:"user"`
	FirstName string   `form:"firstName" json:"firstName"`
	LastName  string   `form:"lastName" json:"lastName"`
	Email     string   `form:"email" json:"email"`
	Phone     string   `form:"phone" json:"phone"`
	Labels    []string `form:"labels" json:"labels"`
}