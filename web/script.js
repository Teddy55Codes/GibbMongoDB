const post_new_entry = document.getElementById("PostNewEntry");

post_new_entry.onclick = async () => {
  const rawResponse = await fetch("https://127.0.0.1:8080/entries", {
    method: "POST",
    headers: {
      "Accept": "application/json",
      "Content-Type": "application/json"
    },
    body: JSON.stringify(
      {
        name: document.getElementById("nameInput"), 
        password: document.getElementById("passwordInput"),
        note: document.getElementById("notesInput")
      })
  });
  const content = await rawResponse.json();
  
  console.log(content);
};

