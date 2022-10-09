package questions

import (
	"fmt"
	"message-tagger/settings"
	"strings"
)


type Question struct {
	Tags             []string `json:"tags"`
	IsAnswered       bool     `json:"is_answered"`
	LastActivityDate int      `json:"last_activity_date"`
	CreationDate     int      `json:"creation_date"`
	QuestionID       int      `json:"question_id"`
	Link             string   `json:"link"`
	Title            string   `json:"title"`
	Body             string   `json:"body"`
	Reasons []string `json:"reasons"`
}

func (q *Question) HasTag(wantedTag string) bool {
	for _, tag := range q.Tags {
		if tag == wantedTag {
			return true
		}
	}
	return false
}

func (q *Question) ContainsWords(words []string) bool {
	for _, ext := range words {
		if strings.Contains(q.Body, ext) {
			return true
		}
	}
	return false
}

func (q *Question) Analize(s *settings.Settings) {
	var reasons []string
	if q.HasTag(s.BareTag) == false {
		reasons = append(reasons, fmt.Sprintf("Has no %v Tag", s.BareTag))
	}
	if q.HasTag(s.DesirableTag) == true {
		reasons = append(reasons, fmt.Sprintf("Has %v Tag", s.DesirableTag))
	}

	lowerBody := strings.ToLower(q.Body)
	for _, word := range s.WarningStrings {
		if strings.Contains(lowerBody, word) {
			reasons = append(reasons, fmt.Sprintf("Has %v word", word))
		}
	}
	if len(reasons) > 0 {
		q.Reasons = reasons
	}
}