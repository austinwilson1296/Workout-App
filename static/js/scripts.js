function trueBeginnerWarmUp(){
    const url = "/api/warmup?level=1"

    fetch(url)
    .then(response => response.json())
    .then(data => {
        const warmUpList = document.querySelector('#warm-up-list');
        warmUpList.innerHTML = '';  // Clear previous exercises

        // Append each new exercise to the list
        data.forEach(item => {
            const listItem = document.createElement('li');
            listItem.textContent = item.exercise_name;
            warmUpList.appendChild(listItem);
        });
    })
    .catch(error => console.error('Error fetching exercises:', error));
}
