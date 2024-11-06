# fullcycle-multithreading-cep-api

Um serviço em Go para buscar o endereço de um CEP, consultando duas APIs diferentes e retornando a resposta mais rápida. O serviço faz requisições concorrentes para as APIs [BrasilAPI](https://brasilapi.com.br) e [ViaCEP](https://viacep.com.br/), e exibe o endereço obtido da primeira API a responder dentro de um limite de 1 segundo.

## Funcionalidades

- Realiza duas requisições concorrentes para buscar o endereço de um CEP.
- Retorna a resposta da API que responder mais rapidamente.
- Limita o tempo de resposta a 1 segundo; caso contrário, retorna um erro de timeout.

## Requisitos

- Go 1.15 ou superior
- Conexão com a internet para acessar as APIs

## Como rodar o serviço

1. Clone o repositório:
   ```bash
   git clone https://github.com/marcosocram/fullcycle-multithreading-cep-api.git
   cd fullcycle-multithreading-cep-api
    ```

2. Rodar o serviço:
   ```bash
   go run main.go
   ```
   O programa irá buscar o endereço do CEP especificado no código (CEP 88110798 por padrão) e exibirá a resposta da API que respondeu mais rapidamente.


3. Alterar o CEP:
   * Caso queira testar com outro CEP, basta alterar a variável cep na função `main` para o valor desejado. Exemplo:
   ```go
   cep := "88110798" // Altere o CEP aqui
   ```

## Exemplo de Saída

Aqui estão exemplos do que serão exibidos no terminal:

**Cenário 1**: A API BrasilAPI responde mais rápido

```bash
Resposta mais rápida da API BrasilAPI:
CEP: 88110798
Logradouro: Avenida Osvaldo José do Amaral
Bairro: Bela Vista
Localidade: São José
UF: SC
```

**Cenário 2**: A API ViaCEP responde mais rápido

```bash
Resposta mais rápida da API ViaCEP:
CEP: 88110-798
Logradouro: Avenida Osvaldo José do Amaral
Bairro: Bela Vista
Localidade: São José
UF: SC
```

**Cenário 3**: Nenhuma das APIs responde dentro de 1 segundo

```bash
Erro: timeout após 1 segundo.
```

