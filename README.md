# MCP Go Echo Server

This is a sample implementation of a server for the Mark3Labs Communication Protocol (MCP) written in Go.

The server exposes a simple `echo` tool that takes a text string and an optional boolean flag to return the text in uppercase.

## Prerequisites

- Go programming language installed.

## Running the Server

1.  **Start the server:**

    ```sh
    go run main.go
    ```

2.  The server will start on `http://127.0.0.1:8080` by default.

3.  **Authentication:**
    By default, the server requires a bearer token for authorization. The default token is `secret123` as seen in the `send.sh` script. You can change this by running the server with the `-token` flag:

    ```sh
    go run main.go -token="your-secret-token"
    ```

### Command-line flags

You can customize the server's behavior using the following flags:

-   `-addr`: The HTTP listen address (default: `:8080`).
-   `-token`: The bearer token required for access.
-   `-path`: The base path for the MCP endpoints (default: `/mcp`).

## Using the Server

The `send.sh` script is provided to demonstrate how to interact with the server. The script uses `curl` to send a series of JSON-RPC messages to the server.

1.  **Make sure the server is running.**

2.  **Execute the script:**

    ```sh
    ./send.sh
    ```

The script performs the following actions:
1.  Initializes the MCP connection.
2.  Lists the available tools on the server (which should include `echo`).
3.  Calls the `echo` tool with the text "Ahoj, pepp" and the `uppercase` flag set to `true`.

You will see the JSON-RPC responses from the server for each of these requests printed to your terminal.
