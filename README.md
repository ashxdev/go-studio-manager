# go-dls

This project is a refactor of an application previously written with SvelteKit and PocketBase, now using Go, HTMX, Alpine.js, and Tailwind CSS.

## Stack

- **Go**: The backend server, responsible for handling HTTP requests and serving static files.
- **HTMX**: Handles interactivity with the backend by enabling dynamic updates to the HTML without the need for full page reloads.
- **Alpine.js**: Provides declarative reactive components for a simpler, more lightweight approach to JavaScript behavior.
- **Tailwind CSS**: A utility-first CSS framework for rapid UI development.
- **Tailwind CLI**: Used for building Tailwind CSS without the need for Node.js.

## Prerequisites

- **Go**: Ensure you have Go installed on your system. You can download it from [golang.org](https://golang.org/dl/).
- **Tailwind CSS CLI**: The standalone CLI is used for building Tailwind CSS. No Node.js required. Read more about this here: [Tailwind CSS CLI](https://tailwindcss.com/blog/standalone-cli).

## Getting Started

1. **Clone the repository**

    ```sh
    git clone https://github.com/yourusername/go-dls.git
    cd go-dls
    ```

2. **Setup environment variables**

    Create a `.env` file or rename `.evn.example` in the root of the project and add/modify the following:

    ```plaintext
    HTTP_LISTEN_PORT=3000
    ```

3. **Run Tailwind CSS CLI**

    Download the Tailwind CSS standalone CLI from the [official website](https://tailwindcss.com/blog/standalone-cli) and place it in your project directory.

    **Build Tailwind CSS:**

    ```sh
    ./tailwindcss -i ./static/css/tailwind-input.css -o ./public/styles.css --watch
    ```

    This command will watch for changes and automatically rebuild the CSS.

4. **Run the Go server**

    ```sh
    go run main.go
    ```

    The server will start and be accessible at `http://localhost:3000`.

## Project Structure

- **/public**: Contains static files including the built Tailwind CSS.
- **/src**: Contains source files for Tailwind CSS.
- **/templates**: Contains template files.
- **/handlers**: Contains Go handler functions for various routes.
