package nucredit

// isTargetOfGradePointAverage GPAの計算の対象かを返す
func (subject Subject) isTargetOfGradePointAverage() bool {
	switch subject.Grade {
	case "S":
		return true
	case "A":
		return true
	case "B":
		return true
	case "C":
		return true
	case "F":
		return true
	default:
		return false
	}
}

// CalcGradePoint GradePointを返す
func (subject Subject) CalcGradePoint() float64 {
	switch subject.Grade {
	case "S":
		return 4.3 * subject.Credit
	case "A":
		return 4.0 * subject.Credit
	case "B":
		return 3.0 * subject.Credit
	case "C":
		return 2.0 * subject.Credit
	case "F":
		return 0.0 * subject.Credit
	default:
		return 0.0
	}
}

// CalcGradePointAverage GPAを返す
func (subjects Subjects) CalcGradePointAverage() float64 {
	creditSum := 0.0
	gradePointSum := 0.0
	for _, subject := range subjects {
		if subject.isTargetOfGradePointAverage() {
			creditSum += subject.Credit
			gradePointSum += subject.CalcGradePoint()
		}
	}
	if creditSum == 0.0 {
		return 0
	}
	return gradePointSum / creditSum
}
