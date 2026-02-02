// package utils

// import (
//     "fmt"
//     "io"
//     "log"
//     "net/http"
//     "net/url"
//     "os"
//     "strings"
// )

// func SendSMS(message string, mobileArray []string) error {
//     apiKey := os.Getenv("SMS_API_KEY")
//     senderID := os.Getenv("SMS_SENDER_ID")
//     apiURL := os.Getenv("SMS_API_URL")

//     if apiKey == "" || senderID == "" || apiURL == "" {
//         return fmt.Errorf("missing SMS config in environment")
//     }

//     numbers := strings.Join(mobileArray, ",")

//     params := url.Values{}
//     params.Set("api_key", apiKey)
//     params.Set("type", "text")
//     params.Set("number", numbers)
//     params.Set("senderid", senderID)
//     params.Set("message", message)

//     finalURL := fmt.Sprintf("%s?%s", apiURL, params.Encode())

//     resp, err := http.Get(finalURL)
//     if err != nil {
//         return fmt.Errorf("failed to send SMS: %v", err)
//     }
//     defer resp.Body.Close()

//     body, _ := io.ReadAll(resp.Body)
//     if resp.StatusCode != http.StatusOK {
//         return fmt.Errorf("SMS API error: %s", string(body))
//     }

//	    log.Printf("SMS sent successfully: %s", string(body))
//	    return nil
//	}
package utils

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func SendSMS(message string, mobileArray []string) error {
    apiKey := os.Getenv("SMS_API_KEY")
    senderID := os.Getenv("SMS_SENDER_ID")
    apiURL := os.Getenv("SMS_API_URL")

    if apiKey == "" || senderID == "" || apiURL == "" {
        return fmt.Errorf("missing SMS config in environment")
    }

    numbers := strings.Join(mobileArray, ",")
	if !strings.HasPrefix(numbers, "88") { numbers = "88" + numbers }

    params := url.Values{}
    params.Set("api_key", apiKey)
    params.Set("type", "text")
    params.Set("number", numbers)
    params.Set("senderid", senderID)
    params.Set("message", message)

    finalURL := fmt.Sprintf("%s?%s", apiURL, params.Encode())

    // 🔥 Debug: URL print
    fmt.Println("Final SMS URL:", finalURL)

    resp, err := http.Get(finalURL)
    if err != nil {
        return fmt.Errorf("failed to send SMS: %v", err)
    }
    defer resp.Body.Close()

    body, _ := io.ReadAll(resp.Body)

    // 🔥 Debug: API response print
    fmt.Println("SMS API Response:", string(body))

    if resp.StatusCode != http.StatusOK {
        log.Println("SMS API Error:", string(body))
        return fmt.Errorf("SMS API error: %s", string(body))
    }

    log.Printf("SMS sent successfully: %s", string(body))
    return nil
}
