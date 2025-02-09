import {clearDivError, clearError, disableForm, enableForm, showDivError, showError} from "./utility.js";

const usernameField = document.getElementById("username");
const passwordField = document.getElementById("password");
const button = document.getElementById("submit");
const divs = {
    "divUser": document.getElementById("divUser"),
    "divPassword": document.getElementById("divPassword")
}
const textBox = {
    "errErr": document.getElementById("errErr"),
    "message": document.getElementById("message")
};

function clearSuggestions() {
    for (const err in textBox) {
        clearError(textBox[err]);
    }
    for (const div in divs) {
        clearDivError(divs[div]);
    }
}

function checkFields() {
    button.disabled = !(usernameField.value.trim() !== "" && passwordField.value.trim() !== "");
}

async function authenticate() {
    const username = document.getElementById("username").value;
    const password = document.getElementById("password").value;
    try {
        const response = await fetch("http://localhost:8080/authenticate", {
            method: "POST",
            headers: {
                "Content-Type": "application/json"
            },
            body: JSON.stringify({username, password})
        });
        if (!response.ok) {
            const errorData = await response.json();
            const errorMessage = errorData.message || "Login failed.";
            throw new Error(errorMessage);
        }
        const data = await response.json();
        if (!data["logged_in"]) {
            showDivError(divs["divUser"]);
            showDivError(divs["divPassword"]);
            showError(textBox["errErr"], data["reason"]);
        } else {
            showError(textBox["message"], "Successfully logged in!")
        }
    } catch (error) {
        alert(error.message);
    } finally {
        enableForm(usernameField);
        enableForm(passwordField);
        button.innerText = "Sign In";
    }
}

usernameField.addEventListener("input", checkFields);
passwordField.addEventListener("input", checkFields);

document.getElementById("form").addEventListener("submit", async function(event) {
    event.preventDefault();

    clearSuggestions();
    disableForm(usernameField);
    disableForm(passwordField);
    button.disabled = true;
    button.innerText = "Signing in...";

    await authenticate()

})