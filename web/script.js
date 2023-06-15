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

document.getElementById("GetEntries").onclick = async () => {
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
          name: passwordDoc["1"].name,
          password: passwordDoc["1"].password,
          note: data.notes[index]["1"].note
        };

        // Add the object to the entries array
        entries.push(entry);
      }
    });

    // Log the entries array
    console.log(entries);
  })
  .catch((error) => {
    console.error('Error:', error);
  });
};

