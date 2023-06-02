const inputField = document.getElementById("url-input");
const shortenedUrl = document.getElementById("shorten-url");
const submitBtn = document.getElementById("submit-btn");

function handleSubmit(event) {
    event.preventDefault();

    const url = inputField.value;
    const shortenUrl = "";

    fetch("http://localhost:8080/api/create", {
        method: 'POST',
        mode: "no-cors",
        body: JSON.stringify({
            url: url,
        }),
        headers: {
            'Content-type': 'application/json; charset=UTF-8',
            'Access-Control-Allow-Origin': '*',
        }
    })
    .then(res => {return res.json(); }) 
    .then(data => {
        shortenUrl = data.url;
    })
    .catch(error => console.error('Error:', error)); 
    shortenedUrl.innerText = shortenUrl;
}

submitBtn.addEventListener("click", handleSubmit);
