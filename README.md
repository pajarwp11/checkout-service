# Checkout Backend Service

## Overview

This project is a simple checkout backend service that supports multiple promotion types.

Implemented promotions:

- Each sale of a MacBook Pro comes with a free Raspberry Pi B
- Buy 3 Google Homes for the price of 2
- Buying more than 3 Alexa Speakers will get a 10% discount on all Alexa speakers

---

# Tech Stack

- Golang
- MySQL
- Docker

---

# Run Application

## Start Services

```bash
docker compose up -d
```
```bash
go run cmd/api/main.go
```
---

# API Server

Server will run on:

```txt
http://localhost:8080
```

---

# Checkout Endpoint

## Endpoint

```http
POST /checkout
```

---

# Request Payload

```json
{
  "items": [
    {
      "sku": "120P90",
      "quantity": 3
    }
  ]
}
```

---

# Response Payload

```json
{
  "subtotal": 149.97,
  "discount": 49.99,
  "total": 99.98
}
```

---

# Example Requests

## Buy 3 Google Homes

### Request

```json
{
  "items": [
    {
      "sku": "120P90",
      "quantity": 3
    }
  ]
}
```
---

## Buy MacBook Pro And Raspberry Pi B

### Request

```json
{
  "items": [
    {
      "sku": "43N23P",
      "quantity": 1
    },
    {
      "sku": "234234",
      "quantity": 1
    }
  ]
}
```
---

## Buy 3 Alexa Speakers

### Request

```json
{
  "items": [
    {
      "sku": "A304SD",
      "quantity": 3
    }
  ]
}
```
---

# Database Design

Database design document:

https://docs.google.com/document/d/1bYAmSexijaXtAVsoOjAeYOJ8UlSuv1SEjwAmWqWeHo4/edit?tab=t.0

---
