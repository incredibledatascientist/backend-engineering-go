# Why Use a CLI Framework for a Simple Go Network Checker?

When building a small Go program that checks whether a domain is reachable using `net.Dial`, it may seem unnecessary to use a full CLI framework like `urfave/cli`. After all, the core functionality can be implemented with just a few lines of Go code.

This document explains the difference between **simple interactive programs** and **proper command-line tools**, and when using a CLI framework makes sense.

---

# 1. Simple Interactive Program (No CLI Framework)

A basic implementation might read input from the terminal and check the domain status repeatedly.

Example:

```go
for {
    input := readFromStdin()
    status(domain)
}
```
