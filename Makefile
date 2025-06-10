.PHONY: all clean run

# Nome do projeto
PROJECT_NAME := template-prototyping

# Verifica o sistema operacional
OS := $(shell uname)

# Verifica se o Go está instalado
GO_CHECK := $(shell command -v go 2>/dev/null)

all: clean run

install-deps:
	@if [ -z "$(GO_CHECK)" ]; then \
		echo "Go não encontrado. Instalando..."; \
		if [ "$(OS)" = "Linux" ]; then \
			sudo apt-get update && sudo apt-get install -y golang; \
		elif [ "$(OS)" = "Darwin" ]; then \
			brew install go; \
		elif [ "$(OS)" = "Windows" ]; then \
			echo "Por favor, instale o Go manualmente no Windows."; \
		else \
			echo "Sistema operacional não suportado para instalação do Go!"; \
			exit 1; \
		fi \
	else \
		echo "Go já está instalado."; \
	fi; \
	\
	if [ "$(OS)" = "Linux" ]; then \
		echo "Instalando wkhtmltopdf para Linux..."; \
		sudo apt-get install -y wkhtmltopdf; \
	elif [ "$(OS)" = "Darwin" ]; then \
		echo "Instalando wkhtmltopdf para macOS..."; \
		brew install wkhtmltopdf; \
	elif [ "$(OS)" = "Windows" ]; then \
		echo "Por favor, instale o wkhtmltopdf manualmente no Windows."; \
	else \
		echo "Sistema operacional não suportado para instalação do wkhtmltopdf!"; \
		exit 1; \
	fi

run:
	go run main.go

clean:
	rm -f output.html output.pdf
