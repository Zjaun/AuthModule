import {disableForm, enableForm} from "./forms";

const month = document.getElementById("month");
const day = document.getElementById("day");
const year = document.getElementById("year");
const months = [
    "January", "February", "March", "April", "May", "June",
    "July", "August", "September", "October", "November", "December"
]
const days = [31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31]
const date = new Date();

function isLeapYear(year) {
    year = Number(year);
    return year % 400 === 0 || (year % 4 === 0 && year % 100 !== 0);
}

function isCurrentYear(year) {
    return Number(year) === date.getFullYear();
}

function generateDays(month, year) {
    day.innerHTML = "";
    let bound = days[Number(month) - 1] + 1
    if (isCurrentYear(year)) {
        bound = date.getDate();
    } else if (month.toString() === "2" && isLeapYear(year)){
        bound++;
    }
    for (let i = 1; i < bound; i++) {
        let option = document.createElement("option");
        option.value = i.toString();
        option.text = i.toString();
        day.appendChild(option);
    }
}

function generateMonths(year) {
    let bound;
    if (isCurrentYear(year)) {
        bound = date.getMonth() + 2;
    } else {
        bound = months.length;
    }
    for (let i = 1; i < bound; i++) {
        let option = document.createElement("option");
        option.value = i.toString();
        option.text = months[i - 1];
        month.appendChild(option);
    }
}

function generateYears() {
    let bound = date.getFullYear() + 1;
    for (let i = 1900; i < bound; i++) {
        let option = document.createElement("option");
        option.value = i.toString();
        option.text = i.toString();
        year.appendChild(option);
    }
}

month.addEventListener("change", function() {
    const selectedMonth = month.value;
    const selectedYear = year.value;
    generateDays(selectedMonth, selectedYear);
    if (day.disabled) {
        enableForm(day)
    }
});

year.addEventListener("change", function() {
    if (month.disabled) {
        enableForm(month);
        generateMonths(year.value);
    }
    if (month.value !== "" || month.value !== null) {
        generateDays(month.value, year.value);
        enableForm(days);
    }
});

generateYears();
disableForm(month);
disableForm(day);