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

	if err := os.MkdirAll(dirName, 0755); err != nil {
		log.Fatalf("Failed to create directory: %v\n", err)
	}

	for _, file := range files {
		filePath := fmt.Sprintf("%s/%s", dirName, file)
		err := createFile(filePath, file, packageName)
		if err != nil {
			log.Printf("Failed to create file %s: %v\n", file, err)
		}
	}

	fmt.Println("Module directory and files created successfully!")
}

func createFile(filePath, fileName, packageName string) error {
	f, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("unable to create file: %w", err)
	}
	defer f.Close()

	packageDeclaration := fmt.Sprintf("package %s\n\n", packageName)
	// Write the package declaration
	if _, err := f.WriteString(packageDeclaration); err != nil {
		return fmt.Errorf("unable to write package declaration: %w", err)
	}

	// Write the file-specific content
	content := getFileContent(fileName, packageName)
	if content != "" {
		if _, err := f.WriteString(content); err != nil {
			return fmt.Errorf("unable to write content: %w", err)
		}
	}

	return nil
}

func getFileContent(fileName, packageName string) string {
	capitalPName := capitalize(packageName)

	switch fileName {
	case "interface.go":
		return fmt.Sprintf(
			"type %sRepository interface {}\n\n"+
				"type %sService interface {}\n",
			capitalPName, capitalPName,
		)

	case "https.go":
		return fmt.Sprintf(
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

	case "service.go":
		return fmt.Sprintf(
			"type service struct {\n"+
				"\trepo %sRepository\n"+
				"}\n\n"+
				"func New%sService(repo %sRepository) %sService {\n"+
				"\treturn &service{repo: repo}\n"+
				"}\n",
			capitalPName, capitalPName, capitalPName, capitalPName,
		)

	case "repository.go":
		return fmt.Sprintf(
			"type repository struct {\n"+
				"\tdb *sql.DB\n"+
				"}\n\n"+
				"func New%sRepository(db *sql.DB) %sRepository {\n"+
				"\treturn &repository{db: db}\n"+
				"}\n",
			capitalPName, capitalPName,
		)

	case "init.go":
		return fmt.Sprintf(
			"import (\n"+
				"\t\"database/sql\"\n"+
				"\t\"github.io/ckshitij/go-service-template/api/wrapper/rest\"\n"+
				")\n\n"+
				"func Init%s(db *sql.DB) rest.IEndpointProvider {\n"+
				"\trepo := New%sRepository(db)\n"+
				"\tservice := New%sService(repo)\n"+
				"\treturn New%sHandler(service)\n"+
				"}\n",
			capitalPName, capitalPName, capitalPName, capitalPName,
		)

	default:
		return ""
	}
}

func capitalize(s string) string {
	if s == "" {
		return s
	}
	return strings.ToUpper(s[:1]) + s[1:]
}
