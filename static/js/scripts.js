// fetches json data from api for warm ups and cooldowns
function fetchWarmUp(level, listId) {
    const url = `/api/warmup?level=${level}`;

    fetch(url)
        .then(response => {
            if (!response.ok) {
                throw new Error(`HTTP error! status: ${response.status}`);
            }
            return response.json();
        })
        .then(data => {
            const listElement = document.querySelector(`#${listId}`);
            if (!listElement) {
                console.error(`Element with ID '${listId}' not found.`);
                return;
            }
            listElement.innerHTML = ''; // Clear previous exercises

            // Append each new exercise to the list
            data.forEach(item => {
                const listItem = document.createElement('li');
                listItem.textContent = item.exercise_name;
                listElement.appendChild(listItem);
            });
        })
        .catch(error => {
            console.error('Error fetching exercises:', error);
            alert('Failed to load exercises. Please try again later.');
        });
}

// fetches json data from api for warm ups and cooldowns
function fetchCooldown(level, listId) {
    const url = `/api/warmup?level=${level}`;

    fetch(url)
        .then(response => {
            if (!response.ok) {
                throw new Error(`HTTP error! status: ${response.status}`);
            }
            return response.json();
        })
        .then(data => {
            const listElement = document.querySelector(`#${listId}`);
            if (!listElement) {
                console.error(`Element with ID '${listId}' not found.`);
                return;
            }
            listElement.innerHTML = ''; // Clear previous exercises

            // Append each new exercise to the list
            data.forEach(item => {
                const listItem = document.createElement('li');
                listItem.textContent = item.exercise_name;
                listElement.appendChild(listItem);
            });
        })
        .catch(error => {
            console.error('Error fetching exercises:', error);
            alert('Failed to load exercises. Please try again later.');
        });
}

function trueBeginnerWarmUp() {
    fetchWarmUp(1, 'warm-up-list');
}

function trueBeginnerCoolDown() {
    fetchCooldown(1, 'coolDownList');
}

function fetchMainExercise(level, listId, name) {
    const url = `/api/exercise?level=${level}&name=${name}`;

    fetch(url)
        .then(response => {
            if (!response.ok) {
                throw new Error(`HTTP error! status: ${response.status}`);
            }
            return response.json();
        })
        .then(data => {
            const listElement = document.querySelector(`#${listId}`);
            if (!listElement) {
                console.error(`Element with ID '${listId}' not found.`);
                return;
            }
            listElement.innerHTML = ''; // Clear previous exercises

            // Handle the single object response
            const listItem = document.createElement('li');
            listItem.className = 'selectedExercise';
            listItem.textContent = data.exercise_name; // Access the exercise name
            listElement.appendChild(listItem);
        })
        .catch(error => {
            console.error('Error fetching exercises:', error);
            alert('Failed to load exercises. Please try again later.');
        });
}

function beginnerChest() {
    fetchMainExercise(1, 'exerciseOption', 'chest');
}
function beginnerBack() {
    fetchMainExercise(1, 'exerciseOption', 'back');
}
function beginnerLegs() {
    fetchMainExercise(1, 'exerciseOption', 'legs');
}
function beginnerShoulders() {
    fetchMainExercise(1, 'exerciseOption', 'shoulders');
}
function beginnerRearShoulders() {
    fetchMainExercise(1, 'exerciseOption', 'rear_shoulders');
}
function beginnerLowerBack() {
    fetchMainExercise(1, 'exerciseOption', 'lower_back');
}
function beginnerCore() {
    fetchMainExercise(1, 'exerciseOption', 'core');
}
function beginnerBalance() {
    fetchMainExercise(1, 'exerciseOption','one_leg_stable')
}

function submitToList() {
    const addedExercise = document.getElementsByClassName('selectedExercise');
    let addExercise = addedExercise[0].innerHTML
    

    const list = document.querySelector('#selected-exercises');
    const listItem = document.createElement('li');

    listItem.textContent = addExercise; 
    list.appendChild(listItem);  
}

function customExercise(){
    const customExercise = document.getElementById('exercise-name').value;
    const list = document.querySelector('#selected-exercises');
    const listItem = document.createElement('li');
    listItem.textContent = customExercise
    list.appendChild(listItem)
}