# To-Do List Application (Angular)

## 🚀 Overview
This is a **To-Do List Application** built using **Angular**, featuring state management with **NgRx**, animations, and local storage support. Users can add, delete, and filter tasks, with smooth UI interactions.

## 📌 Features
- ✅ Add new tasks with a title and description
- ✅ Mark tasks as completed or incomplete
- ✅ Delete tasks
- ✅ Filter tasks (All, Completed, Incomplete)
- ✅ Animations for adding and removing tasks
- ✅ Persistent storage using `localStorage`
- ✅ State management using **NgRx**

## 🛠️ Technologies Used
- **Angular** (v17+)
- **NgRx** for state management
- **Angular Animations** for smooth UI interactions
- **LocalStorage** for data persistence
- **TypeScript & SCSS**

## 📂 Project Structure
```
📦 todo-app
 ┣ 📂 src
 ┃ ┣ 📂 app
 ┃ ┃ ┣ 📂 components
 ┃ ┃ ┃ ┣ 📂 add-to-do       # Add task component
 ┃ ┃ ┃ ┣ 📂 to-do-item      # Single task component
 ┃ ┃ ┃ ┣ 📂 to-do-list      # Task list component
 ┃ ┃ ┣ 📂 services
 ┃ ┃ ┃ ┗ to-do.service.ts   # Handles task logic & local storage
 ┃ ┃ ┣ 📂 store
 ┃ ┃ ┃ ┣ 📂 actions         # NgRx actions
 ┃ ┃ ┃ ┣ 📂 reducers        # NgRx reducers
 ┃ ┃ ┃ ┣ 📂 effects         # NgRx effects (if used)
 ┃ ┃ ┃ ┗ app.state.ts       # Central store
 ┃ ┃ ┣ 📄 app.config.ts      # Application configuration
 ┃ ┃ ┣ 📄 app.routes.ts      # Routing setup
 ┃ ┃ ┗ 📄 main.ts            # Main entry file
```

## 🚀 Installation & Setup
### 1️⃣ Clone the Repository
```sh
git clone https://github.com/your-username/todo-app.git
cd todo-app
```

### 2️⃣ Install Dependencies
```sh
npm install
```

### 3️⃣ Run the Application
```sh
ng serve
```
Then, open your browser and go to: `http://localhost:4200`

## 📝 Usage
- **Add a Task**: Enter a title and optional description, then click **Add**.
- **Filter Tasks**: Use the filter buttons to view **All**, **Completed**, or **Incomplete** tasks.
- **Complete a Task**: Click on a task to mark it as **completed**.
- **Delete a Task**: Click the **delete** button to remove a task.

## 🎯 Bonus Features
✔ **State Management with NgRx** for managing tasks globally.
✔ **Data Persistence with localStorage**, ensuring tasks remain after a page refresh.
✔ **Smooth UI Animations** when adding and removing tasks.

## 📜 License
This project is licensed under the **MIT License**.

---
Made with ❤️ using **Angular**.

