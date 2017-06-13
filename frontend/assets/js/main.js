;! function () {
   window.onload = function () {
        var baseUrl = [
            "http://localhost:8000/search?query=",
            "http://localhost:8080/test_full_stack.php?query="
        ];

        const ctaButton = document.getElementById("submit");
        const queryInputs = document.getElementsByName("query");

        if (ctaButton && queryInputs && queryInputs[0]) {
            ctaButton.addEventListener("click", event =>
            fetch(baseUrl[0] + queryInputs[0].value, getRequestParameters(queryInputs[0]))
                .then(response => response.json())
        .then(handleResponse)
                .catch(handleError)
        );
        }
   };
}();

function getRequestParameters(input) {
    return {
        method: "GET"
    };
}

function handleResponse(response) {
    html = '';
    for (let index in response.tours) {
        response.tours[index].title = convert(response.tours[index].title);
        html += `<article id="${response.tours[index].id}"><h2>${response.tours[index].title}</h2><strong>${response.tours[index].currency} ${response.tours[index].price}</strong><br><mark>${response.tours[index].rating}</mark></article>`;
    }

    const resultsSection = document.getElementById("results");
    if (results) {
        results.innerHTML = html;
    }
}

function handleError(error) {
    console.error('error', error);
}

function convert(str){
    return str.replace(/\\u([a-f0-9]{4})/gi, function (found, code) {
        return String.fromCharCode(parseInt(code, 16));
    });
}