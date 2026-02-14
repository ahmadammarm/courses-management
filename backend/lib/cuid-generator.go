package lib

import "github.com/nrednav/cuid2"

func GenerateCUID() string {
	return cuid2.Generate()
}

func NewInstructorID() string {
	return "ins_" + GenerateCUID()
}

func NewCourseID() string {
    return "crs_" + GenerateCUID()
}
