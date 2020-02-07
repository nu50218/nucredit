package nucredit

import "strings"

// IsTarget GPAの計算の対象かを返す
func (subject *Subject) IsTarget() bool {
	return strings.Contains("SABCF", subject.Grade)
}

// GP GradePointを返す
func (subject *Subject) GP() float64 {
	switch subject.Grade {
	case "S":
		return 4.3 * subject.Credit
	case "A":
		return 4.0 * subject.Credit
	case "B":
		return 3.0 * subject.Credit
	case "C":
		return 2.0 * subject.Credit
	default:
		return 0.0
	}
}

// GPA GPAを返す
func (subjects *Subjects) GPA() float64 {
	creditSum := 0.0
	gradePointSum := 0.0
	for _, subject := range *subjects {
		if subject.IsTarget() {
			creditSum += subject.Credit
			gradePointSum += subject.GP()
		}
	}
	if creditSum == 0.0 {
		return 0
	}
	return gradePointSum / creditSum
}
