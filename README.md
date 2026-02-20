# 🚗 FleetPulse - Real-Time Fleet Management System

<p align="center">
  <img src="https://img.shields.io/badge/Spring%20Boot-3.x-brightgreen?style=for-the-badge&logo=springboot" />
  <img src="https://img.shields.io/badge/Go-1.25-00ADD8?style=for-the-badge&logo=go" />
  <img src="https://img.shields.io/badge/Docker-Compose-2496ED?style=for-the-badge&logo=docker" />
  <img src="https://img.shields.io/badge/Keycloak-OAuth2-orange?style=for-the-badge&logo=keycloak" />
</p>

araç filolarını gerçek zamanlı olarak izlemek, hız ihlallerini tespit etmek ve uyarı oluşturmak için tasarlanmış bir filo yönetim sistemidir.

---

## 📋 İçindekiler

- [Mimari](#-mimari)
- [Servisler](#-servisler)
- [Teknolojiler](#-teknolojiler)
- [Kurulum](#-kurulum)
- [API Endpoints](#-api-endpoints)
- [Konfigürasyon](#-konfigürasyon)
- [Geliştirme](#-geliştirme)

---

## 🏗 Mimari

```
                                    ┌─────────────────┐
                                    │   Keycloak      │
                                    │   (Auth)        │
                                    │   :9090         │
                                    └────────┬────────┘
                                             │
                                             ▼
┌─────────────┐    ┌─────────────────────────────────────────────────────┐
│   Client    │───▶│              API Gateway (:8060)                    │
└─────────────┘    └───────┬─────────────────┬─────────────────┬─────────┘
                           │                 │                 │
                           ▼                 ▼                 ▼
                   ┌───────────────┐ ┌───────────────┐ ┌───────────────┐
                   │  Ingestion    │ │    Fleet      │ │    Alert      │
                   │  Service (Go) │ │  Management   │ │   Service     │
                   │    :8081      │ │    :8083      │ │    :8082      │
                   └───────┬───────┘ └───────┬───────┘ └───────┬───────┘
                           │                 │                 │
              ┌────────────┼─────────────────┼─────────────────┤
              │            │                 │                 │
              ▼            ▼                 ▼                 ▼
        ┌──────────┐ ┌──────────┐    ┌────────────┐    ┌────────────┐
        │  Redis   │ │ RabbitMQ │    │ PostgreSQL │    │ PostgreSQL │
        │  :6379   │ │  :5672   │    │   :5435    │    │   :5434    │
        └──────────┘ └──────────┘    └────────────┘    └────────────┘

                    ┌─────────────────────────────────────┐
                    │         Discovery Server            │
                    │        (Eureka) :8761               │
                    └─────────────────────────────────────┘
                    ┌─────────────────────────────────────┐
                    │          Config Server              │
                    │            :8888                    │
                    └─────────────────────────────────────┘
```

---

## 🔧 Servisler

| Servis | Port | Teknoloji | Açıklama |
|--------|------|-----------|----------|
| **API Gateway** | 8060 | Spring Cloud Gateway | Tüm isteklerin geçtiği merkezi giriş noktası |
| **Discovery Server** | 8761 | Netflix Eureka | Servis keşfi ve kayıt merkezi |
| **Config Server** | 8888 | Spring Cloud Config | Merkezi konfigürasyon yönetimi |
| **Ingestion Service** | 8081 | Go + Gin | Araç verilerini alır, Redis'te cache'ler, hız ihlallerini RabbitMQ'ya gönderir |
| **Fleet Management** | 8083 | Spring Boot | Araç ve filo CRUD operasyonları |
| **Alert Service** | 8082 | Spring Boot | Hız ihlali uyarılarını işler ve saklar |
| **Keycloak** | 9090 | Keycloak | OAuth2/OpenID Connect kimlik doğrulama |

---

## 🛠 Teknolojiler

### Backend
- **Java 21** + Spring Boot 3.x
- **Go 1.25** + Gin Framework
- **Spring Cloud** (Gateway, Config, Eureka)

### Veritabanları & Mesajlaşma
- **PostgreSQL** - Kalıcı veri depolama
- **Redis** - Hızlı cache ve oturum yönetimi
- **RabbitMQ** - Asenkron mesajlaşma

### DevOps & Güvenlik
- **Docker & Docker Compose** - Konteynerizasyon
- **Keycloak** - OAuth2/JWT kimlik doğrulama
- **Netflix Eureka** - Servis keşfi

---

## 🚀 Kurulum

### Gereksinimler

- Docker & Docker Compose
- Java 21 (geliştirme için)
- Go 1.25+ (geliştirme için)

### Hızlı Başlangıç

```bash
# Repository'yi klonla
git clone https://github.com/username/FleetPulse.git
cd FleetPulse

# Tüm servisleri build et ve başlat
cd deployments
docker-compose up --build -d

# Servislerin durumunu kontrol et
docker-compose ps
```

### Servislere Erişim

| Servis | URL |
|--------|-----|
| Eureka Dashboard | http://localhost:8761 |
| API Gateway | http://localhost:8060 |
| Keycloak Admin | http://localhost:9090 (admin/admin) |
| RabbitMQ Management | http://localhost:15672 (guest/guest) |
| Config Server | http://localhost:8888 |

---

## 📡 API Endpoints

Tüm istekler **API Gateway** üzerinden yapılır: `http://localhost:8060`

### Ingestion Service

```http
POST /api/v1/ingest
Content-Type: application/json

{
  "vehicleId": "34ABC123",
  "speed": 85.5,
  "latitude": 41.0082,
  "longitude": 28.9784,
  "timestamp": "2026-01-31T12:00:00Z"
}
```

### Fleet Management Service

```http
# Tüm araçları listele
GET /api/v1/vehicles

# Araç detayı
GET /api/v1/vehicles/{id}

# Yeni araç ekle
POST /api/v1/vehicles
{
  "plateNumber": "34ABC123",
  "brand": "Mercedes",
  "model": "Actros",
  "maxSpeed": 90
}

# Araç güncelle
PUT /api/v1/vehicles/{id}

# Araç sil
DELETE /api/v1/vehicles/{id}
```

### Alert Service

```http
# Tüm uyarıları listele
GET /api/v1/alerts

# Araç bazlı uyarılar
GET /api/v1/alerts/vehicle/{vehicleId}

# Uyarı detayı
GET /api/v1/alerts/{id}
```

---

## ⚙ Konfigürasyon

### Environment Variables

Her servis için Docker Compose'da tanımlanan environment variable'lar:

| Variable | Açıklama | Örnek |
|----------|----------|-------|
| `EUREKA_CLIENT_SERVICEURL_DEFAULTZONE` | Eureka server URL | `http://discovery-server:8761/eureka/` |
| `CONFIG_SERVER_URL` | Config server URL | `http://config-server:8888` |
| `SPRING_DATASOURCE_URL` | PostgreSQL bağlantısı | `jdbc:postgresql://host:5432/db` |
| `SPRING_RABBITMQ_HOST` | RabbitMQ host | `rabbitmq` |
| `REDIS_URL` | Redis bağlantısı | `redis:6379` |
| `KEYCLOAK_ISSUER_URI` | Keycloak realm URL | `http://keycloak:9090/realms/FleetPulse` |

### Port Dağılımı

| Port | Servis |
|------|--------|
| 5433 | PostgreSQL (Keycloak) |
| 5434 | PostgreSQL (Alert Service) |
| 5435 | PostgreSQL (Fleet Management) |
| 5672 | RabbitMQ (AMQP) |
| 6379 | Redis |
| 8060 | API Gateway |
| 8081 | Ingestion Service |
| 8082 | Alert Service |
| 8083 | Fleet Management |
| 8761 | Eureka Discovery |
| 8888 | Config Server |
| 9090 | Keycloak |
| 15672 | RabbitMQ Management |

---

## 💻 Geliştirme

### Proje Yapısı

```
FleetPulse/
├── config-repo/                 # Merkezi konfigürasyon dosyaları
│   ├── application.yml
│   └── alert-service.yml
├── deployments/
│   └── docker-compose.yml       # Docker Compose tanımları
└── services/
    ├── alert-service/           # Spring Boot - Uyarı servisi
    ├── config-server/           # Spring Cloud Config Server
    ├── discovery-server/        # Netflix Eureka Server
    ├── fleet-management-service/# Spring Boot - Filo yönetimi
    ├── gateway-service/         # Spring Cloud Gateway
    └── ingestion-service/       # Go - Veri toplama servisi
```

### Local Geliştirme

```bash
# Sadece altyapı servislerini başlat
cd deployments
docker-compose up -d redis rabbitmq keycloak-db keycloak alert-service-db fleet-service-db discovery-server config-server

# Servisleri IDE'den çalıştır
# Her servis localhost'ta default portlarda çalışır
```

### Servisleri Yeniden Build Etme

```bash
# Tek servis
cd deployments
docker-compose up -d --build <service-name>

# Tüm servisler
docker-compose up -d --build
```

### Logları İzleme

```bash
# Tüm loglar
docker-compose logs -f

# Tek servis
docker-compose logs -f ingestion-service
```

---

## 📊 Monitoring

- **Eureka Dashboard:** http://localhost:8761 - Kayıtlı servisleri görüntüle
- **RabbitMQ Management:** http://localhost:15672 - Kuyrukları ve mesajları izle
- **Keycloak Admin:** http://localhost:9090 - Kullanıcı ve realm yönetimi

---

## 🔐 Güvenlik

- Tüm API istekleri **API Gateway** üzerinden geçer
- **Keycloak** ile OAuth2/JWT tabanlı kimlik doğrulama
- Servisler arası iletişim Docker network içinde izole

---

