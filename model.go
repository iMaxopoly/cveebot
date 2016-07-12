package main

/***
 *      __  __           _      _
 *     |  \/  |         | |    | |
 *     | \  / | ___   __| | ___| |___
 *     | |\/| |/ _ \ / _` |/ _ \ / __|
 *     | |  | | (_) | (_| |  __/ \__ \
 *     |_|  |_|\___/ \__,_|\___|_|___/
 *
 *
 */

type jobCategoryModel struct {
	Name string
	URL  string
}

type jobApplicationModel struct {
	Name   string
	URL    string
	FormID string
}
