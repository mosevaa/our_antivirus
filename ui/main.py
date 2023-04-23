import sys

# записываем данные в stdout
sys.stdout.write("Hello, World!\n")
sys.stdout.flush()  # очищаем буфер

# читаем данные из stdin
data = sys.stdin.read()
print("Input:", data)