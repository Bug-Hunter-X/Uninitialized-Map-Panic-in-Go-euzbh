# Uninitialized Map Panic in Go

This repository demonstrates a common error in Go programs: panics caused by accessing uninitialized maps.  Uninitialized maps are `nil`, and attempting to access or modify them results in a runtime panic.

## Bug

The `bug.go` file shows a simple example of this problem.  A map `m` is declared, but not initialized. The program then attempts to assign a value to `m["a"]`. This causes a runtime panic.

## Solution

The `bugSolution.go` demonstrates the solution: always initialize maps before use.  The easiest way to initialize a map is to use the `make` function.