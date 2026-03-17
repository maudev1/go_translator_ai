# Tradutor de Arquivos Chave=Valor

Ferramenta feita em **Golang** para traduzir arquivos no formato `chave=valor`, muito comum em arquivos de **localização e configuração** de projetos.

O sistema lê um arquivo, traduz os valores e gera uma nova versão traduzida mantendo as mesmas chaves.

⚠️ **Observação:** o projeto ainda está **em desenvolvimento**, então algumas funcionalidades podem não estar totalmente implementadas ou podem sofrer mudanças.

## Métodos de Tradução

A ferramenta foi planejada para suportar três formas de tradução:

* **IA via Groq usando Llama 3**
* **API do Google Tradutor**
* **Tradução manual**, caso você prefira inserir os textos traduzidos sem usar serviços externos.

## Exemplo

Arquivo original:

```
hello=Hello
bye=Goodbye
thanks=Thank you
```

Arquivo traduzido:

```
hello=Olá
bye=Adeus
thanks=Obrigado
```

## Instalação

### 1. Instalar Go

Baixe e instale o Go:
[https://go.dev/dl/](https://go.dev/dl/)

Verifique se está instalado:

```
go version
```

### 2. Clonar o repositório

### 3. Instalar dependências

```
go mod tidy
```

### 4. Configurar APIs (opcional)

Se quiser usar tradução automática, configure suas chaves de API:

* **Groq API (Llama 3)**
* **Google Translate API**

Adicione as chaves nas variáveis de ambiente ou na configuração do projeto.

### 5. Executar

```
go run main.go
```

O programa irá ler o arquivo `chave=valor`, traduzir os textos e gerar a versão traduzida automaticamente ou permitir tradução manual.
