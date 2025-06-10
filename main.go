package main

import (
    "encoding/json"
    "html/template"
    "os"
    "strings"

    "github.com/SebastiaanKlippert/go-wkhtmltopdf"
)

func main() {
    // Caminho fixo para o arquivo JSON
    jsonFilePath := "data.json"

    // Lê o arquivo JSON
    jsonFile, err := os.Open(jsonFilePath)
    if err != nil {
        panic(err)
    }
    defer jsonFile.Close()

    // Decodifica o JSON em um map
    var data map[string]interface{}
    decoder := json.NewDecoder(jsonFile)
    err = decoder.Decode(&data)
    if err != nil {
        panic(err)
    }

    // Gera o HTML a partir do template
    tmpl, err := template.ParseFiles("template.html")
    if err != nil {
        panic(err)
    }

    htmlFile, err := os.Create("output.html")
    if err != nil {
        panic(err)
    }
    defer htmlFile.Close()

    err = tmpl.Execute(htmlFile, data)
    if err != nil {
        panic(err)
    }

    // Adiciona debug: imprime o conteúdo HTML no console
    htmlFile.Seek(0, 0) // Volta para o início do arquivo
    content, _ := os.ReadFile("output.html")
    // println(string(content)) // Imprime o conteúdo no console

    // Gera o PDF a partir do HTML
    pdfg, err := wkhtmltopdf.NewPDFGenerator()
    if err != nil {
        panic(err)
    }

    // Cria uma nova página com o conteúdo HTML
    pdfg.AddPage(wkhtmltopdf.NewPageReader(strings.NewReader(string(content))))
    err = pdfg.Create()
    if err != nil {
        panic(err)
    }

    // Salva o PDF como output.pdf
    err = pdfg.WriteFile("output.pdf")
    if err != nil {
        panic(err)
    }

    println("PDF gerado com sucesso: output.pdf")
}
