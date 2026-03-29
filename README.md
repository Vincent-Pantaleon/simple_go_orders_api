# Simple Go Orders API

A high-performance, concurrent REST API built with **Go** for managing order lifecycles. This project serves as a deep dive into building idiomatic microservices with a focus on system reliability and lightweight routing.

## 🚀 Key Features
* **Lightweight Routing:** Utilizes [go-chi/chi](https://github.com/go-chi/chi) for composable, idiomatic HTTP routing.
* **Concurrency-Safe IDs:** Implements [google/uuid](https://github.com/google/uuid) for collision-resistant entity identification across distributed instances.
* **In-Memory Speed:** Powered by **Redis** for sub-millisecond data retrieval and storage.
* **Systems Rigidity:** Includes **Graceful Shutdown** handling to ensure in-flight requests are completed before the process exits.
* **Environment Driven:** Fully configurable via environment variables for seamless **Docker** integration.

## 🏗️ Technical Architecture
The project follows a modular structure to separate transport logic from business rules:
- `/application`: Core logic, configuration loading, and server initialization.
- `/repository`: Data access layer (Redis implementation).
- `/handler`: HTTP request/response processing.



## 🛠️ Getting Started

### Prerequisites
* Go 1.21+
* Redis server (running on `localhost:6379`)

### Installation & Run
1. **Clone the repo:**
   ```bash
   git clone [https://github.com/vmpantaleon/simple_go_orders_api.git](https://github.com/vmpantaleon/simple_go_orders_api.git)
   cd simple_go_orders_api
  
2. Install dependencies:
    ```bash
    go mod tidy

3. Run the server:

    ```Bash
    # Defaults to port 3000
    go run main.go

🧪 Engineering Decisions
1. Why Chi?
   - I chose chi over heavier frameworks (like Echo or Gin) to maintain a standard-library feel while gaining powerful middleware capabilities.

2. Why Redis?
   - For this practice project, Redis was selected to practice low-latency data handling and networking between a Go binary and a NoSQL store.
