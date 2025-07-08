const getUser = () => {
    return new Promise((resolve, reject) => {
        const request = new XMLHttpRequest();

        request.onreadystatechange = () => {
            if (request.readyState === 4) {
                if (request.status === 200) {
                    try {
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

        request.open("GET", "http://localhost:8080/albums", true);
        request.send();
    });
};
