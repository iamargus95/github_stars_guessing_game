package githubstars

import (
	"os"
	"reflect"
	"testing"
)

func TestGithub_stars(t *testing.T) {
	data, _ := os.ReadFile("./tests/response.json")

	jsonResponse := responseToJson(data)

	want := SearchData{
		Items: []RepoInfo{
			{12821113, "uncss/grunt-uncss", ":scissors: A grunt task for removing unused CSS from your projects.", 3864, 182, "HTML"},
			{3022431, "addyosmani/backbone-fundamentals", ":book: A creative-commons book on Backbone.js for beginners and advanced users alike", 9395, 1511, "JavaScript"},
			{30165704, "Mango/slideout", "A touch slideout navigation menu for your mobile web apps.", 7972, 1240, "JavaScript"},
			{21486287, "sahat/satellizer", "Token-based AngularJS Authentication", 7963, 1195, "TypeScript"},
		},
	}

	if !reflect.DeepEqual(jsonResponse, want) {
		t.Fatal("JSON unmarshall failed.")
	}
}
