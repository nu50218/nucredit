package nucredit

import (
	"encoding/json"
	"io"
)

// Subject 科目
type Subject struct {
	// Category 科目区分
	Category string `json:"category"`
	// Name 科目名
	Name string `json:"name"`
	// Teacher 担当教員名
	Teacher string `json:"teacher"`
	// Credit 単位数
	Credit float64 `json:"credit"`
	// Grade 評価
	Grade string `json:"grade"`
	// Year 修得年度
	Year int `json:"year"`
	// Semester 修得学期
	Semester string `json:"semester"`
}

// Subjects 科目列
type Subjects []*Subject

// FromReader Readerから読んでSubjectsにして返す
func FromReader(r io.Reader) (*Subjects, error) {
	s := &Subjects{}
	if err := json.NewDecoder(r).Decode(s); err != nil {
		return nil, err
	}
	return s, nil
}

// FilterFunc フィルタリングの条件
type FilterFunc func(*Subject) bool

// Filter 与えた条件に一致するSubjectsを返す
func (subjects *Subjects) Filter(filterFn FilterFunc) *Subjects {
	res := Subjects{}
	for _, subject := range *subjects {
		if filterFn(subject) {
			res = append(res, subject)
		}
	}
	return &res
}
