const post_new_entry = document.getElementById("PostNewEntry");

post_new_entry.onclick = async () => {
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

