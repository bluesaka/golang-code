package myexcelize

import (
	"bytes"
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize/v2"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	jsoniter "github.com/json-iterator/go"
	"github.com/spf13/cast"
	"log"
	"math/rand"
	"strings"
	"time"
)

type question struct {
	Content      string
	Resource     string
	Choices      string
	Answer       int
	QuestionType int `gorm:"question_type"`
	Sequence     int
}

//type answer struct {
//	Answer  int
//	Choices []struct {
//		ID      int
//		Content string
//	}
//}

type answer struct {
	FirstChoiceID int
	Choices       []choice
}

type choice struct {
	ID      int
	Content string
}

/**
insert into question(content,resource,choice,answer,question_type,sequence) VALUES
('诸葛亮和诸葛恪是什么关系？', '', '["aaa","bbb"]', 1, 1,1),
('诸葛亮和诸葛恪是什么关系2？', '', '["aaa2","bbb2"]', 1, 1,1);
*/
func WriteExcelDataToDatabase() {
	//打开文件
	file, err := excelize.OpenFile("../file/sgs.xlsx")
	if err != nil {
		panic(err)
	}

	var qs []question
	answers := initAnswer()

	// 获取sheet中所有的行
	rows, err := file.GetRows("题目")
	if err != nil {
		log.Println("get rows error:", err)
		return
	}

	for index, row := range rows {
		if index == 0 {
			continue
		}

		answerID := cast.ToInt(row[3])
		choiceIDs := strings.Split(row[4], ";")
		choices := getChoices(answers, cast.ToInt(choiceIDs[0]))
		shuffle(choices)
		answerIndex := getAnswer(choices, answerID)
		questionType := getQuestionType(row[1])
		resource := ""
		if len(row) >= 6 {
			resource = row[5]
		}

		// 题目	题目类型	题目					正确答案	所有答案					附件编号			备注
		// 1	历史	诸葛亮和诸葛恪是什么关系？	1001	1001;1002;1003;1004
		qs = append(qs, question{
			Content:      row[2],
			Resource:     getResource(resource, questionType),
			Choices:      wrapChoices(choices),
			Answer:       answerIndex,
			QuestionType: questionType,
			Sequence:     cast.ToInt(row[0]),
		})
	}
	//log.Printf("%+v\n", qs)

	bulkInsert(qs)
}

func bulkInsert(qs []question) {
	db, err := gorm.Open("mysql", "root:123456@(127.0.0.1:3306)/sgs-exam?charset=utf8&parseTime=true")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	var buf bytes.Buffer
	sql := "insert into `question`(content,resource,choices,answer,question_type,sequence) VALUES"
	buf.WriteString(sql)
	l := len(qs)

	for k, v := range qs {
		if k == l-1 {
			buf.WriteString(fmt.Sprintf("('%s', '%s', '%s', %d, %d, %d);", v.Content, v.Resource, v.Choices, v.Answer, v.QuestionType, v.Sequence))
		} else {
			buf.WriteString(fmt.Sprintf("('%s', '%s', '%s', %d, %d, %d),", v.Content, v.Resource, v.Choices, v.Answer, v.QuestionType, v.Sequence))
		}

	}

	log.Println(buf.String())
	if err := db.Exec(buf.String()).Error; err != nil {
		log.Println("exec insert error:", err)
	}
}

func initAnswer() []answer {
	//打开文件
	file, err := excelize.OpenFile("../file/sgs.xlsx")
	if err != nil {
		panic(err)
	}

	// 获取sheet中所有的行
	rows, err := file.GetRows("答案")

	var as []answer
	var a answer
	count := 0

	// 答案编号	答案
	// 1001		叔侄
	for index, row := range rows {
		if index == 0 {
			continue
		}

		//a.Choices = append(a.Choices, struct {
		//	ID      int
		//	Content string
		//}{ID: cast.ToInt(row[0]), Content: row[1]})

		a.Choices = append(a.Choices, choice{
			ID:      cast.ToInt(row[0]),
			Content: row[1],
		})

		count++

		if count == 4 {
			a.FirstChoiceID = a.Choices[0].ID
			as = append(as, a)
			a = answer{}
			count = 0
		}
	}

	return as
}

func getChoices(as []answer, choiceID int) []choice {
	for i := 0; i < len(as); i++ {
		if as[i].FirstChoiceID == choiceID {
			return as[i].Choices
		}
	}
	return nil
}

func getResource(resource string, questionType int) string {
	if resource == "" {
		return ""
	}
	s := strings.Split(resource, ";")
	for i := 0; i < len(s); i++ {
		suffix := "png"
		if questionType == 4 {
			suffix = "mp3"
		}
		s[i] = s[i] + "." + suffix
	}
	json, _ := jsoniter.Marshal(s)
	return string(json)
}

func wrapChoices(choices []choice) string {
	m := make([]string, 0, 4)

	//for i := 0; i < len(choices); i++ {
	//	m = append(m, choices[i].Content)
	//}
	for _, v := range choices {
		m = append(m, v.Content)
	}
	json, _ := jsoniter.Marshal(m)
	return string(json)
}

// shuffle shuffle slice
func shuffle(s []choice) {
	for i := len(s) - 1; i >= 0; i-- {
		rand.Seed(time.Now().UnixNano())
		j := rand.Intn(i + 1)
		s[i], s[j] = s[j], s[i]
	}
	//rand.Shuffle(len(s), func(i, j int) {
	//	s[i], s[j] = s[j], s[i]
	//})
}

func getAnswer(s []choice, id int) int {
	for i := 0; i < len(s); i++ {
		if s[i].ID == id {
			return i + 1
		}
	}
	return 0
}

//func getAnswer(s []struct {
//	ID      int
//	Content string
//}, id int) int {
//	for i := 0; i < len(s); i++ {
//		if s[i].ID == id {
//			return i
//		}
//	}
//	return 0
//}

func getQuestionType(s string) int {
	switch s {
	case "历史":
		return 1
	case "游戏":
		return 2
	case "图形":
		return 3
	case "听力":
		return 4
	}
	return 0
}
