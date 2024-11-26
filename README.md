# tectonic_assets:
*building a web-based dashboard for go profiling, real-time monitoring and analysis of CPU usage, Heap / Memory, Threads, Mutexes, Goroutines and Blocks across multiple applications or services.*

Run migrations to postgres
```bash
goose -dir ./migrations postgres "your_connection_string" up
```
Rollback Migration
```bash
goose -dir ./migrations postgres "your_connection_string" down
```

### Pseudocode Plan

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

> **Note**: project still in progress. ![Progress](https://progress-bar.dev/5/)