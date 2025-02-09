import {clearError, clearDivError, isBlank, showError, showDivError} from "./utility.js";

// for showing errors
const errs = {
    "errName": document.getElementById("errName"),
    "errUser": document.getElementById("errUser"),
    "errEmail": document.getElementById("errEmail"),
    "errPassword": document.getElementById("errPassword"),
    "errConfirm": document.getElementById("errConfirm"),
    "errQues1": document.getElementById("errQues1"),
    "errQues2": document.getElementById("errQues2"),
    "errQues3": document.getElementById("errQues3"),
    "errAns1": document.getElementById("errAns1"),
    "errAns2": document.getElementById("errAns2"),
    "errAns3": document.getElementById("errAns3"),
    "errErr": document.getElementById("err")
};

// for input fields
const fields = {
    "First": document.getElementById("firstname"),
    "Last": document.getElementById("lastname"),
    "User": document.getElementById("username"),
    "Email": document.getElementById("email"),
    "Password": document.getElementById("password"),
    "Confirm": document.getElementById("confirmPassword"),
    "Ques1": document.getElementById("ques1"),
    "Ques2": document.getElementById("ques2"),
    "Ques3": document.getElementById("ques3"),
    "Ans1": document.getElementById("ques1Ans"),
    "Ans2": document.getElementById("ques2Ans"),
    "Ans3": document.getElementById("ques3Ans"),
};

// for making input field border glow red
const divs = {
    "divName": document.getElementById("divFirst"),    // 0
    "divLast": document.getElementById("divLast"),     // 1
    "divUser": document.getElementById("divUser"),     // 2
    "divEmail": document.getElementById("divEmail"),    // 3
    "divPassword": document.getElementById("divPassword"), // 4
    "divConfirm": document.getElementById("divConfirm"),  // 5
    "divQues1": document.getElementById("divQues1"),    // 6
    "divQues2": document.getElementById("divQues2"),    // 7
    "divQues3": document.getElementById("divQues3"),    // 8
    "divAns1": document.getElementById("divAns1"),     // 9
    "divAns2": document.getElementById("divAns2"),     // 10
    "divAns3": document.getElementById("divAns3")      // 11
};

const button = document.getElementById("button")
const form = document.getElementById("form");

const questions = [
    "What was the name of your first pet?",
    "What is your favorite book?",
    "In which city were you born?",
    "What is your favorite childhood movie?",
    "Who was your childhood best friend?",
    "What is your favorite food?",
    "What was the model of your first car?",
    "What is your mother's middle name?",
    "Where did you go for your first vacation?",
    "What is your favorite sports team?",
    "What is your favorite color?",
    "What is your favorite historical event?",
    "What is your favorite holiday destination?",
    "What is your favorite song?",
    "Who is your favorite teacher?",
    "What is the first name of your childhood hero?",
    "What is the name of your favorite fictional character?",
    "What is your favorite hobby?",
    "What is the name of your first employer?",
    "What is your favorite quote?",
    "What was the name of your first stuffed animal?",
    "What is the name of the street you grew up on?",
    "What was your favorite game as a child?",
    "What was the first concert you attended?",
    "What is the name of a childhood neighbor?",
    "What is the name of your first teacher?",
    "What was the first thing you learned to cook?",
    "What is the name of a place you always wanted to visit but haven't yet?",
    "What was your favorite subject in school?",
    "What is your least favorite vegetable?",
    "What is the title of the first movie you watched in a theater?",
    "What was your childhood dream job?",
    "What is the name of your first crush?",
    "What was the make of your first bicycle?",
    "What is your favorite homemade dish?",
    "What is the name of your first friend in college?",
    "What is the name of a park you visited often as a child?",
    "What was the first foreign language you learned?",
    "What is the name of your favorite high school teacher?",
    "What was your favorite TV show growing up?"
]

function generateOptions(value, question) {
    const option1 = document.createElement("option");
    const option2 = document.createElement("option");
    const option3 = document.createElement("option");
    option1.value = value.toString();
    option1.text = question;
    option2.value = value.toString();
    option2.text = question;
    option3.value = value.toString();
    option3.text = question;
    return [option1, option2, option3]
}

function populateQuestions() {
    for (let i = 0; i < questions.length; i++) {
        const options = generateOptions(i, questions[i]);
        fields["Ques1"].appendChild(options[0]);
        fields["Ques2"].appendChild(options[1]);
        fields["Ques3"].appendChild(options[2]);
    }
}

function clearSuggestions() {
    for (const err in errs) {
        clearError(errs[err]);
    }
    for (const div in divs) {
        clearDivError(divs[div]);
    }
}

function showErrSame(index1, index2, message) {
    showDivError(divs["div" + index1]);
    showDivError(divs["div" + index2]);
    showError(errs["err" + index1], message);
    showError(errs["err" + index2], message);
}

function hasSameQuestions() {
    let message = "Cannot be the same.";
    if (fields["Ques1"].value === fields["Ques2"].value) {
        showErrSame("Ques1", "Ques2", message);
        return true;
    } else if (fields["Ques2"].value === fields["Ques3"].value) {
        showErrSame("Ques2", "Ques3", message);
        return true;
    } else if (fields["Ques1"].value === fields["Ques3"].value) {
        showErrSame("Ques1", "Ques3", message);
        return true;
    }
    return false;
}

function samePassword() {
    if (fields["Password"].value !== fields["Confirm"].value) {
        showErrSame("Password", "Confirm", "Passwords do not match.");
        return false;
    }
    return true;
}

async function register() {
    try {
        const body = JSON.stringify({
            username: fields["User"].value,
            first: fields["First"].value,
            last: fields["Last"].value,
            email: fields["Email"].value,
            password: fields["Password"].value,
            q1: fields["Ques1"].value,
            q1Ans: fields["Ans1"].value,
            q2: fields["Ques2"].value,
            q2Ans: fields["Ans2"].value,
            q3: fields["Ques3"].value,
            q3Ans: fields["Ans3"].value
        })
        const response = await fetch("http://localhost:8080/validate", {
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
        if (data["valid"]) {
            alert("Success!");
        } else {
            let errField = data["field"];
            if (errField === "Err") {
                showError(errs["err" + errField], "Please contact developer: " + data["reason"]);
            } else {
                showDivError(divs["div" + errField]);
                showError(errs["err" + errField], data["reason"]);
            }
        }
    } catch (error) {
        console.log(error);
        alert(error);
    }
}

function hasErrors() {
    for (const field in fields) {
        if (field === "Ques1" || field === "Ques2" || field === "Ques3") {
            continue;
        }
        let hasError = false;
        if (isBlank(fields[field].value)) {
            showDivError(divs["div" + field]);
            hasError = true;
        }
        if (hasError) {
            if (field === "First" || field === "Last") {
                if (isBlank(fields[field].value)) {
                    showError(errs["errName"], "Cannot be empty.");
                }
            } else {
                showError(errs["err" + field], "Cannot be empty.");
            }
            return true;
        }
    }
    return !samePassword() || hasSameQuestions();
}

async function submit(event) {
    event.preventDefault()
    clearSuggestions();
    if (!hasErrors()) {
        await register();
    }
}

form.addEventListener("submit", submit);
populateQuestions();