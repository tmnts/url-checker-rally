import csv
import time
import requests

# Список сайтов для проверки (в реальности тут могут быть тысячи URL)
# Lista de sites para verificar (na realidade, poderia haver milhares de URLs aqui)
urls = [
    "https://google.com", "https://github.com", "https://microsoft.com",
    "https://ass-we-can.com", "https://out.com", "https://python.org"
] * 10  # Увеличим список для наглядности тестов / Vamos expandir a lista para tornar os testes mais claros.

def check_url(url):
    try:
        start_time = time.time()
        response = requests.get(url, timeout=5)
        duration = round((time.time() - start_time) * 1000) # в миллисекундах / em milissegundos
        return {"url": url, "status": response.status_code, "latency_ms": duration, "error": ""}
    except Exception as e:
        return {"url": url, "status": 0, "latency_ms": 0, "error": str(e)}

def main():
    print("Python: Начало обработки... / Processamento iniciado...")
    start = time.time()
    
    results = []
    for url in urls:
        res = check_url(url)
        results.append(res)
        
    # Сохраняем в CSV
    with open("report_py.csv", "w", newline="", encoding="utf-8") as f:
        writer = csv.DictWriter(f, fieldnames=["url", "status", "latency_ms", "error"])
        writer.writeheader()
        writer.writerows(results)
        
    print(f"Python: Готово за {time.time() - start:.2f} сек. Результаты в report_py.csv / Python: Concluído em {time.time() - start:.2f} s. Resultados em report_py.csv")

if __name__ == "__main__":
    main()