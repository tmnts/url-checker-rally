# --- ЭТАП 1: Сборка Go-бинарника ---
# --- ETAPA 1: Compilação do binário em Go ---
FROM golang:1.26-alpine AS go_builder
WORKDIR /app
COPY go_service/go.mod go_service/main.go ./
# Компилируем оптимизированный бинарник без лишнего трешака
# Compila um binário otimizado, removendo símbolos de depuração desnecessários
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o url_checker_go main.go

# --- ЭТАП 2: Финальный контейнер для ралли ---
# --- ETAPA 2: Container final para a corrida de performance ---
FROM python:3.11-alpine
WORKDIR /app

# Здесь мы устанавливаем зависимости Python (к счастью, она там всего одна)
# Instala dependências do sistema e prepara o ambiente para o Python
COPY python_service/requirements.txt ./
RUN pip install --no-cache-dir -r requirements.txt

# Копируем Python-скрипт
# Copia o script Python e o binário Go já compilado na Etapa 1
COPY python_service/main.py ./url_checker_py.py

# Копируем уже скомпилированный, готовый Go-бинарник из первого этапа
# Copiamos o binário Go já compilado e pronto para uso do primeiro estágio.
COPY --from=go_builder /app/url_checker_go ./url_checker_go

# Копируем bash-скрипт для запуска ралли (опционально) или просто держим контейнер готовым
# Executa os dois serviços em sequência para comparar a velocidade
CMD ["sh", "-c", "echo '=== ЗАПУСК PYTHON - INICIANDO PYTHON ===' && python url_checker_py.py && echo '\n=== ЗАПУСК GO - INICIANDO GO ===' && ./url_checker_go"]
