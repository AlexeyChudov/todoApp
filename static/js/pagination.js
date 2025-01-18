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
const tasksPerPage = 5; // Количество задач на странице.
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
        taskElement.innerHTML = `<ul class="list-group list-group-horizontal rounded-0 bg-transparent pagination-sm">
                            <li class="list-group-item d-flex align-items-center ps-0 pe-3 py-1 rounded-0 border-0 bg-transparent">
                                <div class="form-check">
                                    <input
                                            class="form-check-input me-0"
                                            type="checkbox"
                                            value=""
                                            id="flexCheckChecked1"
                                            aria-label="..."

                                    />
                                </div>
                            </li>
                            <li class="list-group-item px-3 py-1 d-flex align-items-center flex-grow-1 border-0 bg-transparent">
                                <p class="lead fw-normal mb-0">${task.title}</p>
                            </li>
                            <li class="list-group-item ps-3 pe-0 py-1 rounded-0 border-0 bg-transparent">
                                <div class="d-flex flex-row justify-content-end mb-1">
                                    <a href="#!" class="text-info" data-mdb-toggle="tooltip" title="Edit todo"><i class="fas fa-pencil-alt me-3"></i></a>
                                    <a href="#!" class="text-danger" data-mdb-toggle="tooltip" title="Delete todo"><i class="fas fa-trash-alt"></i></a>
                                </div>
                                <div class="text-end text-muted">
                                    <a href="#!" class="text-muted" data-mdb-toggle="tooltip" title="Created date">
                                        <p class="small mb-0"><i class="fas fa-info-circle me-2"></i>${task.str_creation_date}</p></a>
                                </div>
                            </li>
                        </ul>`
            // <div class="card-body">
            //     <h5 class="card-title">${task.title}</h5>
            //     <p class="card-text">Статус: ${task.status}</p>
            //     <p class="card-text">Дедлайн: ${task.str_creation_date}</p>
            // </div>

        taskList.appendChild(taskElement);
    });
}

// Функция для рендеринга кнопок пагинации
function renderPagination() {
    const pagination = document.getElementById('pagination');
    pagination.innerHTML = ''; // Очистить старые кнопки

    const totalPages = Math.ceil(tasks.length / tasksPerPage);
    console.log(tasks.length, totalPages)
    for (let i = 1; i <= totalPages; i++) {
        const pageItem = document.createElement('li');
        pageItem.className = `page-item ${i === currentPage ? 'active' : ''}`;
        pageItem.innerHTML = `<a class="page-link" href="home/tasks?page=${i}">${i}</a>`;
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
fetch('http://localhost:3000/home/tasks') // Замените на ваш API
    .then(response => response.json())
    .then(data => {
        tasks.push(...data); // Загружаем данные в массив
        renderTasks(); // Рендерим первую страницу задач
        renderPagination(); // Создаем кнопки пагинации
    })
    .catch(error => console.error('Ошибка загрузки задач:', error));
