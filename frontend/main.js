const inputField = document.getElementById("url-input");
const shortenedUrl = document.getElementById("shorten-url");
const submitBtn = document.getElementById("submit-btn");

async function handleSubmit(event) {
    event.preventDefault();

    const url = inputField.value;

    const res = await fetch("http://localhost:8080/api/create", {
        method: 'POST',
        mode: "no-cors",
        body: JSON.stringify({
            url: url,
        }),
        headers: {
            'Content-type': 'application/json',
            'Access-Control-Allow-Origin': '*',
        }
    });

    const data = await res.json();
    const shortenUrl = data.url;

    shortenedUrl.innerText = shortenUrl;
}

submitBtn.addEventListener("click", handleSubmit);
