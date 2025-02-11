package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

// Template directory
const templateDir = "templates"

// Template file mapping
var templateFiles = map[string]string{
	"model":      filepath.Join(templateDir, "model.go.tmpl"),
	"repository": filepath.Join(templateDir, "repository.go.tmpl"),
	"service":    filepath.Join(templateDir, "service.go.tmpl"),
	"handler":    filepath.Join(templateDir, "handler.go.tmpl"),
	"request":    filepath.Join(templateDir, "request.go.tmpl"),
	"routes":     filepath.Join(templateDir, "routes.go.tmpl"),
	"provider":   filepath.Join(templateDir, "provider.go.tmpl"),
}

func die(v ...interface{}) {
	fmt.Println(v...)
	os.Exit(1)
}

func main() {
	if len(os.Args) < 3 || os.Args[1] != "make:module" {
		fmt.Println("Usage: go run cmd/modular/main.go make:module <module_name>")
		return
	}

	moduleName := os.Args[2]
	if err := GenerateModule(moduleName); err != nil {
		log.Fatalf("Error: %v", err)
	}
}

// GenerateModule creates a new module
func GenerateModule(moduleName string) error {
	moduleName = strings.Title(moduleName) // Capitalize first letter
	moduleLower := strings.ToLower(moduleName)
	moduleImport := "SparksSport/pkg/" + moduleLower

	// Define module paths
	moduleBase := filepath.Join("../../pkg", moduleLower)
	modulePaths := []string{
		filepath.Join(moduleBase, "http", "handlers"),
		filepath.Join(moduleBase, "http", "routes"),
		filepath.Join(moduleBase, "http", "requests"),
		filepath.Join(moduleBase, "repositories"),
		filepath.Join(moduleBase, "models"),
		filepath.Join(moduleBase, "services"),
		filepath.Join(moduleBase, "providers"),
	}

	if checkModuleExistence(moduleBase) {
		return fmt.Errorf("module '%s' already exists. Skipping generation.", moduleName)
	}

	// Define template-generated file paths
	files := map[string]string{
		"model":      filepath.Join(moduleBase, "models", moduleLower+".go"),
		"repository": filepath.Join(moduleBase, "repositories", moduleLower+"_repository.go"),
		"service":    filepath.Join(moduleBase, "services", moduleLower+"_service.go"),
		"handler":    filepath.Join(moduleBase, "http", "handlers", moduleLower+"_handler.go"),
		"routes":     filepath.Join(moduleBase, "http", "routes", "api.go"),
		"request":    filepath.Join(moduleBase, "http", "requests", moduleLower+"_routes.go"),
		"provider":   filepath.Join(moduleBase, "providers", moduleLower+"_provider.go"),
	}

	// Create directories
	for _, path := range modulePaths {
		if err := os.MkdirAll(path, os.ModePerm); err != nil {
			return fmt.Errorf("failed to create directory '%s': %w", path, err)
		}
	}

	// Create files using templates
	for key, filePath := range files {
		templatePath := templateFiles[key]
		if err := createFileFromTemplate(filePath, templatePath, moduleName, moduleImport, moduleLower); err != nil {
			return fmt.Errorf("failed to create file '%s': %w", filePath, err)
		}
		log.Printf("Created: %s\n", filePath)
	}

	// Update migrate.go
	if err := updateMigrateFile(moduleName, moduleLower); err != nil {
		log.Fatalf("Failed to update migrate.go: %s", err)
	}

	// register the service provider in app
	if err := updateAppGo(moduleName, moduleLower); err != nil {
		log.Fatalf("Failed to update app.go: %s", err)
	}

	fmt.Printf("Module '%s' created successfully!\n", moduleName)
	return nil
}

// createFileFromTemplate generates a file using a template
func createFileFromTemplate(filePath, templatePath, moduleName, moduleImport, moduleLower string) error {
	// Ensure template file exists
	if _, err := os.Stat(templatePath); os.IsNotExist(err) {
		return fmt.Errorf("template file '%s' not found", templatePath)
	}

	// Read template file content
	tmplContent, err := os.ReadFile(templatePath)
	if err != nil {
		return fmt.Errorf("failed to read template file: %w", err)
	}

	// Parse template
	tmpl, err := template.New(filepath.Base(templatePath)).Parse(string(tmplContent))
	if err != nil {
		return fmt.Errorf("failed to parse template: %w", err)
	}

	// Create output file
	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer file.Close()

	// Fill template with module data
	data := map[string]string{
		"ModuleName":   moduleName,
		"ModuleImport": moduleImport,
		"ModuleLower":  moduleLower,
	}

	// Execute template
	if err := tmpl.Execute(file, data); err != nil {
		return fmt.Errorf("failed to execute template: %w", err)
	}

	return nil
}

func updateMigrateFile(moduleName, moduleLower string) error {
	migrateFile := "../app/core/migrations/migrate.go"

	// Read migrate.go
	content, err := os.ReadFile(migrateFile)
	if err != nil {
		return fmt.Errorf("failed to read migrate.go: %w", err)
	}

	// Convert content to string
	migrateContent := string(content)

	// Prepare import and migration lines
	importLine := fmt.Sprintf(`%sModel "SparksSport/pkg/%s/models"`, moduleLower, moduleLower)
	migrationLine := fmt.Sprintf("&%sModel.%s{},\n", moduleLower, moduleName)

	// Check if the import already exists
	if strings.Contains(migrateContent, importLine) {
		fmt.Println("Module already exists in migrate.go. Skipping update.")
		return nil
	}

	// Add import statement
	importInsertPos := strings.Index(migrateContent, ")")
	if importInsertPos == -1 {
		return fmt.Errorf("failed to find import block in migrate.go")
	}
	migrateContent = migrateContent[:importInsertPos] + "\t" + importLine + "\n" + migrateContent[importInsertPos:]

	// Add migration statement
	migrateFuncPos := strings.Index(migrateContent, "db.AutoMigrate(")
	if migrateFuncPos == -1 {
		return fmt.Errorf("failed to find db.AutoMigrate block in migrate.go")
	}
	insertPos := strings.Index(migrateContent[migrateFuncPos:], ")") + migrateFuncPos
	migrateContent = migrateContent[:insertPos] + "\n\t\t" + migrationLine + migrateContent[insertPos:]

	// Write back to migrate.go
	err = os.WriteFile(migrateFile, []byte(migrateContent), 0644)
	if err != nil {
		return fmt.Errorf("failed to write to migrate.go: %w", err)
	}

	fmt.Println("migrate.go updated successfully.")
	return nil
}

// Example of inserting new provider dynamically
func updateAppGo(moduleName, moduleLower string) error {
	appFile := "../app/bootstrap/app.go"

	// Read app.go content
	content, err := os.ReadFile(appFile)
	if err != nil {
		return fmt.Errorf("failed to read app.go: %w", err)
	}

	// Convert content to string
	appContent := string(content)

	// Prepare import and provider registration lines
	importLine := fmt.Sprintf(`%sServiceProviders "%s/pkg/%s/providers"`, moduleLower, "SparksSport", moduleLower)
	providerLine := fmt.Sprintf("&%sServiceProviders.%sServiceProvider{},", moduleLower, moduleName)

	// Check if the import already exists
	if strings.Contains(appContent, importLine) {
		fmt.Println("Module already exists in app.go. Skipping update.")
		return nil
	}

	// Add import statement before the providers block
	importInsertPos := strings.Index(appContent, ")")
	if importInsertPos == -1 {
		return fmt.Errorf("failed to find import block in app.go")
	}
	appContent = appContent[:importInsertPos] + "\t" + importLine + "\n" + appContent[importInsertPos:]

	// Find the RegisterProviders function and the providers slice declaration
	registerFuncPos := strings.Index(appContent, "providers := []providers.ServiceProvider{")
	if registerFuncPos == -1 {
		return fmt.Errorf("failed to find provider registration block in app.go")
	}

	// Insert the provider line before the slice declaration
	insertPos := registerFuncPos + len("providers := []providers.ServiceProvider{")
	appContent = appContent[:insertPos] + "\n\t\t" + providerLine + appContent[insertPos:]

	// Write back to app.go
	err = os.WriteFile(appFile, []byte(appContent), 0644)
	if err != nil {
		return fmt.Errorf("failed to write to app.go: %w", err)
	}

	fmt.Println("app.go updated successfully.")
	return nil
}

func checkModuleExistence(moduleBase string) bool {
	// Check if the base module directory exists
	if _, err := os.Stat(moduleBase); err == nil {
		return true // Module directory already exists
	}
	return false
}
