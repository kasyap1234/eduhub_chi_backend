package model 

type Company struct {
	ID string `json:"id"`
	Name string `json:"name"`
	Email string `json:"email, omitempty"`
	Phone string `json:"phone,omitempty"`
	Website string `json:"website,omitempty"`
	ImageURL string `json:"image_url,omitempty"`

}
