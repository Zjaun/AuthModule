import {disableForm} from "./forms.js";

const emailField = document.getElementById("email");
const passwordField = document.getElementById("password");
const button = document.getElementById("submit");

function checkFields() {
    button.disabled = !(emailField.value.trim() !== "" && passwordField.value.trim() !== "");
}

emailField.addEventListener("input", checkFields);
passwordField.addEventListener("input", checkFields);

document.getElementById("form").addEventListener("submit", async function(event) {
    event.preventDefault();
    const email = document.getElementById("email").value;
    const password = document.getElementById("password").value;
    disableForm(emailField);
    disableForm(passwordField);
    button.disabled = true;
    button.innerText = "Signing in...";
})