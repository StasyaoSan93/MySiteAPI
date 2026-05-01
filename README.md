# 🛠️ ISAProgAPI — Go Site Data Management API

<p align="center">
  <b>A Go REST API for managing website blocks (headers, texts, images, links) with JWT authentication.</b>
</p>

<p align="center">
  <img src="https://img.shields.io/badge/Language-Go-00ADD8?style=flat&logo=go&logoColor=white" />
  <img src="https://img.shields.io/badge/Framework-Gin-00ADD8?style=flat" />
  <img src="https://img.shields.io/badge/Database-PostgreSQL-336791?style=flat&logo=postgresql&logoColor=white" />
</p>

---

# 📑 Table of Contents
- [Project Overview](#-project-overview)
- [Features](#-features)
- [API Endpoints](#-api-endpoints)
- [Tech Stack](#-tech-stack)
- [License](#-license)

---

# 📌 Project Overview

**ISAProgAPI** is a high-performance Go backend application designed to manage website data blocks.  

It allows you to:

- Securely authenticate users via JWT
- Create, read, update, and delete (CRUD) site data
- Easily extend for new site content types

The API is suitable for **portfolio projects, small websites, and learning Go with PostgreSQL**.

---

# 🔍 Features

- 🔐 JWT Authentication  
- 📄 Full CRUD operations for site data  
- 🌐 RESTful API endpoints  
- ⚡ Lightning-fast performance and concurrency powered by Go  
- 🧩 Easy integration with frontend or mobile apps  

---

# 🧩 API Endpoints

| Endpoint | Method | Description |
|----------|--------|-------------|
| `/token` | POST | Authenticate user and receive JWT token |
| `/sitedata/` | GET | Retrieve all site data |
| `/infopagedata/` | GET | Retrieve all data for info page |
| `/sitedata/{id}` | GET | Retrieve site data by ID |
| `/sitedata/` | POST | Create a new site data block |
| ~~`/sitedata/{id}`~~ | ~~PUT~~ | ~~Update an existing site data block~~ |
| ~~`/sitedata/{id}`~~ | ~~DELETE~~ | ~~Delete a site data block~~ |

---

# 🧩 Tech Stack

- **Backend:** Go 1.21+  
- **Framework:** Gin (or standard `net/http`)
- **Database:** PostgreSQL, GORM (or `pgx`)
- **Authentication:** JWT (`golang-jwt`)
- **Environment management:** `godotenv` / `viper`

---

# 📄 License
This project is provided for educational and archival purposes only.
- You may reuse the code for learning or personal non-commercial projects.
- Commercial use by third parties is not allowed.
- Credit to the author is appreciated when used or modified.