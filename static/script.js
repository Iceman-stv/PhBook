function loadContacts() {
    const userID = document.getElementById('userID').value;
    fetch(`/api/contacts?user_id=${userID}`)
        .then(response => response.json())
        .then(contacts => {
            const contactsList = document.getElementById('contacts');
            contactsList.innerHTML = '';
            contacts.forEach(contact => {
                const li = document.createElement('li');
                li.textContent = `${contact.Name}: ${contact.Phone}`;
                contactsList.appendChild(li);
            });
        })
        .catch(error => console.error('Ошибка:', error));
}