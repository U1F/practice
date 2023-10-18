// In a 'datab' package directory
package datab

import (
	"fmt"
	"io"
	"net/http"
	"time"

	_ "github.com/lib/pq"
)

// Assuming we have a global database connection variable in this package (initialized elsewhere, e.g., in connector.go)
// var DB *sql.DB

// CacheEntry represents a cached webpage.
type CacheEntry struct {
	Content   string
	Retrieved time.Time
}

// UserInteraction represents a user's interaction with a webpage element.
type UserInteraction struct {
	UserID    int
	ElementID string // ID of the HTML element interacted with
	Timestamp time.Time
}

// Cache for webpages, key is typically the URL.
var pageCache = make(map[string]CacheEntry)

// CachePage caches the content of a webpage locally.
func CachePage(url string) error {
	// Check if the page is already cached
	if entry, exists := pageCache[url]; exists {
		fmt.Printf("Page retrieved from cache at %s\n", entry.Retrieved)
		return nil // Or you can return the content, based on your needs
	}

	// If not in cache, we retrieve the page
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("error: status code is %d", resp.StatusCode)
	}

	// Read and cache the response
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	bodyString := string(bodyBytes)
	pageCache[url] = CacheEntry{Content: bodyString, Retrieved: time.Now()}

	fmt.Println("Page cached successfully.")
	return nil
}

// LogUserInteraction logs which buttons (or other elements) a user has interacted with.
func LogUserInteraction(interaction UserInteraction) error {
	// This function is typically triggered by some event on the frontend. Here, we're simulating the action.

	// Insert the interaction record into the database. You need to ensure the table structure is correct and already exists in your DB.
	// Here, we are assuming there is a table named 'user_interactions' with columns 'user_id', 'element_id', and 'interaction_time'.
	query := `
		INSERT INTO user_interactions (user_id, element_id, interaction_time)
		VALUES ($1, $2, $3)
	`

	_, err := DB.Exec(query, interaction.UserID, interaction.ElementID, interaction.Timestamp)
	if err != nil {
		return err
	}

	fmt.Println("Logged user interaction:", interaction)
	return nil
}

// Initialize the database connection based on earlier setup from connector.go or similar file.
// You should call this function at the start of your application, ensuring it's done before any database interaction occurs.
func Initialize() {
	// The parameters should be your actual database credentials and information.
	InitializeDB("yourHost", "yourPort", "yourUser", "yourPassword", "yourDBName")
}

// Usage of these functions would be called within the route handlers of your web server,
// where you'd handle the logic for when these should be executed based on client-side interactions and requests.
