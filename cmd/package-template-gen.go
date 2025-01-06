package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {

	if len(os.Args) < 2 {
		log.Fatalln("empty resource name")
	}
	packageName := strings.ToLower(os.Args[1])
	dirName := "./api/pkg/" + packageName

	files := []string{
		"errors.go",
		"https.go",
		"init.go",
		"interface.go",
		"queries.go",
		"repository.go",
		"service.go",
		"structs.go",
	}

	if err := os.Mkdir(dirName, 0755); err != nil {
		fmt.Printf("Failed to create directory: %v\n", err)
		return
	}

	for _, file := range files {
		filePath := fmt.Sprintf("%s/%s", dirName, file)
		f, err := os.Create(filePath)
		if err != nil {
			fmt.Printf("Failed to create file %s: %v\n", file, err)
			continue
		}
		defer f.Close()

		f.WriteString(fmt.Sprintf("package %s\n\n", packageName))

		capitalPName := capitalize(packageName)
		// Write specific content based on file name
		switch file {
		case "interface.go":
			content := fmt.Sprintf(
				"type %sRepository interface {}\n\n"+
					"type %sService interface {}\n",
				capitalPName, capitalPName,
			)
			f.WriteString(content)

		case "https.go":
			content := fmt.Sprintf(
				"type %sHandler struct {\n"+
					"\tservice  %sService\n"+
					"}\n\n"+
					"func New%sHandler(service  %sService) *%sHandler {\n"+
					"\treturn &%sHandler{service: service}\n"+
					"}\n\n"+
					"func (h *%sHandler) Handlers() []rest.HTTPHandler {\n"+
					"\treturn []rest.HTTPHandler{}\n"+
					"}\n",
				capitalPName, capitalPName, capitalPName, capitalPName, capitalPName, capitalPName, capitalPName,
			)
			f.WriteString(content)

		case "service.go":
			content := fmt.Sprintf(
				"type service struct {\n"+
					"\trepo %sRepository\n"+
					"}\n\n"+
					"func New%sService(repo %sRepository) %sService {\n"+
					"\treturn &service{repo: repo}\n"+
					"}\n",
				capitalPName, capitalPName, capitalPName, capitalPName,
			)
			f.WriteString(content)

		case "repository.go":
			content := fmt.Sprintf(
				"type repository struct {\n"+
					"\tdb *sql.DB\n"+
					"}\n\n"+
					"func New%sRepository(db *sql.DB) %sRepository {\n"+
					"\treturn &repository{db: db}\n"+
					"}\n",
				capitalPName, capitalPName,
			)
			f.WriteString(content)

		case "init.go":
			content := fmt.Sprintf(
				"import (\n"+
					"\t\"database/sql\"\n"+
					"\t\"github.io/ckshitij/go-service-template/api/wrapper/rest\"\n"+
					")\n\n"+
					"func Init%s(db *sql.DB) rest.IEndpointProvider {\n"+
					"\trepo := New%sRepository(db)\n"+
					"\tservice := New%sService(repo)\n"+
					"\treturn New%sHandler(service)\n"+
					"}\n",
				capitalize(packageName), capitalize(packageName), capitalize(packageName), capitalize(packageName),
			)
			f.WriteString(content)
		}
	}

	fmt.Println("Module directory and files created successfully!")
}

func capitalize(s string) string {
	if len(s) == 0 {
		return s
	}
	return strings.ToUpper(s[:1]) + s[1:]
}
