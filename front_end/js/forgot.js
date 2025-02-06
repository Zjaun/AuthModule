document.getElementById("form").addEventListener("submit", async function(event) {
    event.preventDefault();

    const email = document.getElementById("email").value;

    const button = document.getElementById("submit");
    button.disabled = true;


})