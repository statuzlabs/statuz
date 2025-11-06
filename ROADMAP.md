# 🗺️ Roadmap

This roadmap outlines the development goals for **Statuz** — a lightweight, self‑hosted status and incident management platform.

---

## ✅ Phase 1: Core Foundation
- [ ] Initialize Go project structure
- [ ] Implement configuration system (YAML / env support)
- [ ] Add logging and structured error handling
- [ ] Core monitor engine
  - [ ] HTTP(s) checks
  - [ ] TCP checks
  - [ ] Interval scheduler with goroutines
- [ ] Status aggregation
  - [ ] In‑memory storage
  - [ ] Basic persistence with SQLite or Postgres

---

## 🚧 Phase 2: Incidents & Maintenance
- [ ] Detection of degraded / down states
- [ ] Automatic incident creation and resolution
- [ ] Manual incident creation via API
- [ ] Maintenance window scheduling
- [ ] Basic web/API interface for managing incidents

---

## 🌐 Phase 3: Web UI
- [ ] Simple dashboard for service overview
- [ ] Incident & maintenance history page
- [ ] Authentication for admin area
- [ ] Choose frontend framework (Svelte / htmx / other)
- [ ] Real‑time updates (SSE or WebSocket)

---

## ✉️ Phase 4: Notifications & Subscriptions
- [ ] Notification channels
  - [ ] Email
  - [ ] Webhook
  - [ ] Discord / Slack integration
- [ ] Notification preferences per user / subscriber
- [ ] Subscription management (per service or global)
- [ ] Incident and maintenance alerts

---

## 🐳 Phase 5: Deployment & Polish
- [ ] Dockerfile & docker‑compose setup
- [ ] Configuration via environment variables
- [ ] CLI commands (`statuz start`, `statuz check`, etc.)
- [ ] Example config for self‑hosting
- [ ] Documentation with screenshots
- [ ] Demo deployment (e.g., statuz.sh)

---

## 💡 Future Ideas
- [ ] Role‑based access control (RBAC)
- [ ] Multi‑tenant support
- [ ] Graphs / metrics dashboards
- [ ] API performance statistics
- [ ] External plugin system for custom monitors

---

> **Note:** This roadmap is iterative — features may be re‑prioritized based on feedback and real‑world use.
