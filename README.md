# API Simulator

A simple application that will help in simulating response after external API System is hit.

## Table of Contents

- [Introduction](#introduction)
- [Getting Started](#getting-started)
  - [Prerequisites](#prerequisites)
  - [Installation](#installation)
- [Usage](#usage)
- [Project Structure](#project-structure)
- [Contributing](#contributing)
- [License](#license)

## Introduction

Simulating the API response.

Features

1. Add simulation to reponse in both success and failure.
2. It enables continuous flow testing in local without breaking the application.

## Getting Started

### Prerequisites

List any prerequisites or dependencies required to run your project. Include links or installation instructions if needed.

### Installation

1. Clone the repository:

   ```sh
   git clone git@github.com:sarweshmaharjan/api-simulator.git
   ```

   ```sh
   cd api-simulator
   ```

2. Start the project

```sh
make run
```

3. Go to PHIL Local DB and set

- key: "truepill.apiUrl" to "http://localhost:9999"

### Project Structure

- cmd/ # Application entry point
- internal/ # Internal packages
- pkg/ # Exported packages
- test/ # Unit tests
- data/ # Data files (if any)
- go.mod # Go module file
- go.sum # Go module checksum file
- main.go # Main application file
- README.md # This README file

### License

This project is licensed under the License Name - see the LICENSE file for details.

### Contributing

- [Sarwesh Maharjan](https://github.com/sarweshmaharjan)
