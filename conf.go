package main

/***
 *       _____             __ _
 *      / ____|           / _(_)
 *     | |     ___  _ __ | |_ _  __ _
 *     | |    / _ \| '_ \|  _| |/ _` |
 *     | |___| (_) | | | | | | | (_| |
 *      \_____\___/|_| |_|_| |_|\__, |
 *                               __/ |
 *                              |___/
 */

// DEBUG variable allows us to watch the progress of the bot.
const DEBUG = true

// cv.ee address
const siteURL = "http://www.cv.ee"

// cv.ee Login username
const username = "contact@kryptodev.com"

// cv.ee Login password
const password = "wszGLm6XMIacPzLrmSZR"

// URL where server receives login data via POST
const urlLoginPost = "http://www.cv.ee/services/xhtml/index.php/loginmessage"

// URL where login form is displayed
const urlLoginForm = "http://www.cv.ee/for-jobseeker/login"

// URL for Job Listing
const urlJobListing = "http://www.cv.ee/job-ads/%s?page=%d"

// URL for Job Listing Categories
const urlJobListingCategories = "http://www.cv.ee/job-ads/"

// URL for application form
const urlJobApply = "http://www.cv.ee/apply/%s"

// Content-type text-html
const contentTypeTextHTML = "text/html;charset=UTF-8"

// Content-type url form encoded
const contentTypeFormEncoded = "application/x-www-form-urlencoded; charset=UTF-8"

// Mock User agent
const userAgent = "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/51.0.2704.106 Safari/537.36"

// To see in english
const cookieSetEnglish = "cookielang=inglise"

// Job categories I deem unfit
var ignoredJobCategories = []string{
	"http://www.cv.ee/kuulutused/pakun",
	"http://www.cv.lv/job-ads",
	"http://www.cvonline.lt/job-ads",
	"http://www.monster.fi/",
}

// Cover letter
var coverLetter = `Dear Hiring Manager,

With great interest, I would earnestly like you to consider me for a job position mentioned in your advertisement "%s" on http://www.cv.ee/

I am deeply excited to learn about the availability and firmly believe that I possess all the necessary skills you are looking for in an ideal candidate for the task. Kindly refer to the enclosed resume along with this letter.

I have been utterly enthusiastic about computers from an early age, and as a result, managed to find myself in various ethical hacking communities and forums. My deep interest in the field led me through multitude of topics ranging from ethical web hacking to reverse engineering and ultimately, choosing to become a developer.

Choosing to become a developer came to me at around age 15 where I started my first private game server which allowed me to be an author of my own product. Managing my linux based server on a daily basis, tightening firewall, working with various CMS, database management and learning C++ to make a game launcher provoked me to learn more than ever I thought I could.

This journey took me through a multitude of projects, most that I owned myself where I programmed to code online shopping carts, paypal integrations and more.

Over years I have matured and have developed great skills and experience with multiple languages and platforms that include but are not limited to HTML, CSS, Java, PHP, Hack, Python, Scala, Kotlin, C, C++, Lua, Javascript​, Ruby, C#, and my most beloved Golang.

In the past I have run several websites, a blogging platform called newjolt.com, now deprecated, a game server at windsofarithia.com, now deprecated, and am now the sole owner and developer of http://animedom.com which is an anime video network containing over 3 thousand anime series. The website runs on golang codebase with rethinkdb database structure. I have plans to make this website into a single page app with the stable release on Angular 2.

I am a quick learner, an adept google searcher, and have the capacity to learn any new programming language or technology in a matter of days, sometimes staying up entire nights until I get it right.

At the time of posting this application, I would like to make you aware that this is being posted by a bot that I created to demonstrate my quick skills. The code for which is available(at https://github.com/kryptodev/cveebot) for your assessment.

Before ending my cover-letter, I would encourage you to interview me by means of skype, teamviewer or any other manner of your choice. I love challenges and I would enjoy being tested well by my future employer where I see myself working on a longterm basis.

Considering I manage to attract your interest, I can manage to pay for my own expenses to reach Estonia, where my sister is currently studying Software Engineering(in University of Tartu).

Thank you for your patience and I intently look forward to a future at your enterprise for a long and productive duration.

Most Sincerely,
Manish Prakash Singh

Enclosed herewith: Resume`
