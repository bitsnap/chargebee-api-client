func Retrieve{{.NameSingular}}ById(site string, id string) (*{{.NameSingular}}, error) {
	client := fiber.AcquireClient()
	defer fiber.ReleaseClient(client)

	parsedUrl, err := url.Parse("https://" + site + "." + "{{.Url}}/" + id)
	if err != nil {
		return nil, err
	}
	
    content, err := getQuery(client, parsedUrl, "", nil)
    if err != nil {
        return nil, err
    }
    	
	type {{.NameSingular}}Item struct {
		{{.NameSingular}} {{.NameSingular}} `json:"{{.NameSingular}}"`
	}

    var item {{.NameSingular}}Item

	err = json.Unmarshal(content, &item)
	if err != nil {
		return nil, err
	}

	return &item.{{.NameSingular}}, nil
}
