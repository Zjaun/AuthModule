import {clearDivError, clearError, delay, hideDiv, isHidden, showDiv, showDivError, showError} from "./utility.js";

const divs = {
    "divUser": document.getElementById("divUser"),
    "divQues": document.getElementById("divQues"),
    "divAns": document.getElementById("divAns"),
    "divPass": document.getElementById("divPass"),
    "divConfirm": document.getElementById("divConfirm")
}
const fields = {
    "Username": document.getElementById("username"),
    "Ques": document.getElementById("ques"),
    "Answer": document.getElementById("answer"),
    "Password": document.getElementById("password"),
    "Confirm": document.getElementById("confirm"),
}
const textBox = {
    "errUser": document.getElementById("errUser"),
    "errAns": document.getElementById("errAns"),
    "errPass": document.getElementById("errPass"),
    "errConfirm": document.getElementById("errConfirm"),
    "errErr": document.getElementById("errErr"),
    "message": document.getElementById("message"),
    "message2": document.getElementById("message2")
};

const button = document.getElementById("submit");

function checkFields() {
    button.disabled = fields["Username"].value.trim() === "";
}

function generateQuestions(q1, q2, q3) {
    const option1 = document.createElement("option");
    const option2 = document.createElement("option");
    const option3 = document.createElement("option");
    option1.value = "sq1";
    option1.text = q1;
    option2.value = "sq2";
    option2.text = q2;
    option3.value = "sq3";
    option3.text = q3;
    fields["Ques"].appendChild(option1);
    fields["Ques"].appendChild(option2);
    fields["Ques"].appendChild(option3);
}

async function submit() {

    // if (!isHidden(divs["divQues"])) {
    //     if (fields["Answer"].value === "") {
    //         showDivError(divs["divAns"]);
    //         showError(textBox["errAns"], "Cannot be empty.");
    //         return;
    //     } else if (fields["Password"].value === "") {
    //         showDivError(divs["divPass"]);
    //         showError(textBox["errPass"], "Cannot be empty.");
    //         return;
    //     } else if (fields["Confirm"].value === "") {
    //         showDivError(divs["divConfirm"]);
    //         showError(textBox["errConfirm"], "Cannot be empty.");
    //         return;
    //     } else if (fields["Password"].value !== fields["Confirm"].value) {
    //         showDivError(divs["divPass"]);
    //         showDivError(divs["divConfirm"]);
    //         showError(textBox["errConfirm"], "Passwords do not match.");
    //         return;
    //     }
    // }

    let body;
    if (isHidden(divs["divQues"])) {
        body = JSON.stringify({
            "valUser": true,
            "username": fields["Username"].value,
            "valQues": false,
            "question": "",
            "answer": "",
            "password": ""
        })
    } else {
        body = JSON.stringify({
            "valUser": true,
            "username": fields["Username"].value,
            "valQues": true,
            "question": fields["Ques"].value,
            "answer": fields["Answer"].value,
            "password": fields["Password"].value
        })
    }

    try {
        const response = await fetch("http://localhost:8080/reset", {
            method: "POST",
            headers: {
                "Content-Type": "application/json"
            },
            body
        });
        if (!response.ok) {
            const errorData = await response.json();
            const errorMessage = errorData.message || "Registration Failed.";
            throw new Error(errorMessage);
        }
        const data = await response.json();
        if (data["field"] === "Err") {
            showError(textBox["errErr"], "Please contact developer: " + data["reason"]);
            return;
        }
        if (data["success"]) {
            hideDiv(divs["divQues"])
            hideDiv(divs["divAns"])
            hideDiv(divs["divPass"]);
            hideDiv(divs["divConfirm"]);
            textBox["message"].innerText = "Successfully Changed Password.";
            textBox["message2"].innerText = "Redirecting to login page...";
            await delay(2500);
            window.location.replace("http://localhost:8080/Login.html");
        } else {
            if (data["field"] === "User") {
                hideDiv(divs["divQues"])
                hideDiv(divs["divAns"])
                hideDiv(divs["divPass"]);
                hideDiv(divs["divConfirm"]);
                showDivError(divs["divUser"]);
                showError(textBox["errUser"], data["reason"]);
                return;
            }
            if (data["field"] === "Answer") {
                showDivError(divs["divAns"]);
                showError(textBox["errAns"], data["reason"]);
            } else if (data["q1"] !== "") {
                showDiv(divs["divQues"]);
                showDiv(divs["divAns"]);
                showDiv(divs["divPass"]);
                showDiv(divs["divConfirm"]);
                generateQuestions(data["q1"], data["q2"], data["q3"]);
            } else if (fields["Answer"].value === "") {
                showDivError(divs["divAns"]);
                showError(textBox["errAns"], "Cannot be empty.");
                return;
            } else if (fields["Password"].value === "") {
                showDivError(divs["divPass"]);
                showError(textBox["errPass"], "Cannot be empty.");
                return;
            } else if (fields["Confirm"].value === "") {
                showDivError(divs["divConfirm"]);
                showError(textBox["errConfirm"], "Cannot be empty.");
                return;
            } else if (fields["Password"].value !== fields["Confirm"].value) {
                showDivError(divs["divPass"]);
                showDivError(divs["divConfirm"]);
                showError(textBox["errConfirm"], "Passwords do not match.");
                return;
            }
        }
    } catch (error) {
        console.log(error);
        alert(error);
    }
}

function clearSuggestions() {
    for (const err in textBox) {
        clearError(textBox[err]);
    }
    for (const div in divs) {
        clearDivError(divs[div]);
    }
}

button.disabled = true;
fields["Username"].addEventListener("input", checkFields);

document.getElementById("form").addEventListener("submit", async function(event) {
    event.preventDefault();
    clearSuggestions();
    await submit();
})