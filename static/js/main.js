document.getElementById("filter-select").addEventListener("change", function(){
    const selectedValue = this.value;
    console.log(selectedValue)
    fetch(`home/tasks?filter=${selectedValue}`)
        .then(response =>{
            if (!response.ok) {
                throw new Error("response not ok");
            }
            return response.text();
        })
        .then(html =>{
            const tasks = document.getElementById("task-list")
            const pafinationNavBar = document.getElementById("pagination-nav-bar")
            tasks.innerHTML = ''
            pafinationNavBar.innerHTML = ''
            tasks.innerHTML = html
        })

        .catch(err => {
            console.error("error: ", err);
        })
})
document.getElementById("sort-select").addEventListener("change", function(){
    const selectedValue = this.value;
    console.log(selectedValue)
    fetch(`home/tasks?sort=${selectedValue}`)
        .then(response =>{
            if (!response.ok) {
                throw new Error("response not ok");
            }
            return response.text();
        })
        .then(html =>{
            const tasks = document.getElementById("task-list")
            const pafinationNavBar = document.getElementById("pagination-nav-bar")
            tasks.innerHTML = ''
            pafinationNavBar.innerHTML = ''
            tasks.innerHTML = html
        })

        .catch(err => {
            console.error("error: ", err);
        })
})


// const taskForm = document.getElementById('task-form');
//
// // Обработка отправки формы
// taskForm.addEventListener('submit', (event) => {
//     event.preventDefault(); // Отключить стандартное поведение
//     const formData = new FormData(taskForm);
//
//
//     const taskData = {
//         title: formData.get("title"),
//         description: formData.get("description"),
//         due_date: formData.get("due_date"),
//         priority: formData.get("priority"),
//     };
//     console.log(taskData.title, taskData.due_date)
//     fetch('/home/tasks/', {
//         method: 'POST',
//         headers: {
//             'Content-Type': 'application/json',
//         },
//         body: JSON.stringify(taskData),
//     });
//
//
// });





