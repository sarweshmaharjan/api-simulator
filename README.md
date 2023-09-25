# API Simulator

[![Status](https://img.shields.io/badge/Status-Progress-yellow.svg)](https://shields.io/)
[![Version](https://img.shields.io/badge/Version-1.0-blue.svg)](https://shields.io/)

## Table of Contents

- [Introduction](#introduction)
- [Getting Started](#getting-started)
  - [Prerequisites](#prerequisites)
  - [Installation](#installation)
- [Usage](#usage)
- [Project Structure](#project-structure)
- [Contributing](#contributing)

## Introduction

The API simulator project replicates API integration responses, primarily Truepill and Thriftywhite. It is designed to expedite testing and debugging within our system.

## Getting Started

### Prerequisites

List any prerequisites or dependencies required to run your project. Include links or installation instructions if needed.

### Installation

1. Clone the repository:

   ```sh
   git clone git@github.com:sarweshmaharjan/api-simulator.git
   cd api-simulator
   ```

2. Start the project

   ```sh
   make run #To create docker instance of postgres.
   make update-go #To set up go version and install
   ```

3. Configure CAPI:

- Set "truepill.apiUrl" to "http://localhost:9999".
- Ensure feature.IsCopayCallbackEnabled() is set to true for copay callback functionality.

### Project Structure

```sh
├── adrs
│   └── index.md
├── api
│   ├── handler.go
│   └── truepill.go
├── common
│   └── common.go
├── config
│   └── config.go
├── docker-compose.yml
├── Dockerfile
├── go.mod
├── go.sum
├── main.go
├── Makefile
├── README.md
├── storage
│   └── db.go
└── types
    └── truepill.go

```

### Contributing

- [Sarwesh Maharjan](https://github.com/sarweshmaharjan)
- [Suman Shah](https://github.com/shahsuman438)
