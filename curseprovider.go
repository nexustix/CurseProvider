package main

import (
	"fmt"
	"net/url"
	"os"
	"strings"

	bp "github.com/nexustix/boilerplate"
	"github.com/nexustix/nxcurse"
	"github.com/nexustix/nxduck"
)

//search => find atoms
//deps => get dependencies
//downinfo => get download URL and filename

// return as URL encoded strings seperated by spaces followed by a Pipe and all Grops seperaded by spaces
// <Provider> <Name> <ID> <Filename> <URL> <RelativePath>|<Group1> <Group2> <Group3>...

//TODO evaluate (old code)
func getModName(theResult nxduck.SearchResult) string {
	var modName string
	//segments := strings.Split(theResult.Title, "-")
	segments := strings.Split(theResult.Title, " - ")
	//fmt.Println(strings.TrimSpace(segments[1]))
	switch segments[0] {
	case "Overview":
		modName = strings.TrimSpace(segments[1])
	case "Files":
		modName = strings.TrimSpace(segments[1])
	case "Addons":
		modName = strings.TrimSpace(segments[1])
	case "Images":
		modName = strings.TrimSpace(segments[1])
	default:
		modName = strings.TrimSpace(segments[0])
	}

	//if strings.HasPrefix(modName, "...") {
	if strings.Contains(modName, "...") {
		urlSegments := strings.Split(theResult.URL, "/")
		modName = urlSegments[4]

		modName = strings.Replace(modName, "-", " ", -1)

		modName = strings.Title(modName)
	}
	return modName
}

func main() {
	//version := "V.0-1-0"

	args := os.Args

	var urlSegments []string

	var groups []string

	groups = append(groups, "curse")
	groups = append(groups, "minecraft")

	switch bp.StringAtIndex(1, args) {

	case "search":
		fmt.Printf("<-> search for %s\n", bp.StringAtIndex(2, args))
		searchPhrase := nxcurse.GetMinecraftModSearchphrase(bp.StringAtIndex(2, args))
		searchURL := nxduck.GenerateSearchURL(searchPhrase)
		searchResults := nxduck.GetSearchResultObjects(searchURL)
		curseResults := nxcurse.GetMinecraftModResults(searchResults)
		for _, v := range curseResults {

			urlSegments = strings.Split(v.URL, "/")

			fmt.Printf("%s %s %s %s %s %s|",
				url.QueryEscape("curse"),                         // Provider
				url.QueryEscape(getModName(v)),                   // Name
				url.QueryEscape(urlSegments[len(urlSegments)-1]), //ID
				url.QueryEscape("N/A"),                           //Filename
				url.QueryEscape(v.URL),                           //URL (this provider does not provide direct download URLs here)
				url.QueryEscape("/mods"),                         //RelativePath (install to mod folder)
			)
			for kk, vv := range groups {
				if kk == len(groups)-1 {
					fmt.Println(vv) //last item should not have trailing space character
				} else {
					fmt.Print(vv + " ")
				}
			}
		}

	case "deps":
		fmt.Printf("<-> depsearch for %s\n", bp.StringAtIndex(2, args))

	case "downinfo":
		fmt.Printf("<-> depsearch for %s\n", bp.StringAtIndex(2, args))

	}
	//fmt.Printf("it works !\n")
}
