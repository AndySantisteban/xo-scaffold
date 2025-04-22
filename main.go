package main

import (
	"flag"
	"fmt"
	"log"
	"os/exec"
	"strings"
)

func main() {
	// Parámetros por consola
	var output string
	var pkg string
	var views string

	flag.StringVar(&output, "out", "gen", "Directorio de salida")
	flag.StringVar(&pkg, "pkg", "gen", "Nombre del paquete")
	flag.StringVar(&views, "views", "", "Nombres de vistas separados por coma (ej: VwUsuarios,VwClientes)")

	flag.Parse()

	if views == "" {
		log.Fatal("Debe proporcionar al menos una vista usando --views")
	}

	viewList := strings.Split(views, ",")

	args := []string{
		"schema", "sqlserver://sysugosrv02:UGO%21dev%210823@develop86.devfci.com/CenturionNotes",
		"--fk-mode", "parent",
		// "--single", "output.go",
		"--out", output,
		"--use-index-names",
		"--go-not-first",
		// "--verbose",
	}
	for _, v := range viewList {
		args = append(args, "--include", v)
	}

	cmd := exec.Command("./xo.exe", args...)
	cmd.Stdout = log.Writer()
	cmd.Stderr = log.Writer()

	if err := cmd.Run(); err != nil {
		log.Fatalf("Fallo al ejecutar Scaffold: %v", err)
	}

	fmt.Println("¡Structs generados con éxito!")
}
