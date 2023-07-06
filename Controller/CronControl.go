package Controller

import (
	"liamelior-api/Model"
	"net/http"
	"os"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/gin-gonic/gin"
)

type ScrapeResult struct {
	Data []map[string]string `json:"data"`
}

func ScrapeHandler(c *gin.Context) {
	// Send GET request to the website
	url := os.Getenv("JKT48_SCHEDULE_URL")
	resp, err := http.Get(url)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to fetch data from the website",
		})
		return
	}
	defer resp.Body.Close()

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to parse HTML",
		})
		return
	}

	// Search for the desired member
	found := false

	doc.Find("table.table tbody tr").Each(func(index int, row *goquery.Selection) {
		// Extract the member data
		memberName := os.Getenv("JKT48_MEMBER_NAME")
		memberData := row.Find("td:nth-child(3)").Text()
		if strings.Contains(memberData, memberName) {
			found = true
			setlist := row.Find("td:nth-child(2)").Text()
			showDateTime := row.Find("td:nth-child(1)").Text()

			// Extract the show date and time using modified regular expressions
			dateRegex := regexp.MustCompile(`\w+,\s*(\d+\.\d+\.\d+)`)
			timeRegex := regexp.MustCompile(`(Show\s*\d+:\d+)`)
			showDateMatches := dateRegex.FindStringSubmatch(showDateTime)
			showTime := timeRegex.FindString(showDateTime)

			if len(showDateMatches) > 1 {
				showDate := showDateMatches[1]

				// save the map to database to model Schedule

				// Check if the data already exists in the database
				exists, err := Model.CheckScheduleExists(setlist, showDate, showTime)
				if err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{
						"error": "Failed to check data in the database",
					})
					return
				}

				// If the data doesn't exist, save it to the database
				if exists {
					c.JSON(http.StatusOK, gin.H{
						"message": "Successfully scraped data from the website, no new data has been saved to the database",
					})
					return

				} else {
					_, err := (&Model.Schedule{
						Setlist:  setlist,
						ShowDate: showDate,
						Time:     showTime,
					}).Save()

					if err != nil {
						// Handle the error if saving fails
						c.JSON(http.StatusInternalServerError, gin.H{
							"error": "Failed to save data into the database",
						})
						return
					}

					// Create a response with the extracted data
					response := gin.H{"message": "Successfully scraped data from the website, new data has been saved to the database"}

					c.JSON(http.StatusOK, response)
				}
			}
		}
	})

	// Check if member was found
	if !found {
		c.JSON(http.StatusOK, gin.H{
			"message": "There are no upcoming shows with the specified member",
		})
		return
	}

}
