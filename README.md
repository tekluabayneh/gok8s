<div align="center">

# ⎈ goK8s

**A Kubernetes Reimplementation, Built From Scratch in Go**

*Learning Kubernetes by rebuilding its core — the API server, scheduler, kubelet, and etcd-backed storage — one component at a time.*

[![Go](https://img.shields.io/badge/Go-88.9%25-00ADD8?style=flat-square&logo=go)](https://go.dev)
[![TypeScript](https://img.shields.io/badge/TypeScript-11.1%25-3178C6?style=flat-square&logo=typescript)](https://www.typescriptlang.org)
[![Status](https://img.shields.io/badge/Status-Work%20in%20Progress-yellow?style=flat-square)](#-status)
[![License](https://img.shields.io/badge/License-MIT-lightgrey?style=flat-square)](#-license)

</div>

---

## 📑 Table of Contents

- [Overview](#-overview)
- [Why Build This](#-why-build-this)
- [Components](#-components)
- [Project Structure](#️-project-structure)
- [Getting Started](#-getting-started)
- [Status](#-status)
- [Author](#-author)

---

## 📖 Overview

**goK8s** is an educational, ground-up reimplementation of core Kubernetes internals, written in Go. Rather than wrapping or extending `kubectl` and the real Kubernetes control plane, this project rebuilds the fundamental pieces — API server, scheduler, kubelet, and cluster state storage — to understand *how* Kubernetes actually works under the hood, not just how to operate it.

It's a systems-level learning project: every component is built by hand, with a deliberate rule of **using AI only for guidance and learning, never to write the core logic itself.**

---

## 🎯 Why Build This

Operating Kubernetes and understanding Kubernetes are two different skills. goK8s exists to close that gap by working through the real mechanics:

- How the **API server** persists and serves cluster state
- How **etcd**-style storage handles object marshaling and watch semantics (via gRPC/HTTP2)
- How the **scheduler** detects unscheduled pods (`spec.nodeName == ""`) and assigns them to nodes
- How a **kubelet** self-registers and continuously reconciles desired vs. actual state
- The distinction between a **node as an etcd record** and a **node as a running process**
- Go-native concurrency patterns (goroutines, channels, `sync.WaitGroup`) as applied to a real control-loop system

---

## 🧩 Components

| Component | Responsibility |
|---|---|
| 🗄️ **API Server / Storage** | Stores cluster object state (JSON-marshaled), exposes a watch mechanism for change notifications. |
| 📅 **Scheduler** | Watches for unscheduled pods and assigns them to available nodes. |
| 🧷 **Kubelet** | Registers its node, reports status, and executes/reconciles pod specs on the node it runs on. |
| 🖥️ **Kubectl** (`Kubectl/`) | A CLI for interacting with the cluster, mirroring the real `kubectl` experience. |
| 🌐 **Cluster** (`cluster/`) | Core cluster orchestration and control-plane logic. |

---

## 🗂️ Project Structure

```
gok8s/
├── Kubectl/          # CLI client for interacting with the cluster
├── cluster/          # Control-plane / cluster orchestration logic
├── Quetion.md         # Notes, open questions, and design decisions
├── mustusedStaff.md   # Reference notes on key patterns and must-know concepts
└── .gitignore
```

> Per-binary commands and shared internal logic follow a `cmd/` + `internals/` convention as components are added — see in-repo notes for current design decisions on `ObjectMeta` and `PodStatus` struct shape, and handler/mapper/decoder organization.

---

## 🚀 Getting Started

```bash
git clone https://github.com/tekluabayneh/gok8s.git
cd gok8s
```

Each component under `cluster/` and `Kubectl/` is a Go module/binary — build and run them individually as they come online:

```bash
go build ./...
```

> As this project is under active development, component entry points and run instructions are evolving — check each directory for its current state.

---

## 🚧 Status

goK8s is an **active, evolving learning project**, not a production-ready Kubernetes alternative. Components are being built incrementally:

- [x] Object storage & JSON marshaling foundations
- [x] Watch mechanism groundwork (gRPC/HTTP2)
- [x] Scheduler: unscheduled pod detection
- [x] Kubelet self-registration
- [ ] Full controller reconciliation loops
- [ ] Networking layer
- [ ] End-to-end `kubectl apply` workflow

---

## 🧑‍💻 Author

**Teklu Abayneh**
*Full-Stack Engineer · Systems & Edge Engineer*

[![GitHub](https://img.shields.io/badge/GitHub-tekluabayneh-181717?style=flat-square&logo=github)](https://github.com/tekluabayneh)

---

<div align="center">

*Understanding Kubernetes, one control loop at a time.*

</div>
