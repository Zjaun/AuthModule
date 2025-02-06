export function disableForm(form) {
    form.classList.add("disabled");
    form.disabled = true;
}

export function enableForm(form) {
    form.classList.remove("disabled");
    form.disabled = false;
}