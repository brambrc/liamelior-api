package Model

import (
	"fmt"
	"liamelior-api/Database"
	"sort"
	"strings"
	"time"

	"gorm.io/gorm"
)

type Schedule struct {
	gorm.Model
	Setlist  string `json:"setlist" binding:"required"`
	ShowDate string `json:"show_date" binding:"required"`
	Time     string `json:"time" binding:"required"`
}

func (schedule *Schedule) Save() (*Schedule, error) {
	var err error

	err = Database.Database.Create(&schedule).Error
	if err != nil {
		return &Schedule{}, err
	}

	return schedule, nil
}

func CheckScheduleExists(setlist, showDate, showTime string) (bool, error) {
	var count int64
	err := Database.Database.Model(&Schedule{}).Where("setlist = ? AND show_date = ? AND time = ?", setlist, showDate, showTime).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func GetUpcomingShows() ([]map[string]string, error) {
	// Get the current date and time
	currentTime := time.Now()

	// Retrieve shows from the database
	var schedules []Schedule
	err := Database.Database.
		Table("schedules").
		Select("setlist, show_date, time").
		Find(&schedules).Error
	if err != nil {
		return nil, err
	}

	// Filter out shows that have already occurred or the time has passed
	upcomingShows := make([]map[string]string, 0)
	layout := "2.1.2006 15:04"
	for _, schedule := range schedules {
		showTime := strings.TrimPrefix(schedule.Time, "Show ")

		// Format the show_date column with a consistent format using PostgreSQL's TO_CHAR function
		// This converts the single digit day to two digits with leading zero
		formattedShowDate := "TO_CHAR(show_date::date, 'DD.MM.YYYY')"

		// Combine date and time as a single string
		showDateTimeCombined := fmt.Sprintf("%s %s", formattedShowDate, showTime)

		// Trim leading/trailing whitespaces from the combined string
		showDateTimeCombined = strings.TrimSpace(showDateTimeCombined)

		// Parse combined date and time using the layout
		showDateTimeParsed, parseErr := time.Parse(layout, showDateTimeCombined)
		if parseErr != nil {
			fmt.Println("Error parsing show date and time:", parseErr)
			continue
		}

		// Compare date and time
		if showDateTimeParsed.After(currentTime) || (showDateTimeParsed.Equal(currentTime) && showDateTimeParsed.After(currentTime.Add(time.Minute))) {
			upcomingShow := map[string]string{
				"setlist":   schedule.Setlist,
				"show_date": schedule.ShowDate,
				"time":      showTime, // Use the trimmed show time
			}
			upcomingShows = append(upcomingShows, upcomingShow)
		}
	}

	// Sort the upcoming shows by show date and time
	sort.Slice(upcomingShows, func(i, j int) bool {
		showDateTime1, _ := time.Parse(layout, upcomingShows[i]["show_date"]+" "+upcomingShows[i]["time"])
		showDateTime2, _ := time.Parse(layout, upcomingShows[j]["show_date"]+" "+upcomingShows[j]["time"])
		return showDateTime1.Before(showDateTime2)
	})

	return upcomingShows, nil
}

func GetPastShowCount() ([]map[string]interface{}, error) {
	// Retrieve past show counts from the database
	var pastShowCount []map[string]interface{}
	err := Database.Database.
		Table("schedules").
		Select("setlist, COUNT(*) as count").
		Where("TO_DATE(show_date || ' ' || SUBSTRING(time, 6), 'DD.MM.YYYY HH24:MI') < current_timestamp").
		Group("setlist").
		Find(&pastShowCount).
		Error
	if err != nil {
		return nil, err
	}

	return pastShowCount, nil
}
