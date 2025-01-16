# Go Map Access Before Initialization
This example demonstrates a common error in Go: accessing a map element before the map has been initialized.  This leads to a runtime panic. The solution shows how to properly initialize the map before access.

## Bug
The `bug.go` file contains the erroneous code that attempts to access a map element before the map is initialized. This results in a panic when the code is run.

## Solution
The `bugSolution.go` file provides a corrected version of the code. The map is properly initialized before attempting to access an element. This prevents the runtime panic and makes the code robust.
