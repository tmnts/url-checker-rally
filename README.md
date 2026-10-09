# Concurrency Rally: Go vs Python (Transatlantic I/O Battle) 🏁🇧🇷

[English] 
A production-ready benchmark demonstrating architectural migration of an I/O-bound microservice from Python (synchronous `requests`) to Go (concurrent `goroutines` and `channels`). 

[Português]
Um benchmark pronto para produção que demonstra a migração arquitetural de um microsserviço focado em I/O de Python (`requests` síncronos) para Go (`goroutines` e `channels` concorrentes).

---

## 🏎️ The Race Track & Overseas Latency / A Pista de Corrida & Latência Transatlântica

[English]
Unlike synthetic tests running on local networks, this rally forces both engines to fetch data from actual live servers across the globe — including South American music labels (e.g., *Love Funk*), localized platforms (*Kit de Pontos*), and specific overseas media portals. 
This puts the focus heavily on handling:
* **High Geo-Ping:** Packets physically traveling across transatlantic submarine cables.
* **Anti-Bot Protections:** Surviving strict Cloudflare rate-limiting and connection throttling without native browser headers.

[Português]
Diferente de testes sintéticos em redes locais, esta corrida força ambos os motores a buscar dados de servidores reais ao redor do mundo — incluindo gravadoras sul-americanas (*Love Funk*), plataformas locais (*Kit de Pontos*) e portais de mídia estrangeiros.
Isso coloca o foco principal no gerenciamento de:
* **Alto Geo-Ping:** Pacotes viajando fisicamente por cabos submarinos transatlânticos.
* **Proteções Anti-Bot:** Sobrevivendo a rate-limiting rígidos do Cloudflare e bloqueios de conexão sem headers nativos de navegador.

---

## 📊 Benchmarks & Reality Check / Benchmarks & Resultados Reais

### The Battle for Concurrency
* **Python Loop:** ~99.88 seconds (Linear blocking. The script falls asleep waiting for overseas handshakes sequentially).
* **Go Engine:** ~1.476275668 seconds (True concurrency. Spawns asynchronous light threads to process all requests simultaneously, ignoring geographic distance).

### A Batalha de Concorrência
* **Loop Python:** ~99.88 segundos (Bloqueio linear. O script "dorme" esperando os handshakes internacionais sequencialmente).
* **Motor Go:** ~1.476275668 segundos (Concorrência real. Cria threads leves assíncronas para processar todas as requisições simultaneamente, ignorando a distância geográfica).

> **Performance Gain / Ganho de Performance:** Go outpaced Python by **~17x** under real-world, high-latency network load! / O Go superou o Python em cerca de **17 vezes** sob carga de rede real de alta latência!

---

## 🛠️ Tech Stack & CI/CD / Tecnologias & CI/CD

* **Go 1.26** (Goroutines, Channels, `sync.WaitGroup`, `httptest` for mocking).
* **Python 3.11** (Data export & comparison logic).
* **Docker Multi-Stage:** Optimized build splitting compilers from final lightweight Alpine images.
* **GitHub Actions / GitVerse:** Automated testing workflow executed on every code push (`go test -v`).

---

## 🚀 How to Run the Rally / Como Executar a Corrida

[English]
You don't need to install Go or Python on your local machine. The entire environment, compilation, and execution are fully containerized using a Multi-stage Docker setup.

[Português]
Você não precisa instalar o Go ou o Python na sua máquina local. Todo o ambiente, compilação e execução são totalmente conteinerizados usando uma estrutura Docker Multi-stage.

### 🏁 Execution Commands / Comandos de Execução

1. **Clone the repository / Clone o repositório:**
   ```bash
   git clone [https://github.com/tmnts/url-checker-rally]
   cd url-checker-rally
   ```

2. **Build the single Multi-stage Image / Construa a imagem Multi-stage única:**
   ```bash
   docker build -t url-checker-rally .
   ```

3. **Run the Rally and extract CSV logs / Execute a corrida e extraia os relatórios CSV:**
   ```bash
   docker run --rm -v $(pwd):/app/out url-checker-rally
   ```

[English]
*Note: The `--rm` flag will automatically clean up the container after execution, and the `-v` volume flag will securely output `report_py.csv` and `report_go.csv` directly into your project root directory.*

[Português]
*Nota: A flag `--rm` limpará automaticamente o container após a execução, e a flag de volume `-v` extrairá com segurança os arquivos `report_py.csv` e `report_go.csv` diretamente para o diretório raiz do seu projeto.*


<span style="color: #00FF00; font-family: monospace;">$ TMNTS A BRABA: Com amor...</span>
