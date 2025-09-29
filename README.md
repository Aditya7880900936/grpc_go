# Go gRPC Communication Patterns 🚀

[![Go Report Card](https://goreportcard.com/badge/github.com/Aditya7880900936/grpc_go)](https://goreportcard.com/report/github.com/Aditya7880900936/grpc_go)
[![MIT License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)
[![Made With Go](https://img.shields.io/badge/made%20with-Go-00ADD8.svg)](https://golang.org/)

> **A feature-rich showcase of all four core gRPC patterns in Go: Unary, Server Streaming, Client Streaming, and Bidirectional Streaming.**  
> Learn, experiment, and build on top of this example project!

---

## 🖼️ Overview

![gRPC Patterns Banner](https://grpc.io/img/logos/grpc-horizontal-color.png)

This repository demonstrates how to implement different gRPC communication types in Go, including both client and server logic for:

- 🔹 **Unary RPC**
- 🔸 **Server Streaming RPC**
- 🔹 **Client Streaming RPC**
- 🔸 **Bidirectional Streaming RPC**

All code is cleanly organized, fully commented, and easy to extend for your own microservices projects.

---

## 📁 Project Structure

```
grpc_go/
├── client/
│   ├── bi_stream.go          # Bidirectional streaming client
│   ├── client_Stream.go      # Client streaming client
│   ├── main.go               # Client entry point
│   ├── server_Stream.go      # Server streaming client
│   └── unary.go              # Unary client
│
├── proto/
│   ├── greet.proto           # Protocol Buffers definition
│   ├── greet.pb.go           # Generated Go code from .proto
│   └── greet_grpc.pb.go      # gRPC-specific generated code
│
├── server/
│   ├── bi_stream.go          # Bidirectional streaming server
│   ├── client_Stream.go      # Client streaming server
│   ├── main.go               # Server entry point
│   ├── server_Stream.go      # Server streaming server
│   └── unary.go              # Unary server
│
├── go.mod                    # Go module definition
├── go.sum                    # Go dependencies lock file
└── main.go                   # (Optional) root main entrypoint
```

---

## 🏁 Quick Start

### 1. Clone the repo

```bash
git clone https://github.com/Aditya7880900936/grpc_go.git
cd grpc_go
```

### 2. Install Dependencies & Tools

- **Go 1.20+**  
- **Protocol Buffers Compiler (`protoc`)**
- **gRPC Go plugins:**

```bash
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```
Make sure `$GOPATH/bin` is in your system `PATH`.

---

### 3. Generate gRPC Code

If you change `greet.proto`, regenerate Go code:

```bash
protoc --go_out=. --go-grpc_out=. proto/greet.proto
```

---

### 4. Run Server

```bash
cd server
go run main.go
```

---

### 5. Run Client

Open a second terminal:
```bash
cd client
go run main.go
```

Or run specific RPC demos:
```bash
go run unary.go
go run server_Stream.go
go run client_Stream.go
go run bi_stream.go
```

---

## 💡 What are the RPC Types?

| Pattern                    | Description                                                  | Example Scenario               |
|----------------------------|-------------------------------------------------------------|-------------------------------|
| **Unary RPC**              | Single request &rarr; single response                       | Simple API call               |
| **Server Streaming RPC**   | Single request &rarr; response stream                       | Fetching log/event stream     |
| **Client Streaming RPC**   | Request stream &rarr; single response                       | Batch data upload             |
| **Bidirectional Streaming**| Request stream &rarr; response stream                       | Real-time chat, telemetry     |

---

## 🧑‍💻 Example Output

> Here’s what running a unary RPC looks like:

```bash
$ go run client/unary.go
Request: name="Aditya"
Response: "Hello, Aditya!"
```

> Bidirectional streaming example:

```bash
$ go run client/bi_stream.go
Sending: "Ping 1"
Received: "Pong 1"
Sending: "Ping 2"
Received: "Pong 2"
...
```

---

## ✨ Why Use This Repo?

- **Clear separation** of client/server logic
- All core gRPC patterns covered
- Modular file structure for easy navigation
- Minimal, idiomatic Go code
- Ready for production or learning

---

## 🤝 Contributing

Contributions, issues, and PRs are welcome!
- Fork the repo
- Create a branch (`feature/foo`)
- Commit your changes
- Open a PR

---

## 📜 License

This project is licensed under the [MIT License](LICENSE).

---

> **Ready to learn, explore, or build awesome gRPC apps? Dive in!**
