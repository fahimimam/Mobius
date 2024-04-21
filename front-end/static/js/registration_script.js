// Get the form element
const registerForm = document.querySelector('#registerForm');
const firstname = document.querySelector('#firstname');
const lastname = document.querySelector('#lastname');
const email = document.querySelector('#email');
const password = document.querySelector('#password');
const confirmPassword = document.querySelector('#confirmPassword');

// Add event listener to form submit
registerForm.addEventListener('submit', (e) => {
    // Prevent the default form submission behavior
    e.preventDefault();

    // Validate form inputs
    const isValid = validateRegistrationInputs();

    console.log("email - ", email.value.trim())
    console.log("firstname - ", firstname.value.trim())
    console.log("lastname - ", lastname.value.trim())
    console.log("password - ", password.value.trim())

    // If form inputs are valid, submit the form
    if (isValid) {
        // Construct the payload
        const payload = {
            action: "registration",
            registration: {
                email: email.value.trim(),
                firstname: firstname.value.trim(),
                lastname: lastname.value.trim(),
                password: password.value.trim()
            }
        };
        console.log("Prepared Pld - ", payload)
        // Construct the fetch options
        const options = {
            method: "POST",
            body: JSON.stringify(payload),
            headers: {
                "Content-Type": "application/json"
            }
        };

        // Send the form data to the broker microservice
        fetch("http://localhost:8080/handle", options)
            .then(response => response.json())
            .then(data => {
                // Handle the response
                window.location.href = "/home";
                console.log(data);
            })
            .catch(error => {
                // Handle errors
                console.log(error);
            });
    }
});

function validateRegistrationInputs() {
    const firstnameVal = firstname.value.trim();
    const lastnameVal = lastname.value.trim();
    const passwordVal = password.value.trim();
    const confirmPasswordVal = confirmPassword.value.trim();
    const emailVal = email.value.trim();
    let success = true;

    if (firstnameVal === '') {
        success = false;
        console.log("No First name provided: ");
        setError(firstname, 'Firstname is required');
    } else {
        console.log("Firstname: ", firstnameVal);
        setSuccess(firstname);
    }

    if (lastnameVal === '') {
        success = false;
        setError(lastname, 'Lastname is required');
    } else {
        setSuccess(lastname);
    }

    success = validateRegistrationEmailAndPassword(emailVal, passwordVal)

    if (confirmPasswordVal === '') {
        success = false;
        setError(confirmPassword, 'Confirm password is required');
    } else if (confirmPasswordVal !== passwordVal) {
        success = false;
        setError(confirmPassword, 'Password does not match');
    } else {
        setSuccess(confirmPassword);
    }

    return success;
}

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

const validateEmailFormat = (email) => {
    return String(email)
        .toLowerCase()
        .match(
            /^(([^<>()[\]\\.,;:\s@"]+(\.[^<>()[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/
        );
};

function validateRegistrationEmailAndPassword (registrationEmail, passwordVal) {
    if (registrationEmail === '') {
        setError(email, 'Email is required');
        return false;
    } else if (!validateEmailFormat(registrationEmail)) {
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
}