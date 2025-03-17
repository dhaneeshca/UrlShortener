# Custom URL Shortener with API Access

## 📌 Objective
Develop a **custom URL shortener service** using **Go** that allows users to shorten URLs with advanced features such as **custom short URLs, custom domains, expiration time, analytics, and API access with API keys**.

---

## 🚀 Features

### 🔗 URL Shortening
- Shorten a **long URL** into a **short URL**.
- Support **custom short URLs**.
- Validate URL format before shortening.

### 🌐 Custom Domain Support
- Users can provide their own domain for shortened URLs (e.g., `short.mydomain.com/xyz`).
- Validate domain ownership.

### ⏳ Expiration Time
- Users can set an **expiration date** for short URLs.
- Expired URLs automatically become inactive.

### 🔀 URL Redirection
- Redirect users to the original long URL when accessing a short URL.
- Display an appropriate message if the URL is expired.

### 📊 Click Analytics
- Track **total clicks** per URL.
- Capture **referrer information**.
- Detect **device & browser** used.
- IP-based **geolocation tracking** (country, city).

### 🔑 API Access with API Keys
- Provide **REST API endpoints** for URL management.
- Secure API requests using **API keys**.
- Users can **generate, revoke, and regenerate** API keys.
- Implement **rate limiting** to prevent abuse.

### 👤 User Authentication (Optional)
- Users can **sign up & log in**.
- Manage URLs from a **user dashboard**.

### 🔒 Security & Rate Limiting
- Implement **rate limiting** to prevent spam.
- Validate user inputs to prevent **malicious URL injections**.

### 🎯 Additional Features (Optional)
- **Password-protected links**.
- **One-time-use links**.
- **QR Code Generation** for shortened URLs.
- **Bulk URL shortening** via CSV upload.
- **Custom redirect page** before redirection.

---

## 🏗️ Tech Stack
- **Language**: Go
- **Web Framework**: Gin or Fiber (for API)
- **Database**: PostgreSQL, SQLite, or Redis (for caching)
- **Authentication**: JWT (for user login)
- **Caching**: Redis (for fast URL lookups)
- **Rate Limiting**: Middleware like `gin-limiter`

---

## 📖 API Endpoints

### 🌍 Public Endpoints
- `POST /shorten` – Shorten a URL (supports custom URLs, expiration).
- `GET /{short_url}` – Redirect to the original URL.

### 🔒 Protected API Endpoints (Require API Key)
- `POST /api/shorten` – Shorten a URL.
- `GET /api/urls/{short_url}` – Retrieve URL details.
- `DELETE /api/urls/{short_url}` – Delete a short URL.
- `GET /api/urls/{short_url}/stats` – Get analytics for a short URL.
- `POST /api/apikey/generate` – Generate a new API key.
- `GET /api/apikey/list` - List all keys
- `POST /api/apikey/revoke` – Revoke an API key.

---

## 🛠️ Development Setup

### 🔹 1. Clone Repository
```bash
git clone https://github.com/yourusername/url-shortener-go.git
cd url-shortener-go
```

### 🔹 2. Install Dependencies
```bash
go mod tidy
```

### 🔹 3. Set Up Environment Variables
Create a `.env` file:
```
DB_CONNECTION=postgres
DB_HOST=localhost
DB_PORT=5432
DB_USER=youruser
DB_PASSWORD=yourpassword
DB_NAME=postgres
REDIS_HOST=localhost
REDIS_PORT=6379
JWT_SECRET=your_secret_key
```

### 🔹 4. Run the Application
```bash
go run main.go
```

---

## 📅 Development Roadmap

### ✅ Phase 1: Basic URL Shortening
- Implement **shorten URL** logic.
- Store shortened URLs in the database.
- Implement **redirection** functionality.

### ✅ Phase 2: Advanced Features
- Add **custom short URLs**.
- Implement **custom domain support**.
- Add **URL expiration logic[default - 1000 days / custom]**.


### ✅ Phase 3: API Key Authentication
- Implement **API key management**.
- Secure API endpoints using API keys.

### ✅ Phase 4: Analytics & Security
- Implement **click tracking**.
- Secure API with **rate limiting & validation**.

### ✅ Phase 5: Deployment
- Dockerize the application.
- Deploy to **AWS, DigitalOcean, or Render**.

---

## 🎯 Expected Outcomes
✔ A **fully functional URL shortener** with an easy-to-use API.
✔ **Secure API key access** for developers.
✔ **Advanced analytics & tracking**.
✔ **Scalable architecture** with caching & rate limiting.

---

## 🤝 Contributing
Feel free to fork this repository and submit pull requests! 🎉

---
