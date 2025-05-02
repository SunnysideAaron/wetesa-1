# TSDR-000 What and TLDR

## What we are building.

An enterprise CRUD application. Enterprise scale monolith. 10s or even 100s of tables.

## TL;DR of decisions

- TSDR-001 Language(s)
  - Go
- TSDR-002 Data Storage
  - PostgreSQL
- TSDR-003 Docker
  - Using Bitnami PostgreSQL Image
- TSDR-004 DB Initial Data Load
  - Postgress image docker initdb folder 
- TSDR-005 SQL Driver
  - pgx
- TSDR-006 Live Reload of Code
  - [air](https://github.com/air-verse/air)
- TSDR-007 API framework
  - Use the standard library.
- TSDR-008 Possible Future Dependencies
  - **PENDING** See document.
- TSDR-009 Commit Messages
  - Commits without messages are OK. For now.