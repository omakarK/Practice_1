document.addEventListener('DOMContentLoaded', () => {
    fetchTasks();
});

async function fetchTasks() {
    const response = await fetch('/api/tasks');
    const tasks = await response.json();
    const taskList = document.getElementById('taskList');
    taskList.innerHTML = '';
    tasks.forEach(task => {
        const li = document.createElement('li');
        li.textContent = task.title;
        taskList.appendChild(li);
    });
}

async function addTask() {
    const input = document.getElementById('taskInput');
    const title = input.value;
    if (!title) return;

    const response = await fetch('/api/tasks', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({ title })
    });

    if (response.ok) {
        input.value = '';
        fetchTasks();
    }
}
