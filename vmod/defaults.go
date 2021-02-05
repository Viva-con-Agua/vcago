package vmod

type (
    //CodeRequest used for handling LinkToken code.
    CodeRequest struct {
        Code string `json:"code"`
        Case string `json:"case,omitempty"`
    }
)
