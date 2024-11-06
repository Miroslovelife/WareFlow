# Используем официальный образ OSRM
FROM osrm/osrm-backend:latest

# Устанавливаем рабочую директорию
WORKDIR /data

# Копируем файл OSM в контейнер
# Предполагается, что файл russia-latest.osm.pbf находится в той же директории, что и Dockerfile
COPY russia-latest.osm.pbf /data/russia-latest.osm.pbf

# Экстракция данных OSM
RUN osrm-extract -p /opt/car.lua /data/russia-latest.osm.pbf

# Разделение данных
RUN osrm-partition /data/russia-latest.osrm

# Настройка весов
RUN osrm-customize /data/russia-latest.osrm

# Открываем порт 5000
EXPOSE 5000

# Запускаем OSRM сервер
CMD ["osrm-routed", "--algorithm", "mld", "/data/russia-latest.osrm"]