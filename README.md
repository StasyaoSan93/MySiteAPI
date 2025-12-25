
# 🛠️ ISAProgAPI — FastAPI Site Data Management API

<p align="center">
  <b>A FastAPI REST API for managing website blocks (headers, texts, images, links) with JWT authentication.</b>
</p>

<p align="center">
  <img src="https://img.shields.io/badge/Language-Python-darkgreen" />
  <img src="https://img.shields.io/badge/Framework-FastAPI-brightgreen" />
  <img src="https://img.shields.io/badge/Database-PostgreSQL-blue" />
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

**ISAProgAPI** is a FastAPI backend application designed to manage website data blocks.  

It allows you to:

- Securely authenticate users via JWT
- Create, read, update, and delete (CRUD) site data
- Easily extend for new site content types

The API is suitable for **portfolio projects, small websites, and learning FastAPI with PostgreSQL**.

---

# 🔍 Features

- 🔐 JWT Authentication  
- 📄 Full CRUD operations for site data  
- 🌐 RESTful API endpoints  
- ⚡ FastAPI async handling for high performance  
- 🧩 Easy integration with frontend or mobile apps  

---

# 🧩 API Endpoints

| Endpoint | Method | Description |
|----------|--------|-------------|
| `/token` | POST | Authenticate user and receive JWT token |
| `/sitedata/` | GET | Retrieve all site data |
| `/sitedata/{id}` | GET | Retrieve site data by ID |
| `/sitedata/` | POST | Create a new site data block |
| `/sitedata/{id}` | PUT | Update an existing site data block |
| `/sitedata/{id}` | DELETE | Delete a site data block |

---

# 🧩 Tech Stack

- **Backend:** FastAPI, Python 3.12  
- **Database:** PostgreSQL, SQLAlchemy  
- **Authentication:** JWT / OAuth2  
- **Async server:** Uvicorn  
- **Environment management:** python-dotenv  

---

# 📄 License
This project is created for educational and portfolio purposes only.
You may reuse the code for learning or personal non-commercial projects.
Commercial use by third parties is not allowed.
Credit to the author is appreciated when publishing or sharing.