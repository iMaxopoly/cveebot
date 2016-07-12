package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"sync"

	"github.com/PuerkitoBio/goquery"
)

/***
 *      __  __      _   _               _
 *     |  \/  |    | | | |             | |
 *     | \  / | ___| |_| |__   ___   __| |___
 *     | |\/| |/ _ \ __| '_ \ / _ \ / _` / __|
 *     | |  | |  __/ |_| | | | (_) | (_| \__ \
 *     |_|  |_|\___|\__|_| |_|\___/ \__,_|___/
 *
 *
 */

var cookieLock sync.Mutex

func setCookies(cookies string) {
	cookieLock.Lock()
	defer cookieLock.Unlock()

	gCookies = cookies
}

func getCookies() string {
	cookieLock.Lock()
	defer cookieLock.Unlock()

	return gCookies
}

func dialer(url string) string {
	// Ready request
	req, err := http.NewRequest("GET", url, nil)
	checkError(err)

	// Set headers
	req.Header.Set("User-Agent", userAgent)
	req.Header.Set("Content-Type", contentTypeFormEncoded)
	req.Header.Set("Cookie", getCookies())

	// Make request
	client := http.Client{}
	resp, err := client.Do(req)
	checkError(err)

	// Read Response
	b, err := ioutil.ReadAll(resp.Body)
	checkError(err)
	body := string(b)

	// Close Response Body
	err = resp.Body.Close()
	checkError(err)

	// Check if session timed out
	if strings.Contains(body, "Wrong username or password!") {
		setCookies("")
		doLogin()
	}

	return body
}

func doLogin() bool {
	// Setup form
	form := url.Values{}

	form.Add("module", "login")
	form.Add("username", username)
	form.Add("password", password)

	// Ready request
	req, err := http.NewRequest("POST", urlLoginPost, strings.NewReader(form.Encode()))
	checkError(err)

	// Set headers
	req.Header.Set("User-Agent", userAgent)
	req.Header.Set("Content-Type", contentTypeFormEncoded)
	req.Header.Set("Cookie", cookieSetEnglish)

	// Make request
	client := http.Client{}
	resp, err := client.Do(req)
	checkError(err)

	// Read Response
	b, err := ioutil.ReadAll(resp.Body)
	checkError(err)
	body := string(b)

	// Close Response Body
	err = resp.Body.Close()
	checkError(err)

	// Verify response
	if strings.Contains(body, "Wrong username or password!") {
		return false
	}

	setCookies(resp.Header.Get("Set-Cookie") + "; " + cookieSetEnglish)
	return true
}

func isLoggedIn() bool {
	req, err := http.NewRequest("GET", urlLoginForm, nil)
	checkError(err)
	req.Header.Set("Cookie", getCookies())
	req.Header.Set("User-Agent", userAgent)
	req.Header.Set("Content-Type", contentTypeTextHTML)

	client := http.Client{}
	resp, err := client.Do(req)
	checkError(err)

	return resp.Header.Get("Location") != "/seeker/index.php"
}

func collectJobCategories() []jobCategoryModel {
	doc, err := goquery.NewDocument(urlJobListingCategories)
	checkError(err)

	var categories []jobCategoryModel

	docJobCats := doc.Find("#Tegvk_body").Find("li")
	for node := range docJobCats.Nodes {

		j := jobCategoryModel{}

		thisNode := docJobCats.Eq(node).Find("a")

		url, exists := thisNode.Attr("href")
		if !exists {
			panic(errors.New("URL not found for JobCategory"))
		}

		// If we are ignoring this link
		if stringInArray(url, ignoredJobCategories) {
			continue
		}

		name := thisNode.Text()

		j.URL, j.Name = siteURL+url, name

		categories = append(categories, j)

	}
	return categories
}

func collectJobs() []jobApplicationModel {

	jobs := []jobApplicationModel{}

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(dialer(fmt.Sprintf(urlJobListing, "information-technology", 0))))
	checkError(err)

	jobTitleDoc := doc.Find(".contentJobTitle")

	for node := range jobTitleDoc.Nodes {
		name := jobTitleDoc.Eq(node).Text()
		link, exists := jobTitleDoc.Eq(node).Attr("href")
		if !exists {
			panic(errors.New("URL not found for Job List"))
		}

		if strings.Contains(link, "plid") {
			continue
		}

		s := jobApplicationModel{}
		s.Name = name
		// Make it completely valid link
		s.URL = "http:" + link
		s.FormID = extractApplicationID(s.URL)

		// Verify the extraction succeeded
		if s.FormID == "" {
			panic(errors.New("Extracted Application ID is nil"))
		}

		jobs = append(jobs, s)
	}

	return jobs
}

func applyForJob(formID string) {
	url := fmt.Sprintf(urlJobApply, formID)

}
