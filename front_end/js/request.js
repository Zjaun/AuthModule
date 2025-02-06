document.getElementById("form").addEventListener("submit", async function(event) {
    event.preventDefault();

    const first = document.getElementById("first_name").value;
    const last = document.getElementById("last_name").value;

    const response = await fetch ("http://localhost:8080/validate", {
        method: "POST",
        headers: {
            "Content-Type": "application/json"
        },
        body: JSON.stringify({first, last})
    });

    const result = await response.json();
    document.getElementById("response").innerText = result.valid ?
        "Input looks good." : result.reason;
})