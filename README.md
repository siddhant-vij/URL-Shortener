# URL Shortener

A compact, easy-to-use URL shortening service developed in Go. It allows users to create shortened versions of long URLs, making them easier to share and manage.
<br>
<br>
The service is designed with simplicity and performance in mind, providing both a command-line interface and an HTTP server for web redirection.

<br>

## Table of Contents

1. [Features](#features)
1. [Technical Scope](#technical-scope)
1. [Installation & Usage](#installation-and-usage)
1. [Contributing](#contributing)
1. [Future Improvements](#future-improvements)
1. [License](#license)

<br>

## Features

- **URL Shortening**: Generate short, unique identifiers for long URLs.
- **HTTP Redirection**: Redirects short URLs to their original destinations through an embedded HTTP server.
- **CLI Support**: Offers a command-line interface for shortening URLs and retrieving original URLs.
- **Concurrency Support**: Handles multiple requests concurrently, ensuring high performance.
- **In-Memory Storage**: Utilizes an in-memory repository for fast retrieval and storage of URL mappings.
- **Persistent Storage**: Integrate a database (MongoDB) to persist URL mappings beyond the application lifecycle.

<br>

## Technical Scope

- **Languages and Frameworks**: The project is entirely written in Go, leveraging its standard library for HTTP server functionality & concurrency management.
- **Architecture**: Follows the Model-View-Controller (MVC) design pattern, separating data handling (Model), user interface (View), and application logic (Controller) for clarity and maintainability.
- **Data Storage**: Uses an in-memory data store for quick access and simplicity, with a straightforward path to adapt to database storage for persistence. *Expanded to include MongoDB.*

<br>

## Installation and Usage

1. **Clone the Repository**:
    ```bash
    git clone https://github.com/siddhant-vij/URL-Shortener.git
    ```
2. **Build the Project**: Navigate to the project directory and run the following command:
    ```bash
    cd URL-Shortener
    go build -o urlShortener
    ```
3. **Run the Server**: Start the HTTP server and CLI tool by running the following command:
    ```bash
    ./urlShortener
    ```
4. **Usage: CLI**: Follow the on-screen prompts to perform url shortening operations.
5. **Usage: HTTP Redirection**: Access `http://localhost:8080/<shortened-path>` in a web browser to be redirected to the original URL.

<br>

## Contributing

Contributions are what make the open-source community such an amazing place to learn, inspire, and create. Any contributions you make are **greatly appreciated**.

1. **Fork the Project**
2. **Create your Feature Branch**: 
    ```bash
    git checkout -b feature/AmazingFeature
    ```
3. **Commit your Changes**: 
    ```bash
    git commit -m 'Add some AmazingFeature'
    ```
4. **Push to the Branch**: 
    ```bash
    git push origin feature/AmazingFeature
    ```
5. **Open a Pull Request**

<br>

## Future Improvements

- **Custom Short URL Aliases**: Allow users to specify custom aliases for their shortened URLs.
- **User Authentication**: Implement user accounts and authentication to enable personalized URL management.
- **Dockerization**: Containerize the application with Docker for easy deployment and scaling.
- **Access Statistics**: Track and report on the number of times each shortened URL is accessed.

<br>

## License

Distributed under the MIT License. See [`LICENSE`](https://github.com/siddhant-vij/URL-Shortener/blob/main/LICENSE) for more information.