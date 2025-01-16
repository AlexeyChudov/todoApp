// const content = document.querySelector('.tasks-content');
// const itemsPerPage = 10;
// let currentPage = 0;
// const items = Array.from(content.getElementsByTagName('ul'));
//
// function showPage(page) {
//     const startIndex = page * itemsPerPage;
//     const endIndex = startIndex + itemsPerPage;
//     items.forEach((item, index) => {
//         item.classList.toggle('visually-hidden', index < startIndex || index >= endIndex);
//     });
//     updateActiveButtonStates();
// }
// function createPageButtons() {
//     const totalPages = Math.ceil(items.length / itemsPerPage);
//     const paginationContainer = document.createElement('div');
//     const paginationDiv = document.body.appendChild(paginationContainer);
//
//     paginationContainer.classList.add('page-item');
//
//     for (let i = 0; i < totalPages; i++) {
//         const pageButton = document.createElement('button');
//         pageButton.textContent = i + 1;
//         pageButton.addEventListener('click', () => {
//             currentPage = i;
//             showPage(currentPage);
//             updateActiveButtonStates();
//         });
//
//         content.appendChild(paginationContainer);
//         paginationDiv.appendChild(pageButton);
//     }
// }
//     function updateActiveButtonStates() {
//         const pageButtons = document.querySelectorAll('.pagination button');
//         pageButtons.forEach((button, index) => {
//             if (index === currentPage) {
//                 button.classList.add('active');
//             } else {
//                 button.classList.remove('active');
//             }
//         });
//     }
//     createPageButtons();
//     showPage(currentPage);


const tasks = []; // Этот массив будет содержать задачи (полученные с сервера).
const tasksPerPage = 10; // Количество задач на странице.
let currentPage = 1; // Текущая страница.

// Функция для отображения задач на текущей странице
function renderTasks() {
    const taskList = document.getElementById('task-list');
    taskList.innerHTML = ''; // Очистить старый список

    // Определяем задачи для отображения
    const start = (currentPage - 1) * tasksPerPage;
    const end = start + tasksPerPage;
    const currentTasks = tasks.slice(start, end);

    // Добавляем задачи в DOM
    currentTasks.forEach(task => {
        const taskElement = document.createElement('div');
        taskElement.className = 'card mb-2';
        taskElement.innerHTML = `
            <div class="card-body">
                <h5 class="card-title">${task.title}</h5>
                <p class="card-text">Статус: ${task.status}</p>
                <p class="card-text">Дедлайн: ${task.dueDate}</p>
            </div>
        `;
        taskList.appendChild(taskElement);
    });
}

// Функция для рендеринга кнопок пагинации
function renderPagination() {
    const pagination = document.getElementById('pagination');
    pagination.innerHTML = ''; // Очистить старые кнопки

    const totalPages = Math.ceil(tasks.length / tasksPerPage);

    for (let i = 1; i <= totalPages; i++) {
        const pageItem = document.createElement('li');
        pageItem.className = `page-item ${i === currentPage ? 'active' : ''}`;
        pageItem.innerHTML = `<a class="page-link" href="#">${i}</a>`;
        pageItem.addEventListener('click', (e) => {
            e.preventDefault();
            currentPage = i;
            renderTasks();
            renderPagination();
        });
        pagination.appendChild(pageItem);
    }
}

// Получаем данные с сервера и инициализируем пагинацию
fetch('/api/tasks') // Замените на ваш API
    .then(response => response.json())
    .then(data => {
        tasks.push(...data); // Загружаем данные в массив
        renderTasks(); // Рендерим первую страницу задач
        renderPagination(); // Создаем кнопки пагинации
    })
    .catch(error => console.error('Ошибка загрузки задач:', error));
