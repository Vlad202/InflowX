<p align="center">
  <img src="assets/logo.png" alt="Inflow X Logo" width="100%"/>
</p>


**Inflow X** is an event-driven, real-time workflow platform for capturing form submissions from any website and delivering them instantly to external systems (CRMs, Telegram, Webhooks, Emails, and more).  
Built for developers and teams who need ultra-fast lead handling, flexible routing logic, and simple analytics.

---

# 🚀 Features

- ⚡ **Real-time delivery** — Process and route events in milliseconds.
- 📦 **SDK & API** — Easily embed forms or send events from any frontend/backend.
- 🔁 **Workflow engine** — Define multi-step logic: filters, conditions, routing.
- 🔔 **Instant notifications** — Send to Telegram, Webhook, Email, Discord, etc.
- 📊 **Built-in analytics** — Measure drop-offs, lead sources, time to response.
- 🔐 **Secure API access** — Token-based authentication for forms and API.
- 🧩 **Modular architecture** — Each service is scalable, isolated, and replaceable.

---

# 🧱 Architecture Overview

```text
                            +-------------+
                            | Website SDK |
                            +-------------+
                                   |
                             [ Ingest API ]
                                   |
                             +-----------+
                             | Ingest Go |
                             +-----------+
                                   |
                            [ Redis Streams ]
                                   |
        +----------------+----------------+----------------+
        |                |                |                |
+----------------+  +----------------+  +----------------+  +----------------+
| Workflow Engine|  | Notification   |  | Analytics Saver|  | Admin API (Nest)|
|     (NestJS)   |  |   (NestJS)     |  |     (Go)       |  |     (NestJS)    |
+----------------+  +----------------+  +----------------+  +----------------+
                                                          |
                                                      [ PostgreSQL / Clickhouse ]
                                                          |
                                                      +------------------+
                                                      | Nuxt 3 Dashboard |
                                                      +------------------+
