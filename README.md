# 🧹 Temp Cleaner Go

**Temp Cleaner** is a small utility written in Go to automatically clean up the temporary folders on a Windows system.  
It clears both:

- `C:\Windows\Temp`
- `%USERPROFILE%\AppData\Local\Temp`

## 💡 What It Does

- Lists the contents of the two temp folders.
- Deletes every file and folder found inside them.
- Logs any error that happens.
- Prints progress messages to the terminal.
- Waits for user input before exiting.

## 📁 Log File

If something goes wrong, a log file will be created in:

`C:\Users<YourName>\TEMP\temp<timestamp>.txt`

Make sure the `TEMP` folder exists in your user directory, or it may fail to write.

## ⚙️ Requirements

- **Windows OS**
- **Go installed** (version 1.18+ recommended)

## ▶️ How to Run

Run directly with:

```bash
go run main.go
```

Or build it first:

```bash
go build -o tempcleaner.exe
./tempcleaner.exe
```

> 💡 Tip: Run as Administrator to fully clean C:\Windows\Temp.

## ⚠️ Warning

This tool permanently deletes files and folders from temp directories.
Use it with caution — deleted items **CANNOT** be recovered.
