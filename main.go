package main

import (
	"bufio"
	"fmt"
	"os"
)

// Palabras Reservadas
var palabras_reservadas = []string{
	"auto", "const", "double", "float", "int",
	"short", "struct", "unsigned", "break",
	"continue", "else", "for", "long", "signed",
	"switch", "void", "case", "default", "enum",
	"goto", "register", "sizeof", "typedef", "volatile",
	"char", "do", "extern", "if", "return", "static",
	"union", "while", "asm", "dynamic_cast", "namespace",
	"reinterpret_cast", "try", "bool", "explicit", "new",
	"static_cast", "typeid", "catch", "false", "operator",
	"template", "typename", "class", "friend", "private", "this",
	"using", "const_cast", "inline", "public", "throw", "virtual",
	"delete", "mutable", "protected", "true", "wchar_t",
}

//Definimos una lista con todos los operadores
var operadores = []string{
	"=", "+", "+=", "-", "-=", "*", "*=",
	"/", "/=", "%", "%=", "++", "--",
	"<", ">", "<=", ">=", "==", "!=", "<=>",
	"!", "&&", "||", "<<=", ">>=", "~", "&=",
	"|", "|=", "^", "^=", "(", ")", "[", "]", "{",
	"}", "->", ".", "->.", ".*", ",", "::", "::*",
	"<?", ">?", ";",
}

// Alphabetical Values
var alphabet = []string{
	"A", "B", "C", "D", "E",
	"F", "G", "H", "I", "J",
	"K", "L", "M", "N", "O",
	"P", "Q", "R", "S", "T",
	"U", "V", "W", "X", "Y", "Z",
	"a", "b", "c", "d", "e",
	"f", "g", "h", "i", "j",
	"k", "l", "m", "n", "o",
	"p", "q", "r", "s", "t",
	"u", "v", "w", "x", "y", "z",
}

// Numbers
var numbuhs = []string{
	"1", "2", "3", "4", "5",
	"6", "7", "8", "9", "0",
}

var initialString string = `
<!DOCTYPE html>
    <html>
    <head>
        <link href="https://allfont.es/allfont.css?fonts=lucida-console" rel="stylesheet" type="text/css" />
        <meta charset="UTF-8">
        <title> Analizador de Sintaxis </title>
        <link rel="stylesheet" href="formato.css">
    </head>
    <body>
	`
var finalStirng string = `
	</body>
	</html>
`

func stringInSlice(element string, slice []string) bool {
	for _, word := range slice {
		if element == word {
			return true
		}
	}
	return false
}

func iterativeLexer(sourceFile string, resultFile string) {
	file, err := os.Open(sourceFile)
	resFile, err2 := os.Create(resultFile)

	defer file.Close()
	defer resFile.Close()

	if err != nil || err2 != nil {
		os.Exit(1)
	}

	resFile.WriteString(initialString)

}

func main() {

	file, err := os.Open("./test.txt")
	resultFile, err2 := os.Create("./highlight.html")

	defer file.Close()
	defer resultFile.Close()

	if err != nil || err2 != nil {
		os.Exit(1)
	}

	resultFile.WriteString(initialString)
	m_scaner := bufio.NewScanner(file)

	var linea string
	var element string
	var counter int

	for m_scaner.Scan() {
		counter++
		linea = "<p>"
		tam := len(m_scaner.Text())
		currentLine := m_scaner.Text()
		// fmt.Printf("tam: %v, linea actual: %v\n", tam, currentLine)
		i := 0
		var j int

		for i <= tam {
			// fmt.Printf("%v\n", i)
			if i == tam {
				fmt.Printf("%v\n", counter)
				linea += "</p>\n"
				fmt.Printf("Linea Agregada: %v\n", linea)
				resultFile.WriteString(linea)
				break

				// Encuentra # puede ser un include
			} else if string(currentLine[i]) == "#" {
				element += string(currentLine[i])
				j = i + 1

				// Mientras J sea menor que la linea
				for j < tam {
					//Si el valor es alfabetico lo agrega a element y aumenta j
					if stringInSlice(string(currentLine[i]), alphabet) {
						element += string(currentLine[i])
						j++

					} else {
						// Si es un include lo agrega como palabra reservada a la linea
						if element == "#include" {
							element = "<span class='palabra_reservada'>" + element + "</span>"
							linea += element
							element = ""
							break
							// si no es un #include lo trata como operador
						} else {
							element = "<span class='operador'>" + element + "</span>"
							linea += element
							element = ""
							break

						}
					}
				}
				i = j
				// Si encuentra un espacio
			} else if string(currentLine[i]) == " " {
				j = 0

				for string(currentLine[i]) == " " {
					// si el siguiente elemente no es un epsacio agrega el espacio a la linea
					if string(currentLine[i+1]) != " " {
						linea += string(currentLine[i])
						i++
					} else {
						i++
						j++

						// si j llega a tres significa que encontro una tabulacion
						if j == 3 {
							linea += "<span class='tab'></span>"
							break
						}
					}
				}
				// Si encuentra una letra
			} else if stringInSlice(string(currentLine[i]), alphabet) {
				element += string(currentLine[i])
				j = i + 1

				for j < tam {
					// si j es una letra o un numero la agrega y aumenta j
					if stringInSlice(string(currentLine[j]), alphabet) || stringInSlice(string(currentLine[j]), numbuhs) {
						element += string(currentLine[j])
						j++
						// si no lo es
					} else {
						// si element esta en las palabras reservadas, la clasifica como una y la agrega a la linea
						if stringInSlice(element, palabras_reservadas) {
							element = "<span class='palabra_reservada'>" + element + "</span>"
							linea += element
							element = ""
							break
							// si no es una palabra reservada lo trata como variable y la agrega

						} else {
							element = "<span class='variable'>" + element + "</span>"
							linea += element
							element = ""
							break
						}

					}
				}
				i = j
				// si encuentra un numero
			} else if stringInSlice(string(currentLine[i]), numbuhs) {
				element += string(currentLine[i])
				j = i + 1

				for j < tam {
					// Si es un numero, punto o una e o agrega
					if stringInSlice(string(currentLine[j]), alphabet) || string(currentLine[j]) == "e" || string(currentLine[j]) == "." {
						element += string(currentLine[j])
						j++
						// si es una letra lee todo mientras sea numero o letra
					} else if stringInSlice(string(currentLine[j]), alphabet) {
						for stringInSlice(string(currentLine[j]), alphabet) || stringInSlice(string(currentLine[j]), numbuhs) {
							element += string(currentLine[j])
							j++
						}
						// si es diferente a letra o numero clasifica la variable como error y lo agrega a la linea
						element = "<span  class='error'>" + element + "</span>"
						linea += element
						element = ""
						break
						// si no es, lo clasifica com un numero y lo agrega
					} else {
						element = "<span  class='numero'>" + element + "</span>"
						linea += element
						element = ""
						break

					}
				}
				i = j
				//si encuentra un operador
			} else if stringInSlice(string(currentLine[i]), operadores) {
				// revisa si es un simbolo de division
				if string(currentLine[i]) == "/" {
					// si hay otro simbolo de division mas adelante significa que es un comentario
					// lo agrega a la linea e iguala i al tamaño
					if string(currentLine[i+1]) == "/" {
						element = "<span class='comentario'>" + string(currentLine[i:]) + "</span>"
						linea = linea + element
						element = ""
						i = tam
						// si no, lo trata como un simbolo de division y lo agrega
					} else {
						element = "<span class='operador'>" + string(currentLine[i]) + "</span>"
						linea += element
						element = ""
						i++
					}
					// si no es un simbolo de division lo trata como un operador
				} else {
					element = "<span class='operador'>" + string(currentLine[i]) + "</span>"
					linea += element
					element = ""
					i++

				}
				// Si encuentra comillas
			} else if string(currentLine[i]) == "\"" {
				element += string(currentLine[i])
				j = i + 1
				for j < tam {
					// si encuentra la comilla de cierra, agrega var como un string
					if string(currentLine[j]) == "\"" {
						element = "<span class='string'>" + element + "</span>"
						j++
						linea += element
						element = ""
						break
					} else {
						element += string(currentLine[j])
						j++
					}
				}
				i = j
			} else {
				i++
			}

		}
	}
	resultFile.WriteString(finalStirng)
	resultFile.Sync()

}
