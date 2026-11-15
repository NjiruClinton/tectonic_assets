# Tectonic Assets

Tectonic Assets is an early-stage development profiler for Go applications. The
current implementation records process-level CPU and resident-memory samples in
PostgreSQL and displays them in a web dashboard.

The dashboard and collector are intentionally independent: the dashboard can
start when PostgreSQL or the configured target process is unavailable.

## Local development

1. Copy `.env.template` to `.env` and adjust the connection string.
2. Start PostgreSQL with `docker compose up -d postgres`.
3. Run the migrations shown below.
4. Start the application with `go run .`.
5. Open `http://localhost:8080`.

To collect process-level metrics, set `TARGET_PROCESS_PATH` to the absolute path
of a running executable. If it is empty, only the dashboard is started.

### Configuration

| Variable | Default | Purpose |
| --- | --- | --- |
| `SERVER_ADDRESS` | `:8080` | Dashboard listen address |
| `DB_CONNECTION_STRING` | none | PostgreSQL connection URL |
| `TARGET_PROCESS_PATH` | none | Executable to monitor |
| `TARGET_NAME` | executable filename | Display name for the target |
| `PROFILE_INTERVAL` | `5s` | CPU and memory sample interval |

The target path must refer to a process visible from the Tectonic process. A
Docker container cannot normally inspect arbitrary processes running on the
host, so run Tectonic directly on the host while developing the collector.

## Database migrations

Run migrations to postgres
```bash
goose -dir ./migrations postgres "your_connection_string" up
```
Rollback Migration
```bash
goose -dir ./migrations postgres "your_connection_string" down
```

## Roadmap

1. **Define Data Collection Agent**:
    - Create a lightweight agent that can be integrated into different applications to collect CPU usage data.
    - The agent sends data to your central server at specified intervals.

2. **Set Up Central Server**:
    - Develop a server application to receive, process, and store data sent by agents.
    - Implement authentication to ensure data security.

3. **Design Database Schema**:
    - Design a schema to efficiently store and query CPU usage data, timestamps, application identifiers, etc.

4. **Implement Data Processing Logic**:
    - Develop logic to aggregate and analyze data from multiple sources.
    - Implement features like real-time monitoring, historical analysis, and alerting based on predefined thresholds.

5. **Develop Web Dashboard**:
    - Create a web-based interface where users can view real-time data, historical charts, and analytics.
    - Implement user authentication and authorization for data access control.

6. **Testing and Deployment**:
    - Thoroughly test the entire system for performance and security.
    - Deploy the server and web dashboard, and distribute the data collection agent.

### Components

1. **Data Collection Agent**
2. **Central Server**
3. **Web Dashboard**

The next milestone is bounded CPU-profile capture through Go's `pprof`
endpoints. Process CPU/RSS sampling will remain as the low-cost live metrics
layer.

> **Note**: project is still in progress.
