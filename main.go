// package main

// import (
// 	"fmt"
// 	"html/template"
// 	"net/http"
// 	"os"
// 	"path/filepath"
// 	"strings"
// )

// var tmpl = template.Must(template.ParseFiles("templates/index.html"))

// func printStartupDiagram() {
// 	fmt.Println("\033[1;34m") // Set color to blue

// 	fmt.Println(`
// ╔══════════════════════════════════════════════╗
// ║                                              ║
// ║   🚀 Go Web Terminal Server is Running!     ║
// ║                                              ║
// ╠══════════════════════════════════════════════╣
// ║                                              ║
// ║   📂 Endpoints:                              ║
// ║      • /            => Index Page           ║
// ║      • /run         => Terminal Emulator    ║
// ║      • /file-structure => List Directory    ║
// ║                                              ║
// ║   🖥️  Listening on: http://localhost:8080     ║
// ║                                              ║
// ╚══════════════════════════════════════════════╝
// `)

// 	fmt.Print("\033[0m") // Reset color
// }

// func main() {
// 	printStartupDiagram()
// 	http.HandleFunc("/", serveIndex)
// 	http.HandleFunc("/run", runCommand)
// 	http.HandleFunc("/file-structure", fileStructure)
// 	http.ListenAndServe(":8080", nil)
// }

// func serveIndex(w http.ResponseWriter, r *http.Request) {
// 	tmpl.Execute(w, nil)
// }

// func runCommand(w http.ResponseWriter, r *http.Request) {
// 	cmdStr := r.FormValue("cmd")
// 	cmdParts := strings.Fields(cmdStr)

// 	// If the user types clear, clear the terminal
// 	if strings.ToLower(cmdStr) == "clear" {
// 		w.Write([]byte("<pre class='text-sm bg-black text-green-400 p-2'># Terminal Cleared</pre>"))
// 		return
// 	}

// 	// If the user types exit, exit the terminal
// 	if strings.ToLower(cmdStr) == "exit" {
// 		w.Write([]byte("<pre class='text-sm bg-black text-green-400 p-2'># Session Terminated</pre>"))
// 		return
// 	}

// 	// Handle the 'pwd' command (print working directory)
// 	if strings.ToLower(cmdStr) == "pwd" {
// 		wd, err := os.Getwd()
// 		if err != nil {
// 			http.Error(w, "Error getting current directory", http.StatusInternalServerError)
// 			return
// 		}
// 		w.Write([]byte(fmt.Sprintf("<pre class='text-sm bg-black text-green-400 p-2'>%s</pre>", wd)))
// 		return
// 	}

// 	// Handle the 'ls' command (list files in the current directory)
// 	if strings.ToLower(cmdParts[0]) == "ls" {
// 		wd, err := os.Getwd()
// 		if err != nil {
// 			http.Error(w, "Error getting current directory", http.StatusInternalServerError)
// 			return
// 		}

// 		files, err := os.ReadDir(wd)
// 		if err != nil {
// 			http.Error(w, "Error reading current directory", http.StatusInternalServerError)
// 			return
// 		}

// 		var fileList string
// 		for _, file := range files {
// 			fileList += file.Name() + "\n"
// 		}
// 		w.Write([]byte(fmt.Sprintf("<pre class='text-sm bg-black text-green-400 p-2'>%s</pre>", fileList)))
// 		return
// 	}

// 	// Handle the 'cat' command (output the contents of a file)
// 	if len(cmdParts) > 1 && strings.ToLower(cmdParts[0]) == "cat" {
// 		filename := cmdParts[1]
// 		wd, err := os.Getwd()
// 		if err != nil {
// 			http.Error(w, "Error getting current directory", http.StatusInternalServerError)
// 			return
// 		}

// 		filePath := filepath.Join(wd, filename)
// 		data, err := os.ReadFile(filePath)
// 		if err != nil {
// 			http.Error(w, "Error reading file", http.StatusInternalServerError)
// 			return
// 		}

// 		w.Write([]byte(fmt.Sprintf("<pre class='text-sm bg-black text-green-400 p-2'>%s</pre>", string(data))))
// 		return
// 	}

// 	// Handle the 'cat' command (output the contents of a file)
// 	if len(cmdParts) > 1 && strings.ToLower(cmdParts[0]) == "cat" {
// 		filename := strings.Join(cmdParts[1:], " ")
// 		wd, err := os.Getwd()
// 		if err != nil {
// 			http.Error(w, "Error getting current directory", http.StatusInternalServerError)
// 			return
// 		}

// 		filePath := filepath.Join(wd, filename)
// 		data, err := os.ReadFile(filePath)
// 		if err != nil {
// 			http.Error(w, fmt.Sprintf("Error reading file: %v", err), http.StatusInternalServerError)
// 			return
// 		}

// 		escaped := template.HTMLEscapeString(string(data))
// 		w.Write([]byte(fmt.Sprintf("<pre class='text-sm bg-black text-green-400 p-2'>%s</pre>", escaped)))
// 		return
// 	}

// 	// Handle file removal (rm command)
// 	if len(cmdParts) > 1 && cmdParts[0] == "rm" {
// 		filename := cmdParts[1]
// 		removeFile(filename, w)
// 		return
// 	}

// 	// Handle file creation (touch command)
// 	if len(cmdParts) > 1 && cmdParts[0] == "touch" {
// 		filename := cmdParts[1]
// 		createFile(filename, w)
// 		return
// 	}

// 	// Handle directory creation (mkdir command)
// 	if len(cmdParts) > 1 && cmdParts[0] == "mkdir" {
// 		dirname := cmdParts[1]
// 		createDirectory(dirname, w)
// 		return
// 	}

// 	// Handle other commands (for now, just echo them)
// 	cmdResult := fmt.Sprintf("$ %s\n", cmdStr)
// 	w.Header().Set("Content-Type", "text/html")
// 	w.Write([]byte("<pre class='text-sm bg-black text-green-400 p-2'>" + cmdResult + "</pre>"))
// }

// func fileStructure(w http.ResponseWriter, r *http.Request) {
// 	wd, err := os.Getwd()
// 	if err != nil {
// 		http.Error(w, "Error getting current directory", http.StatusInternalServerError)
// 		return
// 	}

// 	var structure strings.Builder
// 	structure.WriteString(fmt.Sprintf("Current Directory: %s\n", wd))

// 	files, err := os.ReadDir(wd)
// 	if err != nil {
// 		http.Error(w, "Error reading current directory", http.StatusInternalServerError)
// 		return
// 	}

// 	for _, file := range files {
// 		structure.WriteString(fmt.Sprintf("  %s\n", file.Name()))
// 	}

// 	w.Header().Set("Content-Type", "text/plain")
// 	w.Write([]byte(structure.String()))
// }

// func createFile(filename string, w http.ResponseWriter) {
// 	wd, err := os.Getwd()
// 	if err != nil {
// 		http.Error(w, "Error getting current directory", http.StatusInternalServerError)
// 		return
// 	}

// 	filePath := filepath.Join(wd, filename)
// 	_, err = os.Create(filePath)
// 	if err != nil {
// 		http.Error(w, "Error creating file", http.StatusInternalServerError)
// 		return
// 	}

// 	w.Write([]byte(fmt.Sprintf("<pre class='text-sm bg-black text-green-400 p-2'>%s created successfully in %s</pre>", filename, wd)))
// }

// func removeFile(filename string, w http.ResponseWriter) {
// 	wd, err := os.Getwd()
// 	if err != nil {
// 		http.Error(w, "Error getting current directory", http.StatusInternalServerError)
// 		return
// 	}

// 	filePath := filepath.Join(wd, filename)
// 	err = os.Remove(filePath)
// 	if err != nil {
// 		http.Error(w, "Error removing file", http.StatusInternalServerError)
// 		return
// 	}

// 	w.Write([]byte(fmt.Sprintf("<pre class='text-sm bg-black text-green-400 p-2'>%s removed successfully from %s</pre>", filename, wd)))
// }

// func createDirectory(dirname string, w http.ResponseWriter) {
// 	wd, err := os.Getwd()
// 	if err != nil {
// 		http.Error(w, "Error getting current directory", http.StatusInternalServerError)
// 		return
// 	}

// 	dirPath := filepath.Join(wd, dirname)
// 	err = os.Mkdir(dirPath, 0755)
// 	if err != nil {
// 		http.Error(w, "Error creating directory", http.StatusInternalServerError)
// 		return
// 	}

//		w.Write([]byte(fmt.Sprintf("<pre class='text-sm bg-black text-green-400 p-2'>%s directory created successfully in %s</pre>", dirname, wd)))
//	}
package main

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

var tmpl = template.Must(template.ParseFiles("templates/index.html"))

func printStartupDiagram() {
	fmt.Println("\033[1;34m") // Set color to blue

	fmt.Println(`
╔══════════════════════════════════════════════╗
║                                              ║
║   🚀 Go Web Terminal Server is Running!     ║
║                                              ║
╠══════════════════════════════════════════════╣
║                                              ║
║   📂 Endpoints:                              ║
║      • /            => Index Page           ║
║      • /run         => Terminal Emulator    ║
║      • /file-structure => List Directory    ║
║                                              ║
║   🖥️  Listening on: http://localhost:8080     ║
║                                              ║
╚══════════════════════════════════════════════╝
`)

	fmt.Print("\033[0m") // Reset color
}

var currentWorkingDir = "."

func main() {
	printStartupDiagram()
	http.HandleFunc("/", serveIndex)
	http.HandleFunc("/run", runCommand)
	http.HandleFunc("/file-structure", fileStructure)
	http.HandleFunc("/change-directory", changeDirectory)
	http.ListenAndServe(":8080", nil)
}

func serveIndex(w http.ResponseWriter, r *http.Request) {
	tmpl.Execute(w, nil)
}

func runCommand(w http.ResponseWriter, r *http.Request) {
	cmdStr := r.FormValue("cmd")
	cmdParts := strings.Fields(cmdStr)

	// clear
	if strings.ToLower(cmdStr) == "clear" {
		w.Write([]byte("<pre class='text-sm bg-black text-green-400 p-2'># Terminal Cleared</pre>"))
		return
	}

	// exit
	if strings.ToLower(cmdStr) == "exit" {
		w.Write([]byte("<pre class='text-sm bg-black text-green-400 p-2'># Session Terminated</pre>"))
		return
	}

	// pwd
	if strings.ToLower(cmdStr) == "pwd" {
		wd, err := os.Getwd()
		if err != nil {
			http.Error(w, "Error getting current directory", http.StatusInternalServerError)
			return
		}
		w.Write([]byte(fmt.Sprintf("<pre class='text-sm bg-black text-green-400 p-2'>%s</pre>", wd)))
		return
	}

	// ls
	if strings.ToLower(cmdParts[0]) == "ls" {
		wd, err := os.Getwd()
		if err != nil {
			http.Error(w, "Error getting current directory", http.StatusInternalServerError)
			return
		}
		files, err := os.ReadDir(wd)
		if err != nil {
			http.Error(w, "Error reading current directory", http.StatusInternalServerError)
			return
		}

		var fileList string
		for _, file := range files {
			fileList += file.Name() + "\n"
		}
		w.Write([]byte(fmt.Sprintf("<pre class='text-sm bg-black text-green-400 p-2'>%s</pre>", fileList)))
		return
	}

	// cat
	if len(cmdParts) > 1 && strings.ToLower(cmdParts[0]) == "cat" {
		filename := strings.Join(cmdParts[1:], " ")
		wd, err := os.Getwd()
		if err != nil {
			http.Error(w, "Error getting current directory", http.StatusInternalServerError)
			return
		}
		filePath := filepath.Join(wd, filename)
		data, err := os.ReadFile(filePath)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error reading file: %v", err), http.StatusInternalServerError)
			return
		}
		escaped := template.HTMLEscapeString(string(data))
		w.Write([]byte(fmt.Sprintf("<pre class='text-sm bg-black text-green-400 p-2'>%s</pre>", escaped)))
		return
	}

	// rm
	if len(cmdParts) > 1 && cmdParts[0] == "rm" {
		filename := cmdParts[1]
		removeFile(filename, w)
		return
	}

	// touch
	if len(cmdParts) > 1 && cmdParts[0] == "touch" {
		filename := cmdParts[1]
		createFile(filename, w)
		return
	}

	// mkdir
	if len(cmdParts) > 1 && cmdParts[0] == "mkdir" {
		dirname := cmdParts[1]
		createDirectory(dirname, w)
		return
	}

	// fallback echo
	cmdResult := fmt.Sprintf("$ %s\n", cmdStr)
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte("<pre class='text-sm bg-black text-green-400 p-2'>" + cmdResult + "</pre>"))
}

// 🔥 File structure API
func fileStructure(w http.ResponseWriter, r *http.Request) {
	queryPath := r.URL.Query().Get("path")

	basePath, err := os.Getwd()
	if err != nil {
		http.Error(w, "Error getting current directory", http.StatusInternalServerError)
		return
	}

	targetPath := filepath.Join(basePath, queryPath)

	info, err := os.Stat(targetPath)
	if err != nil || !info.IsDir() {
		http.Error(w, "Invalid or non-existent directory", http.StatusBadRequest)
		return
	}

	var structure strings.Builder
	structure.WriteString(fmt.Sprintf("📁 Directory: %s\n\n", targetPath))
	getFileStructure(targetPath, "", &structure)

	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte(structure.String()))
}

// 🔥 Recursive file walker
func getFileStructure(path string, prefix string, builder *strings.Builder) {
	files, err := os.ReadDir(path)
	if err != nil {
		builder.WriteString(fmt.Sprintf("%s[error reading dir]\n", prefix))
		return
	}

	for _, file := range files {
		if file.IsDir() {
			builder.WriteString(fmt.Sprintf("%s📂 %s/\n", prefix, file.Name()))
			getFileStructure(filepath.Join(path, file.Name()), prefix+"    ", builder)
		} else {
			builder.WriteString(fmt.Sprintf("%s📄 %s\n", prefix, file.Name()))
		}
	}
}

// File creation handler
func createFile(filename string, w http.ResponseWriter) {
	wd, err := os.Getwd()
	if err != nil {
		http.Error(w, "Error getting current directory", http.StatusInternalServerError)
		return
	}

	filePath := filepath.Join(wd, filename)
	_, err = os.Create(filePath)
	if err != nil {
		http.Error(w, "Error creating file", http.StatusInternalServerError)
		return
	}

	w.Write([]byte(fmt.Sprintf("<pre class='text-sm bg-black text-green-400 p-2'>%s created successfully in %s</pre>", filename, wd)))
}

// File deletion handler
func removeFile(filename string, w http.ResponseWriter) {
	wd, err := os.Getwd()
	if err != nil {
		http.Error(w, "Error getting current directory", http.StatusInternalServerError)
		return
	}

	filePath := filepath.Join(wd, filename)
	err = os.Remove(filePath)
	if err != nil {
		http.Error(w, "Error removing file", http.StatusInternalServerError)
		return
	}

	w.Write([]byte(fmt.Sprintf("<pre class='text-sm bg-black text-green-400 p-2'>%s removed successfully from %s</pre>", filename, wd)))
}

// Directory creation handler
func createDirectory(dirname string, w http.ResponseWriter) {
	wd, err := os.Getwd()
	if err != nil {
		http.Error(w, "Error getting current directory", http.StatusInternalServerError)
		return
	}

	dirPath := filepath.Join(wd, dirname)
	err = os.Mkdir(dirPath, 0755)
	if err != nil {
		http.Error(w, "Error creating directory", http.StatusInternalServerError)
		return
	}

	w.Write([]byte(fmt.Sprintf("<pre class='text-sm bg-black text-green-400 p-2'>%s directory created successfully in %s</pre>", dirname, wd)))
}

// changeDirectory is now handling folder uploads and updates the `currentWorkingDir`
// without using a temporary directory.
func changeDirectory(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseMultipartForm(32 << 20); err != nil {
		http.Error(w, "Failed to parse form", http.StatusBadRequest)
		return
	}

	// Get the folder files uploaded
	files := r.MultipartForm.File["folder"]
	if len(files) == 0 {
		http.Error(w, "No folder uploaded", http.StatusBadRequest)
		return
	}

	// Define the directory where the files should be uploaded
	// You can modify this to point to any directory on your server.
	uploadDir := "./uploaded_files" // Change this path as needed

	// Ensure the target directory exists
	if err := os.MkdirAll(uploadDir, 0755); err != nil {
		http.Error(w, "Failed to create directory", http.StatusInternalServerError)
		return
	}

	// Iterate over the uploaded files and save them under the chosen directory
	for _, fileHeader := range files {
		file, err := fileHeader.Open()
		if err != nil {
			continue
		}
		defer file.Close()

		// Save file under the same relative path
		targetPath := filepath.Join(uploadDir, fileHeader.Filename)
		if err := os.MkdirAll(filepath.Dir(targetPath), 0755); err != nil {
			continue
		}

		out, err := os.Create(targetPath)
		if err != nil {
			continue
		}
		defer out.Close()

		io.Copy(out, file)
	}

	// Update the currentWorkingDir to the uploadDir after folder upload
	currentWorkingDir = uploadDir

	// Respond with the updated file structure after the folder is uploaded
	var builder strings.Builder
	builder.WriteString("✅ Folder uploaded into: " + currentWorkingDir + "\n\n")
	filepath.Walk(uploadDir, func(path string, info os.FileInfo, err error) error {
		if err == nil {
			rel, _ := filepath.Rel(uploadDir, path)
			builder.WriteString(rel + "\n")
		}
		return nil
	})

	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprint(w, builder.String())
}
