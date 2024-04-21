const loginForm = document.querySelector('#loginForm')
const email = document.querySelector('#email')
const password = document.querySelector('#password')

loginForm.addEventListener('submit', (event) => {
    console.log("Entered Login form Submit action")
    event.preventDefault()
    const isValid = validateLoginInputs()
    if (isValid) {
        console.log("All fields are validated")
        const payload = {
            action: "login",
            login: {
                email: email.value.trim(),
                password: password.value.trim()
            }
        };

        const options ={
            method: "POST",
            body: JSON.stringify(payload),
            headers: {
                "Content-Type": "application/json"
            }
        };

        fetch("http://localhost:8080/handle", options)
            .then(response => response.json())
            .then(data => {
                window.location.href = "/home";
                console.log(data)
            })
            .catch(err => {
                console.log(err)
            })
    }
});

function validateLoginInputs() {
    const emailVal = email.value.trim();
    const passwordVal = password.value.trim();
    console.log("Email - ", emailVal, "\nPassword - ", passwordVal)
    return validateLoginEmailAndPassword(emailVal, passwordVal)
}

function validateLoginEmailAndPassword (loginEmail, passwordVal) {
    if (loginEmail === '') {
        setError(email, 'Email is required');
        return false;
    } else if (!validateEmailFormat(loginEmail)) {
        setError(email, 'Please enter a valid email');
        return false;
    } else {
        setSuccess(email);
    }

    if (passwordVal === '') {
        setError(password, 'Password is required');
        return false;
    } else if (passwordVal.length < 8) {
        setError(password, 'Password must be at least 8 characters long');
        return false;
    } else {
        setSuccess(password);
    }

    return true
}

const validateEmailFormat = (email) => {
    return String(email)
        .toLowerCase()
        .match(
            /^(([^<>()[\]\\.,;:\s@"]+(\.[^<>()[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/
        );
};

function setError(element, message) {
    const inputGroup = element.parentElement;
    const errorElement = inputGroup.querySelector('.error');

    errorElement.innerText = message;
    inputGroup.classList.add('error');
    inputGroup.classList.remove('success');
}

function setSuccess(element) {
    const inputGroup = element.parentElement;
    const errorElement = inputGroup.querySelector('.error');

    errorElement.innerText = '';
    inputGroup.classList.add('success');
    inputGroup.classList.remove('error');
}