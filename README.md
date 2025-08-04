Auth Service + Monitoring
Микросервис аутентификации для форума по поиску разработчиков в команду.
Реализует регистрацию и авторизацию пользователей с JWT, а также мониторинг с помощью Prometheus, Grafana и Node
Exporter.

## 🚀 Запуск проекта

### 1. Клонируй репозиторий
   ```bash
   git clone https://github.com/<твой-репозиторий>.git
   cd <папка-проекта>
   ```
### 2. Настрой .env
   Создай файл .env на основе .env.example:

```bash
cp .env.example .env
```

Заполни переменные своими значениями.

.env.example:

```env
# Database settings

DB_HOST=DB_CONTAINER_NAME
DB_USER=DB_USER
DB_PASSWORD=DB_PASSWORD
DB_NAME=DB_NAME
DB_PORT=DB_PORT

# Auth-service settings

AUTH_PORT=PORT_AUTH_SERVICE
JWT_SECRET=JWT_SECRET
JWT_EXPIRES_IN=JWT_EXPIRES_IN_TIME
```

### 3. Запусти сервисы
   ```bash
   docker-compose --env-file .env -f docker-compose.yml -f docker-compose.monitoring.yml up -d --build
   ```
### 4. Доступ к сервисам
   * Auth-service API: **http://localhost:${AUTH_PORT}**
   * Prometheus: **http://localhost:9090**
   * Grafana: **http://localhost:3000 (логин: admin, пароль: admin)**
   * Node Exporter: **http://localhost:9100/metrics**

## 📡 API эндпоинты
POST **/auth/register**. Регистрация нового пользователя.

Пример запроса:

```json
{
"email": "user@example.com",
"password": "123456"
}
```
Пример ответа:

```json
{
"status": "success",
"data": {
  "id": 1,
  "email": "user@example.com",
  "access_token": "jwt_token",
  "token_type": "Bearer",
  "expires_in": 900
  }
}
```
POST **/auth/login**. Авторизация пользователя.

Пример запроса:
```json
{
"email": "user@example.com",
"password": "123456"
}
```
Пример ответа:

```json
{
"status": "success",
"data": {
  "access_token": "jwt_token",
  "token_type": "Bearer",
  "expires_in": 900
  }
}
```
GET **/metrics**. Метрики Prometheus.

GET **/health**. Проверка состояния сервиса.

## 📊 Мониторинг
В проект интегрированы:

* **Prometheus** — сбор метрик
* **Grafana** — визуализация данных (предустановлены 2 дашборда)
* **Node Exporter** — метрики хоста

После запуска в Grafana уже будут доступны дашборды для:

* HTTP-запросов сервиса
* Нагрузка машины

## 🛠 Технологии
* Go (Gin, GORM, bcrypt, JWT)
* PostgreSQL
* Prometheus
* Grafana
* Node Exporter
* Docker, Docker Compose