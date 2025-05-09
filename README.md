# 🔍 Certstream-Domains

This tool opens a websocket connection to [Certstream Server Go](https://github.com/d-Rickyy-b/certstream-server-go) in order to extract domains from the stream of certificate transparency logs and saves to jsonl in the domain_data directory.

---

## 📋 Features

- Real-time aggregating of Domains from Certificate Transparency logs
- Results saved as structured JSONL file

## 🔧 Requirements

- Docker - That's it! All other dependencies are containerized

## 🐳 Installation

```bash
# Clone the repository
git clone https://github.com/derekrgreene/certstream-domains.git
cd certstream-domains

# Build and start Certstream-domains
docker-compose up --build
```

## 📂 Output Format

Results are saved as JSONL in the `domain_data` directory with the following format:

```json
"*.leakednews20.pages.dev"
"leakednews20.pages.dev"
"crisonchemistry.com"
"www.crisonchemistry.com"
"media1.evmlabs.com"
"media1.tenxerlabs.com"
"cashless.intectrafic.com"
"b2.francillon.net"
"b3.francillon.net"
"francillon.net"
"vw.francillon.net"
"redbussecondarydns.picturegamer.com"
"cctamo.net"
"*.cctamo.net"
"vurkxwww.singed.internetdruckerei.com"
"*.elitedentalmy.com"
"elitedentalmy.com"
"3a5h6.com"
```

⚠️ Many of the domains collected will be prepended with an '*' (wildcard) as domains are collected from Certificate Transparency Logs 

## 📝 License

[MIT License](LICENSE)

## 📧 Contact

For support or questions, please open an issue on GitHub.