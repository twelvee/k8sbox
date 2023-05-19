package structs

type Box struct {
	Tag          string
	Type         string
	Applications []Application
	Chart        string
	Values       string
}

type BoxService struct {
	ValidateBoxes func([]Box, string, ApplicationService) error
}
