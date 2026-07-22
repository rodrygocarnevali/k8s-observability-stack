# Kubernetes Observability Stack

## 📑 Índice

...

## 📖 Visão Geral

...

## 🎯 Objetivos

...

## 🏗️ Arquitetura

...

## 🛠️ Tecnologias

...

## 📂 Estrutura

...

## 🚀 Implantação

...

## ✅ Validação

...

## 📊 Dashboards

...

## 🚨 Alertas

...

## 🤖 Automação

...

## 📸 Evidências

...

## 📚 Aprendizados

````markdown
## 🔧 Troubleshooting

### Problema de validação SSL durante instalação do Helm

Durante a instalação do Helm através do repositório APT, foi identificado um problema de validação do certificado SSL ao acessar o repositório utilizado para instalação.

Comando executado:

```bash
curl https://baltocdn.com/helm/signing.asc
````

Erro apresentado:

```text
curl: (60) SSL certificate problem: unable to get local issuer certificate
```

---

### Diagnóstico

Inicialmente foi validado se o problema estava relacionado à conectividade da máquina.

Teste de acesso HTTPS:

```bash
curl -I https://google.com
```

Resultado:

```text
HTTP/2 301
```

O teste confirmou que:

* A conectividade de rede estava funcionando.
* A resolução DNS estava operacional.
* O acesso HTTPS para outros serviços estava funcionando.

---

### Validação dos certificados

Foi analisada a cadeia de certificados apresentada pelo domínio:

```bash
echo | openssl s_client -connect baltocdn.com:443 -servername baltocdn.com 2>/dev/null | openssl x509 -noout -issuer -subject -dates
```

Resultado:

```text
issuer=C = US, O = Let's Encrypt, CN = YR1
subject=CN = baltocdn.com
notBefore=Jul 18 17:32:27 2026 GMT
notAfter=Oct 16 17:32:26 2026 GMT
```

A análise indicou um problema relacionado à validação da cadeia de certificados do domínio utilizado pelo repositório.

---

### Solução aplicada

Como alternativa, a instalação do Helm foi realizada utilizando o instalador oficial disponibilizado pelo projeto Helm.

Comando utilizado:

```bash
curl https://raw.githubusercontent.com/helm/helm/main/scripts/get-helm-3 | bash
```

Após a instalação, foi realizada a validação:

```bash
helm version
```

Resultado esperado:

```text
version.BuildInfo{Version:"v3.x.x"...}
```

---

### Aprendizado

Durante a instalação de ferramentas em ambientes Linux, é importante validar as camadas da infraestrutura antes de alterar componentes da aplicação.

O processo de troubleshooting seguiu a seguinte sequência:

1. Verificação da conectividade de rede.
2. Validação da resolução DNS.
3. Verificação de data e hora do sistema.
4. Análise dos certificados SSL/TLS.
5. Definição de uma estratégia alternativa de instalação.

Esse processo evitou alterações desnecessárias no ambiente e permitiu identificar corretamente a origem do problema.

```
```
