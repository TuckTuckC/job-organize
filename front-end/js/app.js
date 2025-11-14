const getUser = () => {
    return new Promise((resolve, reject) => {
        const request = new XMLHttpRequest();

        request.onreadystatechange = () => {
            if (request.readyState === 4) {
                console.log(request.status)
                if (request.status === 200) {
                    try {
                        console.log(request)
                        const data = JSON.parse(request.responseText);

                        console.log(data);

                        resolve(data);
                    } catch (e) {
                        reject("Failed to parse response JSON");
                    };
                } else {
                    reject(`Request failed with status ${request.status}`);
                };
            };
        };

        request.open("GET", "http://localhost:8080/users", true);

        request.send();
    });
};

const createUser = (email, firstName, lastName) => {
    return new Promise((resolve, reject) => {
        const request = new XMLHttpRequest();

        request.onreadystatechange = () => {
            if (request.readyState === 4) {
                if (request.status === 201) {
                    try {
                        const data = JSON.parse(request.responseText);
                        console.log("User created:", data);
                        resolve(data);
                    } catch (e) {
                        reject("Failed to parse response JSON");
                    }
                } else {
                    reject(`Request failed with status ${request.status}`);
                }
            }
        };

        request.open("POST", "http://localhost:8080/user", true);
        request.setRequestHeader("Content-Type", "application/json");
        request.send(JSON.stringify({
            email: email,
            firstName: firstName,
            lastName: lastName
        }));
    });
};

const createTimesheetEntry = (userID, startDateTime, hours) => {
    return new Promise((resolve, reject) => {
        const request = new XMLHttpRequest();

        request.onreadystatechange = () => {
            if (request.readyState === 4) {
                if (request.status === 201) {
                    try {
                        const data = JSON.parse(request.responseText);
                        console.log("Timesheet Entry created:", data);
                        resolve(data);
                    } catch (e) {
                        reject("Failed to parse response JSON");
                    }
                } else {
                    reject(`Request failed with status ${request.status}`);
                }
            }
        };

        request.open("POST", "http://localhost:8080/timesheetEntry", true);
        request.setRequestHeader("Content-Type", "application/json");
        request.send(JSON.stringify({
            userID: userID,
            startDateTime: startDateTime,
            hours: hours
        }));
    });
};


document.getElementById("userForm").addEventListener("submit", function (e) {
    e.preventDefault();

    const email = document.getElementById("email").value;
    const firstName = document.getElementById("firstName").value;
    const lastName = document.getElementById("lastName").value;

    createUser(email, firstName, lastName)
        .then(user => {
            getUser(); // refresh list in console
        })
        .catch(err => {
            console.error(err);
            alert("Failed to create user.");
        });
});

document.getElementById("timesheetForm").addEventListener("submit", function (e) {
    e.preventDefault();

    const userID = document.getElementById("userID").value;
    const startDateTime = document.getElementById("startDateTime").value;
    const hours = document.getElementById("hours").value;

    createTimesheetEntry(userID, startDateTime, hours)
        .then(timesheetEntry => {
            getUser(); // refresh list in console
        })
        .catch(err => {
            console.error(err);
            console.log(userID, startDateTime, hours)
            alert("Failed to create timesheet entry.");
        });
});
