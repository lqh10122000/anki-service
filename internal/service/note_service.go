package service

import (
	"bytes"
	"complaint-service/internal/model"
	"complaint-service/internal/repository"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type NoteService interface {
	AddNotes(note *model.Notes) error
}

type noteService struct {
	repo     repository.NoteRepository
	drepo    repository.DeskRepository
	dService DeskService
}

type Audio struct {
	URL      string   `json:"url"`
	Filename string   `json:"filename"`
	SkipHash string   `json:"skipHash"`
	Fields   []string `json:"fields"`
}

type Note struct {
	DeckName  string            `json:"deckName"`
	ModelName string            `json:"modelName"`
	Fields    map[string]string `json:"fields"`
	Options   map[string]bool   `json:"options"`
	Tags      []string          `json:"tags"`
	Audio     []Audio           `json:"audio"`
}

type AddNoteRequest struct {
	Action  string `json:"action"`
	Version int    `json:"version"`
	Params  struct {
		Note Note `json:"note"`
	} `json:"params"`
}

func NewNoteService(repo repository.NoteRepository, drepo repository.DeskRepository, deskService DeskService) NoteService {
	return &noteService{
		repo:     repo,
		drepo:    drepo,
		dService: deskService,
	}
}

func (s *noteService) AddNotes(note *model.Notes) error {

	audioURL := fmt.Sprintf(
		"https://translate.google.com/translate_tts?ie=UTF-8&q=%s&tl=%s&client=tw-ob",
		url.QueryEscape(note.Word),
		note.Lang,
	)

	randomString := GetPersistentRandomString()

	var desk, err = s.drepo.FindByID(note.DeskID)

	var translatedWord string

	if err != nil {
		s.dService.InputDesk()
		desk, err = s.drepo.FindByID(note.DeskID)
		fmt.Errorf("find error: ", err)
		fmt.Errorf("find iddd: ", desk)
	}

	if err != nil {
		return fmt.Errorf("desk with ID %d not found", note.DeskID)
	}

	if strings.TrimSpace(note.TranslateWord) == "" {
		translatedWord, err = TranslateToVietnamese(note.Word, note.Lang)
		fmt.Println("translatedWord in lang with word: ", note.Word, " lang: ", note.Lang)
		fmt.Println("translatedWord in lang:", translatedWord)
	} else {
		translatedWord = note.TranslateWord
		fmt.Println("translatedWord not in lang:", translatedWord)
	}

	fmt.Println("note.TranslateWord:", note.TranslateWord)

	notes := Note{
		DeckName:  desk.Name,
		ModelName: note.ModelName, //"Basic"
		Fields: map[string]string{
			"word":           note.Word,
			"ipa":            "", // Có thể lấy qua API khác nếu cần
			"translate_word": translatedWord,
			"create_date":    time.Now().Format("2006-01-02"),
			"audio_filename": randomString + ".mp3",
		},
		Options: map[string]bool{
			"allowDuplicate": false,
		},
		Tags: []string{"ielts", "vocab"},
		Audio: []Audio{
			{
				URL:      audioURL,
				Filename: randomString + ".mp3",
				SkipHash: "7e2c2f954ef6051373ba916f000168dc",
				Fields:   []string{"audio_word"},
			},
		},
	}

	req := AddNoteRequest{
		Action:  "addNote",
		Version: 6,
	}
	req.Params.Note = notes

	jsonData, err := json.Marshal(req)
	if err != nil {
		return err
	}

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Post("http://localhost:8765", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	var result map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return err
	}

	if result["error"] != nil {
		return fmt.Errorf("Anki error: %v", result["error"])
	}

	fmt.Println("✅ Added note with audio successfully!")
	s.repo.AddNotes(note)
	return nil
}

func GetPersistentRandomString() string {
	rand.Seed(time.Now().UnixNano())
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	randomRunes := make([]rune, 10)
	for i := range randomRunes {
		randomRunes[i] = letters[rand.Intn(len(letters))]
	}
	return string(randomRunes)
}

func TranslateToVietnamese(text, fromLang string) (string, error) {
	// Tạo URL cho Google Translate API
	apiURL := "https://translate.googleapis.com/translate_a/single"
	params := url.Values{}
	params.Add("client", "gtx")
	params.Add("sl", fromLang) // source language
	params.Add("tl", "vi")     // target language
	params.Add("dt", "t")
	params.Add("q", text)

	finalURL := fmt.Sprintf("%s?%s", apiURL, params.Encode())

	// Gửi request
	resp, err := http.Get(finalURL)
	if err != nil {
		return "", fmt.Errorf("request error: %w", err)
	}
	defer resp.Body.Close()

	// Đọc response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("read error: %w", err)
	}

	// Debug raw response nếu cần
	// fmt.Println(string(body))

	// Parse JSON kết quả
	var response []interface{}
	if err := json.Unmarshal(body, &response); err != nil {
		return "", fmt.Errorf("unmarshal error: %w", err)
	}

	// Trích ra phần dịch
	sentences, ok := response[0].([]interface{})
	if !ok || len(sentences) == 0 {
		return "", fmt.Errorf("unexpected format in response")
	}

	firstSentence, ok := sentences[0].([]interface{})
	if !ok || len(firstSentence) < 1 {
		return "", fmt.Errorf("unexpected sentence format")
	}

	translated, ok := firstSentence[0].(string)
	if !ok {
		return "", fmt.Errorf("translation not found")
	}

	return translated, nil
}
