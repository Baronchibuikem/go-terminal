```markdown
# Terminal Web Server with Folder Selection and Management

This project is a Go-based web server that simulates a terminal interface in a browser. It allows you to:

- **Navigate directories**: View and change the working directory.
- **Upload folders**: Upload folders to the server, and the structure is displayed.
- **Execute commands**: Run basic terminal commands like `pwd`, `ls`, `touch`, `rm`, and `mkdir`.

## Features

- Upload and display folder contents.
- Change and manage directories dynamically.
- Execute terminal commands from the browser.

## Setup

1. Clone the repository:
   ```bash
   git clone https://github.com/yourusername/file-upload-server.git
   cd file-upload-server
   ```

2. Build and run:
   ```bash
   go build -o file-upload-server
   ./file-upload-server
   ```

3. Access the terminal interface at [http://localhost:8080](http://localhost:8080).

## Commands

- **`pwd`**: Displays the current working directory.
- **`ls`**: Lists files in the current directory.
- **`mkdir <folder>`**: Creates a new folder.
- **`touch <file>`**: Creates a new file.
- **`rm <file>`**: Deletes a file.

## License

This project is available under the [MIT License](LICENSE).
```

This is the entire project README in a single `.md` file. Just save it as `README.md` in your project directory.