export function disableForm(form) {
    form.classList.add("disabled");
    form.disabled = true;
}

export function enableForm(form) {
    form.classList.remove("disabled");
    form.disabled = false;
}

export function isBlank(string) {
    return !string.trim();
}

export function showError(element, error) {
    element.innerText = error;
}

export function showDivError(element) {
    element.classList.add("error");
}

export function clearError(element) {
    element.innerText = "";
}

export function clearDivError(element) {
    element.classList.remove("error");
}

export function isHidden(div) {
    return div.classList.contains("hidden");
}

export function showDiv(div) {
    div.classList.remove("hidden");
}

export function hideDiv(div) {
    div.classList.add("hidden");
}

export function displayBlock(div) {
    div.style.display = "block";
}

export function displayNone(div) {
    div.style.display = "none";
}

export const delay = ms => new Promise(res => setTimeout(res, ms));