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
