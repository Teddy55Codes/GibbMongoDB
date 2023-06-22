document.getElementById("PostNewEntry").onclick = async () => {
  const rawResponse = await fetch("http://127.0.0.1:8080/entries", {
    method: "POST",
    headers: {
      "Accept": "application/json",
      "Content-Type": "application/json"
    },
    body: JSON.stringify(
      {
        name: document.getElementById("nameInput").value, 
        password: document.getElementById("passwordInput").value,
        note: document.getElementById("notesInput").value
      })
  });
  const content = await rawResponse.json();
  
  console.log(content);
};

async function edit(id, name, pass, note) {
  fetch("http://127.0.0.1:8080/entries/" + id, {
    method: "PATCH",
    headers: {
      "Accept": "application/json",
      "Content-Type": "application/json"
    },
    body: JSON.stringify(
        {
          name: name,
          password: pass,
          note: note
        })
  })
      .then(response => {
        getData();
      })
}

async function remove(id) {
  fetch("http://127.0.0.1:8080/entries/" + id, {
    method: "DELETE",
    headers: {
      "Accept": "application/json",
      "Content-Type": "application/json"
    }
  })
      .then(response => {
        getData();
      })
}

document.getElementById("GetEntries").onclick = () => {
  getData();
}

async function getData() {
  fetch("http://127.0.0.1:8080/entries", {
    method: "GET",
    headers: {
      "Accept": "application/json"
    }
  })
      .then(response => response.json())
      .then(data => {
        // Create an array to store the converted data
        let entries = [];

        // For each password document in the response data
        data.passwords.forEach((passwordDoc, index) => {
          // If there is a corresponding notes document
          if (data.notes[index]) {
            // Create an object with the name, password, and note


            let entry = {
              id: passwordDoc._id,
              name: passwordDoc.name,
              password: passwordDoc.password,
              note: data.notes[index].note
            };

            // Add the object to the entries array
            entries.push(entry);
          }
        });

        // Log the entries array
        let entriesContainer = document.getElementById("EntriesContainer");
        entriesContainer.innerHTML = "";


        for (const entry of entries) {
          let div = document.createElement("div");
          div.style.borderStyle = "solid"

          let idTitle = document.createElement("h3");
          idTitle.innerHTML = "Id";
          let id = document.createElement("p")
          id.innerHTML = entry.id;

          let userNameTitle = document.createElement("h3");
          userNameTitle.innerHTML = "Username";
          let userName = document.createElement("input");
          userName.value = entry.name;

          let passwordTitle = document.createElement("h3");
          passwordTitle.innerHTML = "Password";
          let password = document.createElement("input");
          password.value = entry.password;

          let noteTitle = document.createElement("h3");
          noteTitle.innerHTML = "Note"
          let note = document.createElement("input");
          note.value = entry.note

          let editButton = document.createElement("button")
          editButton.innerHTML = "edit";
          editButton.onclick = () => {
            edit(entry.id, userName.value, password.value, note.value);
          }

          let removeButton = document.createElement("button")
          removeButton.innerHTML = "remove";
          removeButton.onclick = () => {
            remove(entry.id);
          }


          div.appendChild(idTitle);
          div.appendChild(id);
          div.appendChild(userNameTitle);
          div.appendChild(userName);
          div.appendChild(passwordTitle);
          div.appendChild(password);
          div.appendChild(noteTitle);
          div.appendChild(note);
          div.appendChild(editButton);
          div.appendChild(removeButton);

          entriesContainer.appendChild(div);
        }
      })
      .catch((error) => {
        console.error('Error:', error);
      });
}

