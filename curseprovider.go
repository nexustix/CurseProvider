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
// return as URL encoded strings seperated by spaces followed by a Pipe and all Grops seperated by spaces
// <Provider> <Name> <ID> <Filename> <URL> <RelativePath>|<Group1> <Group2> <Group3>...

//depsearch => get dependencies
// tbd

//downinfo => get download URL and filename
// tbd

// return as URL encoded strings seperated by spaces followed by a Pipe and all Grops seperated by spaces
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

	providerAction := bp.StringAtIndex(1, args)
	providerQuerry := bp.StringAtIndex(2, args)

	switch providerAction {

	case "search":
		fmt.Printf("<-> search for %s\n", providerQuerry)
		searchPhrase := nxcurse.GetMinecraftModSearchphrase(providerQuerry)
		searchURL := nxduck.GenerateSearchURL(searchPhrase)
		searchResults := nxduck.GetSearchResultObjects(searchURL)
		curseResults := nxcurse.GetMinecraftModResults(searchResults)
		for _, v := range curseResults {

			urlSegments = strings.Split(v.URL, "/")

			fmt.Printf("%s %s %s %s %s %s|",
				url.QueryEscape("curse"),                         // Provider
				url.QueryEscape(getModName(v)),                   // Name
				url.QueryEscape(urlSegments[len(urlSegments)-2]), //ID
				url.QueryEscape(""),                              //Filename
				url.QueryEscape(v.URL),                           //URL (this provider does not provide direct download URLs here)
				url.QueryEscape("mods"),                          //RelativePath (install to mod folder)
			)
			//XXX no URL encoding ?
			for kk, vv := range groups {
				if kk == len(groups)-1 {
					//fmt.Println(vv) //last item should not have trailing space character
					fmt.Print(vv + "\n") //last item should not have trailing space character
				} else {
					fmt.Print(vv + " ")
				}
			}
		}

	case "depsearch":
		//fmt.Printf("<-> depsearch for %s\n", bp.StringAtIndex(2, args))
		deps := nxcurse.GetDependencies(providerQuerry)

		//fmt.Printf("<-> found %s\n", deps)

		for kk, vv := range deps {
			if kk == len(deps)-1 {
				fmt.Print(vv) //last item should not have trailing space character
			} else {
				fmt.Print(vv + " ")
				//fmt.Println(vv)
			}
		}
		fmt.Println()

	case "downinfo":
		//fmt.Printf("<-> downinfo for %s\n", bp.StringAtIndex(2, args))
		downinfo := nxcurse.GetCurseDownloads("http://minecraft.curseforge.com/projects/"+bp.StringAtIndex(2, args)+"/files", "1.10")
		if len(downinfo) >= 1 {
			//print URL first since URLs are more predictable due to URL encoding

			tmpFilename := nxcurse.GetFilenameFromDownloadURL(downinfo[0].URL)
			//XXX overwriting cheap-ish filename
			fmt.Printf("%s|%s\n", downinfo[0].URL, tmpFilename)

		}

	}
	//fmt.Printf("it works !\n")
}
