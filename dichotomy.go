// This tool cuts a target file recursively
// I used this script to identify malware in archive file without any compression

/*
    Copyright (C) 2023  Maurice Lambert
    This program is free software: you can redistribute it and/or modify
    it under the terms of the GNU General Public License as published by
    the Free Software Foundation, either version 3 of the License, or
    (at your option) any later version.
    This program is distributed in the hope that it will be useful,
    but WITHOUT ANY WARRANTY; without even the implied warranty of
    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
    GNU General Public License for more details.
    You should have received a copy of the GNU General Public License
    along with this program.  If not, see <https://www.gnu.org/licenses/>.
*/

// Compilation on Linux:
// env GOOS=windows GOARCH=amd64 go build -ldflags "-s -w" -trimpath -o CutFileDichotomy.exe CutFileDichotomy.go

package main

import (
    "strconv"
    "fmt"
    "os"
)

type FileSize struct {
    path string
    size int64
}

/*
    This function writes in a new file a chunk of the source file.
*/
func read_write(file *os.File, newfile FileSize, start int64, maxsize int64) {
    halfFile, err := os.Create(newfile.path)
    if err != nil {
        fmt.Fprintln(os.Stderr, "Error creating a half file:", err)
        return
    }
    defer halfFile.Close()

    buffer := make([]byte, newfile.size)

    _, err = file.ReadAt(buffer, start)
    if err != nil {
        fmt.Fprintln(os.Stderr, "Error reading a half file:", err)
        return
    }

    _, err = halfFile.Write(buffer)
    if err != nil {
        fmt.Fprintln(os.Stderr, "Error writing a half file:", err)
        return
    }

    cut(newfile, maxsize)
}

/*
    This function cuts a file on the middle and write two files
    (one for the first part and one for the second).
*/
func cut(filesize FileSize, maxsize int64) {
    if (filesize.size < maxsize) {
        return
    }

    new_file_size := (filesize.size / 2) + 1

    file, err := os.Open(filesize.path)
    if err != nil {
        fmt.Fprintln(os.Stderr, "Error opening a file:", err)
        return
    }
    defer file.Close()

    newfile := FileSize{filesize.path + "0", new_file_size}
    read_write(file, newfile, 0, maxsize)

    newfile.path = filesize.path + "1"
    read_write(file, newfile, filesize.size - new_file_size, maxsize)
}

/*
    The main function to starts the executable.
*/
func main() {
    if len(os.Args) < 2 || len(os.Args) > 3 {
        fmt.Fprintln(os.Stderr, "USAGES: CutFileDichotomy.exe [string:filename] (integer:size=500)")
        return
    }

    size := 500;

    if len(os.Args) == 3 {
        size1, err := strconv.Atoi(os.Args[2])
        size = size1
        if err != nil {
            fmt.Fprintln(os.Stderr, "USAGES: CutFileDichotomy.exe [string:filename] (integer:size=500)")
            fmt.Fprintln(os.Stderr, "Invalid 'size' parameter, should be an integer...")
            return
        }
    }

    fmt.Println("Start cutting file recursively...")
    startFile := FileSize{os.Args[1], 0}

    file, err := os.Open(startFile.path)
    if err != nil {
        fmt.Fprintln(os.Stderr, "Error opening the source file:", err)
        return
    }

    fileInfo, err := file.Stat()
    if err != nil {
        file.Close()
        fmt.Fprintln(os.Stderr, "Error getting file information:", err)
        return
    }

    startFile.size = fileInfo.Size()
    file.Close()

    cut(startFile, int64(size))

    fmt.Println("File successfully cut")
}