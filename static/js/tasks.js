fetch("GET /home/tasks").
    then(response =>{
        if (!response.ok) {
            throw new Error('Ошибка загрузки данных');
        }
})