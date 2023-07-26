package main

import (
	"os"
	"text/template"
)

func main() {
	t1 := template.New("t1")                // Defines a Template and it's name.
	t1, err := t1.Parse("Value is {{.}}\n") // Defines the format of the template '{{.}}' indicates the position of a value.
	if err != nil {
		panic(err)
	}

	t1 = template.Must(t1.Parse("Value: {{.}}\n")) // Defines format of templates, check for "Parse" errors
	t1.Execute(os.Stdout, "some text")             // Applying template.
	t1.Execute(os.Stdout, 5)                       // Applying template.
	t1.Execute(os.Stdout, []string{
		"Go",
		"Rust",
		"C++",
		"C#",
	}) // Applying template.

	// Defines a fatoory of templates.
	Create := func(name, t string) *template.Template {
		return template.Must(template.New(name).Parse(t))
	}
	t2 := Create("t2", "Name: {{.Name}}\n") // Defines a template with spacific value vacancy (applyed to structs and maps).
	t2.Execute(os.Stdout, struct {
		Name string
	}{"Jane Doe"})
	t2.Execute(os.Stdout, map[string]string{
		"Name": "Mikey Mouse",
	})

	//Defines a template with conditions (if value is not a 'zero value' present prints 'yes' otherwhise prints 'no')
	t3 := Create("t3",
		"{{if . -}} yes {{else -}} no {{end}}\n") // NOTE: character '-' trims white spaces.
	t3.Execute(os.Stdout, "not empty")
	t3.Execute(os.Stdout, "")

	//Defines a template with can loops trough slices, arrays, maps of channels ('{{range .}}' block).
	t4 := Create("t4",
		"Range {{range .}}{{.}} {{end}}\n")
	t4.Execute(os.Stdout,
		[]string{
			"Go",
			"Rust",
			"C++",
			"C#",
		})
}
