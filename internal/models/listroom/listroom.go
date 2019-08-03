package listroom

import (
	"github.com/signaller-matrix/signaller/internal/models/createroom"
)

type Request struct {
	Visibility createroom.VisibilityType `json:"visibility"` // The new visibility setting for the room. Defaults to 'public'. One of: ["private", "public"]
}

type Response struct {
	Visibility createroom.VisibilityType `json:"visibility"` //  	The visibility of the room in the directory. One of: ["private", "public"]
}
