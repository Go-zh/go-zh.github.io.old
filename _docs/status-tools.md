---
layout: default
title: 工具文档翻译状态
summary: 记录了工具文档的翻译状态及任务分配。
---

本文档用于记录工具文档的翻译状态及任务分配。工具文档位于`tools`源码库内的各工具源码注释中。
**注意**：本源码库由于变动太快，现在只翻译了网站模版，其它包暂时不译。

<pre>
tools
├── astutil
├── benchmark
│   └── parse
├── blog
│   └── atom
├── cmd
│   ├── benchcmp
│   ├── callgraph
│   ├── cover
│   ├── digraph
│   ├── eg
│   ├── godex
│   ├── godoc
│   ├── goimports
│   ├── gorename
│   ├── gotype
│   ├── html2article
│   ├── oracle
│   ├── present
│   │   ├── static
│   │   └── templates
│   ├── ssadump
│   ├── stringer
│   ├── tipgodoc
│   └── vet
│       └── whitelist
├── container
│   └── intsets
├── cover
├── dashboard
│   ├── app
│   │   ├── build
│   │   ├── cache
│   │   ├── key
│   │   └── static
│   ├── builder
│   ├── buildlet
│   │   └── stage0
│   ├── coordinator
│   │   └── buildongce
│   ├── env
│   │   ├── commit-watcher
│   │   │   └── scripts
│   │   ├── linux-x86-base
│   │   │   └── scripts
│   │   ├── linux-x86-clang
│   │   │   ├── scripts
│   │   │   └── sources
│   │   ├── linux-x86-gccgo
│   │   │   └── scripts
│   │   ├── linux-x86-nacl
│   │   │   └── scripts
│   │   ├── linux-x86-sid
│   │   │   └── scripts
│   │   ├── openbsd-amd64
│   │   ├── plan9-386
│   │   └── windows
│   ├── retrybuilds
│   ├── types
│   ├── updater
│   └── watcher
├── go
│   ├── buildutil
│   ├── callgraph
│   │   ├── cha
│   │   ├── rta
│   │   └── static
│   ├── exact
│   ├── gccgoimporter
│   ├── gcimporter
│   ├── importer
│   ├── loader
│   ├── pointer
│   ├── ssa
│   │   ├── interp
│   │   └── ssautil
│   ├── types
│   │   └── typeutil
│   └── vcs
├── godoc
│   ├── analysis
│   ├── redirect
│   ├── static        // OlingCat: 完成
│   │   ├── analysis  // OlingCat: *翻译*
│   │   └── images
│   ├── util
│   └── vfs
│       ├── gatefs
│       ├── httpfs
│       ├── mapfs
│       └── zipfs
├── imports
├── oracle
│   └── serial
├── playground
│   └── socket
├── present
└── refactor
    ├── eg
    ├── importgraph
    ├── lexical
    ├── rename
    └── satisfy
</pre>
