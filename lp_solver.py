import sys
from scipy.optimize import linprog

# Чтение данных из аргументов командной строки
costs = [float(sys.argv[1]), float(sys.argv[4]), float(sys.argv[7])]
times = [float(sys.argv[2]), float(sys.argv[5]), float(sys.argv[8])]
distances = [float(sys.argv[3]), float(sys.argv[6]), float(sys.argv[9])]

# Ограничения
max_time = float(sys.argv[10])
max_distance = float(sys.argv[11])

# Целевая функция: минимизация стоимости
c = costs

# Матрица ограничений (каждое ограничение — строка)
A = [
    times,       # Ограничение на время
    distances    # Ограничение на расстояние
]

# Правые части ограничений
b = [max_time, max_distance]

# Ограничение на сумму долей маршрутов
A_eq = [[1, 1, 1]]
b_eq = [1]

# Ограничения на неотрицательность
bounds = [(0, 1), (0, 1), (0, 1)]

# Решение задачи
result = linprog(c, A_ub=A, b_ub=b, A_eq=A_eq, b_eq=b_eq, bounds=bounds, method='highs')

if result.success:
    print("Оптимальное распределение по маршрутам:", result.x)
    print("Минимальная стоимость:", result.fun)
else:
    print("Решение не найдено:", result.message)
