package cmd

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/manifoldco/promptui"
)

func parseInput(search string, lang string, desc string) url.Values {
	queryString := fmt.Sprintf("%s in:name", search)
	if lang != "" {
		queryString = queryString + fmt.Sprintf(" language:%s", lang)
	}
	if desc != "" {
		queryString = queryString + fmt.Sprintf(" %s in:description", desc)
	}
	query := url.Values{}
	query.Add("q", queryString)
	query.Add("sort", "stars")
	query.Add("per_page", "30")
	return query
}

func getTemplate() *promptui.SelectTemplates {
	funcMap := promptui.FuncMap
	funcMap["parseStars"] = func(starCount float64) string {
		if starCount >= 1000 {
			return fmt.Sprintf("%.1f k", starCount/1000)
		}
		return fmt.Sprint(starCount)
	}

	funcMap["truncate"] = func(input string) string {
		length := 80
		if len(input) <= length {
			return input
		}
		return input[:length-3] + "..."
	}

	return &promptui.SelectTemplates{
		Active:   "\U0001F449 {{ .Name | cyan | bold }}",
		Inactive: "   {{ .Name | cyan }}",
		Selected: `{{ "✔" | green | bold }} {{ .Name | cyan | bold }}`,
		Details: `
	{{ "Name:" | faint }} 	 {{ .Name }}
	{{ "Description:" | faint }} 	 {{ .Description | truncate }}
	{{ "Url address:" | faint }} 	 {{ .URL }}
	{{ "⭐" | faint }}	{{ .Stars | parseStars }}`,
	}

}

func getSelectionPrompt(repos []repoInfo) *promptui.Select {
	return &promptui.Select{
		Label:     "repository list",
		Items:     repos,
		Templates: getTemplate(),
		Size:      20,
		Searcher: func(input string, idx int) bool {
			repo := repos[idx]
			title := strings.ToLower(repo.Name)

			return strings.Contains(title, input)
		},
	}
}