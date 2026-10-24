# 🖋️ Inkwell — User Documentation & Technical Guide

Welcome to **Inkwell**, a stateful, lightweight, and esoteric journaling scripting language engineered to seamlessly track, log, and map your core life metrics directly within the command line interface. 

Inkwell compiles down into native, single-file standalone binaries with **zero external dependencies** and features an internal timeline database that automatically protects your chronological records, parses inline script commands, and maps human sentiment tracking parameters across a rolling 365-day year trajectory loop.

---

## 🚀 1. Quick Start Installation

Inkwell is distributed as a pre-compiled standalone execution tool. Download the binary matching your operational platform from your secure distribution box and place it directly into your local project workspace folder.

### 🍏 macOS Setup

. Advanced Directory Setup & Permissions

1. Download inkwell_mac from the website.
2. Create a folder named Journal inside your Documents directory (~/Documents/Journal) and move the downloaded binary file there.
3. Open your native Terminal app, navigate to your workspace directory, and grant the binary system execution permissions:bash
cd ~/Documents/Journal
chmod +x inkwell_mac
Use code with caution.
Note: Because this is an unsigned standalone binary, if macOS Gatekeeper blocks execution, navigate to System Settings ➔ Privacy & Security, scroll down to the Security section, and click "Open Anyway" to authorize your engine.

2. Run the Autonomous Initialization Module

Run the engine completely by itself inside your directory with no script arguments attached:
bash
./inkwell_mac
Use code with caution.
Inkwell will detect that no target script was specified, intercept the blank execution, and instantly write a flawless template file named entry.ink right inside your folder directory.

3. Execute Your Script & Timeline Telemetry

To run your daily interactive logging cycle, pass your newly generated blueprint file as an argument into the compiler:
bash
./inkwell_mac entry.ink
Use code with caution.

🎮 The Live Experience Sequence:

• Step 1: The engine pauses execution and prompts you: Rate how today felt (1-10):. Type a score number and hit Enter.
• Step 2: It immediately prompts you: Reflecting on today's milestones, wins, and progress context:. Type your journal reflection text and hit Enter.
• Step 3: The parser automatically serializes your data payloads securely into a localized journal_db.json database on your hard drive, and renders your historical timeline stream as a visual character sparkline tracking graph right inside your console viewport window!

🛑 Multi-Entry Guardrail Verification

Execute ./inkwell_mac entry.ink a second time on the same day. Because your computer's internal system calendar detects that it is still the exact same date, Inkwell intelligently protects your timeline counter. It will refuse to jump ahead to Day 002. Instead, it maintains your perfect active location row index block, updates your database logs with your newest thoughts, and stamps your exact second-accurate execution clock string onto the display row layout. It will wait perfectly until midnight passes your machine's system clock to unlock your next sequential day track!

### 💻 Windows Setup
1. Download `inkwell.exe` to your workspace folder.
2. Open your Command Prompt application (`cmd`).
3. *Note: As an unsigned standalone compiled tool, Windows SmartScreen may show an alert on your first run. Simply click **"More Info"** and select **"Run Anyway"** to grant operational clearance.*

---

## 📜 2. The Inkwell Script Syntax Blueprint (`.ink`)

Inkwell processes custom script files line-by-line using unique execution operators. Create a file named `entry.ink` using any text editor and establish the following flawless procedural script layout flow:

```text
~ =========================================================
~                  INKWELL SCRIPT PROFILE
~ =========================================================
~ Note: Lines starting with a tilde (~) are comments 
~ and are silently skipped by the interpreter engine.

~ Step 1: Interactive Numerical Metric Tracker Element
~ The compiler pauses execution here to collect your daily score benchmark.
note mood = prompt "Rate how today felt (1-10):"

~ Step 2: Interactive Open Narrative Paragraph Logger
~ Pauses a second time to capture open-ended textual thoughts down to the second.
entry "Reflecting on today's milestones, wins, and progress context:"

~ Step 3: Visual Telemetry Dispensation Graph Layout
~ Compiles historical states and updates the vertical block dashboard timeline.
graph mood 365
```

---

## 🏃‍♂️ 3. Executing The Language

You can run the code utilizing the standalone binary directly, or use the automated macro runtime launchers included in your repository.

### Method A: Using the Automated Wrapper Launchers (Recommended)
Open your terminal inside your project workspace folder directory and execute the unified shorthand scripts:
* **On macOS:** `./run entry.ink`
* **On Windows:** `run entry.ink`

*How it works behind the scenes:* The intelligent `run` utility script automatically scans your machine folder system, locates your downloaded compiler tool (even checking your system `~/Downloads` folder directory autonomously), handles all execution file validation parameters, and boots up the parsing sequence natively.

### Method B: Executing the Binary Directly
If you prefer running the direct underlying compiler engine path strings:
* **On macOS:** `./inkwell_mac entry.ink`
* **On Windows:** `inkwell.exe entry.ink`

---

## 📊 4. Core Features & Under-The-Hood Architecture

### 🛡️ Stateful Local Disk Persistence
Inkwell does not require database connection layers or cloud servers. The moment you execute a script, the interpreter serializes your tracking parameters dynamically into a structured, locally stored **`journal_db.json`** workspace file right inside your directory. Your actual entries are entirely secure and reside strictly on your physical hard drive.

### ⏱️ Precise Timestamp Ingestion
Every text record captured through the `entry` keyword is stamped with microsecond-level accuracy. The console visual layout timeline outputs your strict creation tracking marks down to the exact second (e.g., `Entered at 10:24:05 PM`).

### 🛑 Multi-Entry Chronological Guardrails
Unlike primitive text loggers, running your script multiple times back-to-back on the exact same day will never falsely advance your calendar counter or create broken day index placeholders. The parser uses an internal date-tag validation check: if you log entries on the same calendar date, it intelligently preserves your active track index and simply updates your core data structures with your latest real-time reflections and exact execution time variables.

### 🗑️ Data Purging & Administrative Overrides
If you ever make a syntax recording mistake or want to remove an old testing block index out of your active database map stream, simply add the destruction operator directly onto a fresh line inside your script file:
```text
delete entry 1
```
Executing the code file will instantly purge that specific calendar track block index permanently out of your database records.

### ♻️ 365-Day Automatic Reset & Archiving
The architecture scales seamlessly up to a full year. The moment your active logs successfully document **Day 365**, the engine automatically compresses and dumps your entire 365-day tracking cycle timeline matrix into a permanent archive file (e.g., `archive_year_2026.json`), clears the active live workspace layout, and resets your operational counter cleanly back to Day 1 for the upcoming loop cycle.

---

## 🌍 5. Global Terminal Installation (Optional)

To call your language interpreter globally from any folder workspace on your machine without requiring folder navigation prefixes, register the binary paths directly inside your system settings:

#### 🍏 For macOS Users:
1. Open terminal and run this copy command to link the app to your global system path folder:
   ```bash
   sudo cp inkwell_mac /usr/local/bin/inkwell
   ```
2. You can now execute your journal code scripts instantly from any destination folder layout on your machine by simply typing:
   ```bash
   inkwell entry.ink
   ```

#### 💻 For Windows Users:
1. Move `inkwell.exe` into a permanent system directory (e.g., `C:\Program Files\Inkwell`).
2. Search for **"Environment Variables"** in your Windows Search Bar.
3. Select **"Path"** under your System variables registry, click **"Edit"**, choose **"New"**, and paste your folder directory path string (`C:\Program Files\Inkwell`).
4. Open a brand new Command Prompt window—you can execute your language anywhere by simply typing:
   ```cmd
   inkwell entry.ink
   ```
