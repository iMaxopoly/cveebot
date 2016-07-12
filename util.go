package main

import "log"

/***
 *      _    _ _   _ _ _ _
 *     | |  | | | (_) (_) |
 *     | |  | | |_ _| |_| |_ _   _
 *     | |  | | __| | | | __| | | |
 *     | |__| | |_| | | | |_| |_| |
 *      \____/ \__|_|_|_|\__|\__, |
 *                            __/ |
 *                           |___/
 */

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func debugPrint(msg ...interface{}) {
	if DEBUG {
		log.Println(msg)
	}
}

func extractApplicationID(url string) string {
	var end int
	for i := len(url) - 1; i >= 0; i-- {
		if string(url[i]) == "." {
			end = i
		} else if string(url[i]) == "-" {
			return url[i+2 : end]
		}
	}
	return ""
}

func stringInArray(value string, array []string) bool {
	for _, v := range array {
		if v == value {
			return true
		}
	}
	return false
}
