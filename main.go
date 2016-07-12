package main

/***
 *      __  __       _
 *     |  \/  |     (_)
 *     | \  / | __ _ _ _ __
 *     | |\/| |/ _` | | '_ \
 *     | |  | | (_| | | | | |
 *     |_|  |_|\__,_|_|_| |_|
 *
 *
 */

var gCookies string

func main() {

	// Step 1: Collect Job Categories
	categories := collectJobCategories()
	debugPrint("Categories collected:", categories)

	// Step 2: Login
	loginSuccess := doLogin()
	if loginSuccess {
		debugPrint("Logged in successfully, received cookies:", gCookies)
	} else {
		debugPrint("Wrong Login Credentials")
	}

	// Step 3: Verify if login cookies are valid
	debugPrint("Are we logged in?", isLoggedIn())

	// Step 4: Collect job applications
	jobForms := collectJobs()
	debugPrint("Job forms", jobForms)

	// Step 5: Apply forms

}
