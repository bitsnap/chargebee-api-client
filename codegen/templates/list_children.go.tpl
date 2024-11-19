func List{{.Name}}PageSortBy(site string, id string, offset string, sortBy *SortBy) ([]{{.NameSingular}}, string, error) {
	client := fiber.AcquireClient()
	defer fiber.ReleaseClient(client)

	parsedUrl, err := url.Parse("https://" + site + "." + "{{.Url}}/{{ if .AddSuffix }}" + id + "/{{ end }}{{.UrlSuffix}}")
	if err != nil {
		return nil, "", err
	}
	{{ if not .AddSuffix }}
	q := parsedUrl.Query()
	q.Add("{{ .RenameId }}[is]", id)
	parsedUrl.RawQuery = q.Encode()
	{{ end }}	
    content, err := GetQuery(client, parsedUrl, offset, sortBy)
    if err != nil {
        return nil, "", err
    }
    	
	type {{.NameSingular}}ListItem struct {
		{{.NameSingular}} {{.NameSingular}} `json:"{{.NameSingular}}"`
	}

    type {{.NameSingular}}Page struct {
        List       []{{.NameSingular}}ListItem `json:"list"`
        NextOffset string `json:"next_offset,omitempty"`
    }

	entries := {{.NameSingular}}Page{
		List:       make([]{{.NameSingular}}ListItem, 0, 10),
	}

	err = json.Unmarshal(content, &entries)
	if err != nil {
		return nil, "", err
	}
	
	if len(entries.List) == 0 {
        return []{{.NameSingular}}{}, "", nil
    }
	
	result := make([]{{.NameSingular}}, 0, len(entries.List))
	for _, r := range entries.List {
		result = append(result, r.{{.NameSingular}})
	}

    if len(entries.NextOffset) > 0 {
        return ResultWithOffset(result, offset, entries.NextOffset)
    }
	
	return result, "", nil
}

func List{{.Name}}Page(site string, id string, offset string) ([]{{.NameSingular}}, string, error) {
	return List{{.Name}}PageSortBy(site, id, offset, nil)
}
