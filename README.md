# Fit&Roll
A cross-platform desktop application for managing courses and tracking participant engagement.

## Table of Contents
- [⚙️ Starting the application](#-starting-the-application)
- [🌍 Changing locale](#-changing-locale)
- [📚 Managing Courses: Create, Update & Archive](#-managing-courses-create-update--archive)
- [🗂️ Managing Participants: Create, Update & Archive](#-managing-participants-create-update--archive)
- [🏅 Member card overview](#-member-card-overview)
- [🧾 Attending courses](#-attending-courses)
- [📈 Course Participation History](#-course-participation-history)
- [📊 Statistics](#-statistics)
- [🗃️ Archive](#-archive)
- [📤 Export](#-export)
- [🖼️ Application Gallery](#-application-gallery)
- [👨‍💻 For developers](#-for-developers)
- [📑 License](#-license)


## ⚙️ Starting the Application

- The compiled executables are located in the `release` directory.
- On **Linux**, download and use the provided binary file.
- On **Windows**, download the `.exe` file to run the application.
- Create a new folder on your machine and download the application into this folder.
- When the application is started for the first time, a new file with the `.db` extension will be generated inside the folder.
- The `.db` file is a local database that stores all the application's data.
- **⚠️ Caution**: Deleting or renaming the `.db` file will erase all the application data.
- You are encouraged to make copies of the `.db` file for backup purposes.

## 🌍 Changing Locale

- The application supports two languages: **German** (default) and **English**.
- To access the language selector, simply click on the **logo** in the top bar.
- To change the language, click on the **language selector** and choose your preferred language.

Here’s a preview of the language selection process:

![Change Locale](documentation/change_locale.gif)

## 📚 Managing Courses: Create, Update & Archive

- To track the activity, you first need to **create a course**.
- Click the **➕ plus icon** to add a new course.
- Enter the **course name** and its **schedule**. A **description** is optional.
- You can **edit course details** (name, schedule, description) at any time.
- If a course is no longer active, you can **archive** it.
- All currently active courses are listed in the **Course Overview** table.

Here’s a preview of how to manage your courses:

![Manage Courses](documentation/manage_courses.gif)

## 🗂️ Managing Participants: Create, Update & Archive

- To track participation, you need to **add participants** to the application.
- Use the **Participants Overview** section to create a new participant.
- Each participant must have a **full name** and be assigned to a **group**.
- Participant details can be **updated at any time**.
- If a participant no longer attends any trainings, you can **archive** them.
- All active participants are displayed in the **Participants Overview** table.

Here’s a preview of how to manage participants:

![Manage Participants](documentation/manage_participants.gif)

## 🏅 Member Card Overview

- To participate in courses, a participant must first purchase a **member card**.
- A participant can own **multiple member cards**.
- If a member card was issued by mistake and the cardholder hasn't attended any courses, the card can be **returned**.
- Clicking on a member card will display the **history of attended courses**.

Here’s a preview of the **Member Card** component in action:

![Member Card](documentation/member_card.gif)

## 🧾 Attending Courses

- Participants can attend a course in **three ways**:  
  1. Using a **member card**  
  2. Taking a **trial attendance**  
  3. Attending **without a member card**
  
- If a participant was marked as attending by mistake, you can remove the entry via the **Overall Participant History** component.

- ⚠️ **Important:**  
  Removing an attendance that used a **member card** will **restore one available slot** on that card.

![Attending Courses](documentation/attending_courses.gif)

## 📈 Course Participation History

- Participation can be tracked in **two ways**:
  1. **Per course** — View attendance history for a specific course.
  2. **Overall history** — Use the **Participation History** component to see attendance across all courses.

- 🧹 You can also **remove attendance entries** directly from the **Overall Participation History** view if needed.

- Each participation history view offers **advanced filtering options** to help you quickly find relevant data.

![Participation History](documentation/participation_history.gif)

## 📊 Statistics

- For each participant, you can view:
  - **How many courses** they have attended.
  - **How they attended**: using a member card, as a trial, or without a card.

- For each course, you can see:
  - **How often** it was attended.
  - **The attendance method** used by participants.

- All statistics views include a **date range filter** to help you narrow down and focus your analysis.
- All statistics also include archived entries.

![Statistics](documentation/statistics.gif)

## 🗃️ Archive

- View all **archived participants** and **restore** them when needed.
- Access **archived courses** and bring them back into the active view.
- Both archive views include **filtering options** to help you quickly find specific archived entries.

![Archive](documentation/archive.gif)

## 📤 Export

- For advanced analysis, you can **export all current database data** to an **Excel file**.

![Export](documentation/export.gif)

## 🖼️ Application Gallery

### Course Overview
![All Courses](documentation/images/all_courses.png)
![Add Course](documentation/images/new_course.png)
![Update Course](documentation/images/update_course.png)
![Archive Course](documentation/images/archive_course.png)

### Participants Overview
![All Participants](documentation/images/all_participants.png)
![New Participant](documentation/images/new_participant.png)
![Update Participant](documentation/images/update_participant.png)
![Archive Participant](documentation/images/archive_participant.png)


### Cards Overview
![Cards Overview](documentation/images/cards.png)
![Card Participation History](documentation/images/card_history.png)
![Remove Unused Card](documentation/images/unused_card.png)


### Attending Courses
![Attending Courses Overview](documentation/images/attend_course.png)
![Confirmation Modal](documentation/images/attend_course_confirm.png)

### Course Participation History
![Course Participation History](documentation/images/participation_history.png)

### Statistics
![Participant statistics](documentation/images/participants_statistics.png)
![Course statistics](documentation/images/course_statistics.png)

### Archive
![Overall](documentation/images/archive.png)
![Archived courses](documentation/images/archived_courses.png)
![Archived participants](documentation/images/archived_participants.png)

## 👨‍💻 For developers

### 🚀 Local Development

To run the application in live development mode, execute:
```bash
wails dev
```

This project uses Tailwind CSS for styling. To enable automatic generation of classes during development, run from the project directory:
```bash
npx @tailwindcss/cli -i ./frontend/src/assets/tailwind/style.css -o ./frontend/src/tailwind.css --watch
```
This command watches for changes and updates the compiled CSS in real time.
### 🛠️ Building the app
To build a redistributable, production mode package, use
```bash
wails build
```
For more details and advanced options, visit the official Wails documentation.

## 📑 License

This project is licensed under the Apache License 2.0 - see the [LICENSE](./LICENSE) file for details.
