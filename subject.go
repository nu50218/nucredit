package nucredit

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"golang.org/x/text/width"
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

// FromReader Readerからjson形式で読んでSubjectsにして返す
func FromReader(r io.Reader) (Subjects, error) {
	s := Subjects{}
	if err := json.NewDecoder(r).Decode(&s); err != nil {
		return nil, err
	}
	return s, nil
}

// FromFile Fileから読んでSubjectsにして返す
func FromFile(filename string) (Subjects, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	return FromReader(f)
}

// FromResponse 成績ページをパースしてSubjectsにして返す
func FromResponse(res *http.Response) (Subjects, error) {
	doc, err := goquery.NewDocumentFromResponse(res)
	if err != nil {
		return nil, err
	}

	subjects := Subjects{}

	// パース
	log.Println(doc.Text())
	table := doc.Find("table").Eq(1)
	log.Println(table.Text())
	body := table.Find("tbody").Eq(0)
	log.Println(body.Text())
	data := body.Find("tr")
	log.Println(data.Text())

	// 【科目区分 】 -> 科目区分
	categoryFmt := func(s string) string {
		s = strings.TrimLeft(s, "【")
		s = strings.TrimRight(s, "】")
		return strings.TrimSpace(s)
	}

	// dataを処理していく
	// TODO: refactor

	var currentCategory string

	data.EachWithBreak(func(i int, s *goquery.Selection) bool {
		// カラム名はスキップ
		if i == 0 {
			return true
		}

		td := s.Find("td")

		// 科目区分の行 背景の色で判定
		if td.Eq(0).AttrOr("bgColor", "white") != "white" {
			currentCategory = categoryFmt(td.Eq(0).Text())
			return true
		}

		// 通常の行
		subject := &Subject{
			Category: currentCategory,
			Name:     td.Eq(0).Text(),
			Teacher:  td.Eq(1).Text(),
			// 半角にする
			Grade: width.Narrow.String(td.Eq(3).Text()),
		}

		if subject.Credit, err = strconv.ParseFloat(td.Eq(2).Text(), 64); err != nil {
			return false
		}

		term := strings.Split(td.Eq(4).Text(), " ")
		// 単位が出ていないと存在しないのでうまくパースできたときだけ追加する
		if len(term) == 2 {
			if subject.Year, err = strconv.Atoi(term[0]); err != nil {
				return false
			}
			subject.Semester = term[1]
		}

		subjects = append(subjects, subject)
		return true
	})

	if err != nil {
		return nil, err
	}
	return subjects, nil
}

// FilterFunc フィルタリングの条件
type FilterFunc func(*Subject) bool

// Filter 与えた条件に一致するSubjectsを返す
func (subjects Subjects) Filter(filterFn FilterFunc) Subjects {
	res := Subjects{}
	for _, subject := range subjects {
		if filterFn(subject) {
			res = append(res, subject)
		}
	}
	return res
}
