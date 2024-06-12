package model 

type Question struct {
	ID string `json:"id"`
	Question string `json:"question"`
	Answer string `json:"answer"` // can be empty string 
    Company string `json:"company"`

}
