=== ASCII Art Generator ===

This program transforms regular text into beautiful ASCII art using three different banner styles. 
Built as part of the Zone01 programming curriculum, it demonstrates file handling, string 
manipulation, and clean modular design in Go.

--- How to Use ---

Basic usage (uses standard banner by default):
    go run . "your text here"

Choose a specific banner style:
    go run . "your text" standard
    go run . "your text" shadow  
    go run . "your text" thinkertoy

--- Examples ---

Simple text conversion:
    go run . "Hello"

Multiple lines with newlines:
    go run . "Hello\nWorld"

Mixed content (letters, numbers, symbols):
    go run . "Hello123!"

--- Error Handling ---

Invalid characters (non-ASCII):
    go run . "Γεια"  # Error: unsupported characters: 'Γ', 'ε', 'ι', 'α'

Wrong number of arguments:
    go run .  # Usage: go run . "text" [banner]

Invalid banner name:
    go run . "Hello" invalid  # Error: Invalid banner name. Available options: standard, shadow, thinkertoy

--- Code Organization ---

The project follows a clean modular structure:

    main.go       -> Program entry point and command line handling
    banner.go     -> Banner file loading and character mapping
    render.go     -> ASCII art rendering and text processing
    validation.go -> Input validation and error handling
    main_test.go  -> Complete test coverage

--- Key Features ---

• Three distinct banner styles (standard, shadow, thinkertoy)
• Proper newline and empty line handling
• Input validation for ASCII characters (32-126)
• Clean, maintainable code architecture
• Robust error handling for edge cases
• Optimized string building for performance

--- Implementation Notes ---

The program works by reading banner template files where each printable ASCII character 
is represented as an 8-line art block. Each character block is separated by one empty line, 
making each character occupy exactly 9 lines in the file.

The rendering process builds output line by line, concatenating the appropriate row from 
each character's art block. Input validation prevents crashes from unsupported characters 
while providing helpful error messages.

--- Testing ---

Run the complete test suite:
    go test

Tests include banner loading, rendering functions, input validation, and various edge 
cases to ensure the program works reliably in all scenarios.