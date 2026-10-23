package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"time"
)

type DayRecord struct {
	Mood        int    `json:"mood"`
	JournalText string `json:"journal_text"`
	Timestamp   string `json:"timestamp"` 
	DateKey     string `json:"date_key"`  
}

type JournalDatabase struct {
	CurrentDay int                `json:"current_day"`
	History    map[int]DayRecord  `json:"history"`
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Error: Run with an entry file. E.g., inkwell entry.ink")
		os.Exit(1)
	}

	filename := os.Args[1]
	code, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		os.Exit(1)
	}

	db := loadDatabase()
	todayRecord := DayRecord{Mood: 5, JournalText: ""}

	
	nowKey := time.Now().Format("2006-01-02")
	for d, rec := range db.History {
		if rec.DateKey == nowKey {
			db.CurrentDay = d
			todayRecord = rec
			break
		}
	}

	executeInkwell(string(code), db, &todayRecord)
}

func executeInkwell(code string, db *JournalDatabase, today *DayRecord) {
	lines := strings.Split(code, "\n")
	reader := bufio.NewReader(os.Stdin)

	for i := 0; i < len(lines); i++ {
		trimmed := strings.TrimSpace(lines[i])

		if trimmed == "" || strings.HasPrefix(trimmed, "~") {
			continue
		}

		
		if strings.HasPrefix(trimmed, "note mood = prompt ") {
			promptMsg := strings.Trim(strings.TrimSpace(trimmed[19:]), "\"")
			fmt.Println(promptMsg)
			fmt.Print("> ")
			
			input, _ := reader.ReadString('\n')
			input = strings.TrimSpace(input)
			
			if num, err := strconv.Atoi(input); err == nil {
				today.Mood = num
			} else {
				today.Mood = 5 
			}
			continue
		}

		
		if strings.HasPrefix(trimmed, "entry ") {
			promptMsg := strings.Trim(strings.TrimSpace(trimmed[6:]), "\"")
			fmt.Println(promptMsg)
			fmt.Print("> ")
			
			paragraph, _ := reader.ReadString('\n')
			today.JournalText = strings.TrimSpace(paragraph)

			
			now := time.Now()
			today.Timestamp = now.Format("03:04:05 PM")
			today.DateKey = now.Format("2006-01-02")

			isNewCalendarDay := true
			targetDayIndex := db.CurrentDay

			for d, record := range db.History {
				if record.DateKey == today.DateKey {
					isNewCalendarDay = false
					targetDayIndex = d
					break
				}
			}

			if isNewCalendarDay {
				db.CurrentDay++
				targetDayIndex = db.CurrentDay
			}

			db.History[targetDayIndex] = *today
			saveDatabase(db)
			continue
		}

		if trimmed == "graph mood 365" {
			renderRealTimelineDashboard(db)
			continue
		}
	}
}

func loadDatabase() *JournalDatabase {
	dbFile := "journal_db.json"
	if _, err := os.Stat(dbFile); os.IsNotExist(err) {
		return &JournalDatabase{CurrentDay: 0, History: make(map[int]DayRecord)}
	}
	data, _ := ioutil.ReadFile(dbFile)
	var db JournalDatabase
	json.Unmarshal(data, &db)
	return &db
}

func saveDatabase(db *JournalDatabase) {
	data, _ := json.MarshalIndent(db, "", "  ")
	ioutil.WriteFile("journal_db.json", data, 0644)
}

func renderRealTimelineDashboard(db *JournalDatabase) {
	fmt.Println("\n========================= 📑 INKWELL JOURNAL TIMELINE STREAM =========================")
	fmt.Printf("Current loop cycle progress: (Day %d/365):\n\n", db.CurrentDay)

	for d := 1; d <= 365; d++ {
		record, exists := db.History[d]
		if !exists {
			continue
		}

		textLog := record.JournalText
		if len(textLog) > 25 {
			textLog = textLog[:22] + "..."
		}

		var sparkSymbol string
		switch {
		case record.Mood <= 3:  sparkSymbol = "■ [LOW]"
		case record.Mood <= 6:  sparkSymbol = "■■■ [MED]"
		case record.Mood <= 8:  sparkSymbol = "■■■■■ [HIGH]"
		default:            sparkSymbol = "■■■■■■■ [PEAK]"
		}

		paddedText := fmt.Sprintf("%-25s", textLog)
		fmt.Printf("Day %03d | %s | Entered at %s | Mood: %s\n", d, paddedText, record.Timestamp, sparkSymbol)
	}
	fmt.Println("======================================================================================\n")
}

