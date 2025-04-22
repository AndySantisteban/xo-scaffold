package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/joho/godotenv"
)

func goDotEnvVariable(key string) string {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

func main() {
	godotenv.Load()
	var output string
	var views string

	flag.StringVar(&output, "out", "gen", "Directorio de salida")
	flag.StringVar(&views, "views", "", "Nombres de vistas separados por coma (ej: VwUsuarios,VwClientes)")

	flag.Parse()

	if views == "" {
		log.Fatal("Debe proporcionar al menos una vista usando --views")
	}

	viewList := strings.Split(views, ",")
	schema := goDotEnvVariable("DATABASE_URL")
	var path = filepath.Join(".", "gen")
	args := []string{
		"schema", schema,
		"--fk-mode", "smart",
		// "--single", "output.go",
		"--out", path,
		"--use-index-names",
		"--template", "go",
		// "--go-append",
		// "--go-not-first",
		// "--verbose",
	}
	for _, v := range viewList {
		args = append(args, "--include", v)
	}
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err := os.Mkdir(path, os.ModePerm)
		if err != nil {
			fmt.Println(err.Error())
		}
	}

	cmd := exec.Command("./xo.exe", args...)
	cmd.Stdout = log.Writer()
	cmd.Stderr = log.Writer()

	if err := cmd.Run(); err != nil {
		log.Fatalf("Fallo al ejecutar Scaffold: %v", err)
	}

	fmt.Println("¡Structs generados con éxito!")
}
