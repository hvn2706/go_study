document
    .getElementById("get-album-by-id")
    .addEventListener("click", function () {
        id = document.getElementById("album-id-get").value;

        if (id == "") {
            alert("Please enter an album id");
            return;
        }

        fetch(`http://localhost:8080/albums/${id}/`, {
            mode: "cors",
            method: "GET",
            headers: {
                "Access-Control-Allow-Origin": "*",
                "Content-Type": "application/json",
            },
        })
            .then(function (response) {
                console.log(response);
                return response.json();
            })
            .then(function (data) {
                console.log(data);
                document.getElementById("album-by-id-result").innerHTML =
                    JSON.stringify(data);
            });
    });

document.getElementById("new-album").addEventListener("click", function () {
    var album = {
        title: document.getElementById("album-title").value,
        artist: document.getElementById("album-artist").value,
        price: parseFloat(document.getElementById("album-price").value),
    };

    if (album.title == "" || album.artist == "" || album.price == "") {
        alert("Please enter all album details");
        return;
    }

    fetch("http://localhost:8080/albums/", {
        method: "POST",
        mode: "cors",
        headers: {
            "Access-Control-Allow-Origin": "*",
            "Content-Type": "application/json",
        },
        body: JSON.stringify(album),
    })
        .then(function (response) {
            console.log(response);
            return response.json();
        })
        .then(function (data) {
            console.log(data);
            document.getElementById("new-album-result").innerHTML =
                JSON.stringify(data);
        });
});

document.getElementById("update-album").addEventListener("click", function () {
    var album = {
        ID: parseInt(document.getElementById("album-id-update").value),
        title: document.getElementById("new-album-title").value,
        artist: document.getElementById("new-album-artist").value,
        price: document.getElementById("new-album-price").value,
    };

    if (document.getElementById("album-id-update").value == "") {
        alert("Please enter an album id");
        return;
    }

    console.log(album);

    fetch("http://localhost:8080/albums/edit", {
        method: "PUT",
        mode: "cors",
        headers: {
            "Access-Control-Allow-Origin": "*",
            "Content-Type": "application/json",
        },
        body: JSON.stringify(album),
    })
        .then(function (response) {
            console.log(response);
            return response.json();
        })
        .then(function (data) {
            console.log(data);
            document.getElementById("update-album-result").innerHTML =
                JSON.stringify(data);
        });
});

document
    .getElementById("delete-album-by-id")
    .addEventListener("click", function () {
        id = document.getElementById("album-id-delete").value;

        if (id == "") {
            alert("Please enter an album id");
            return;
        }

        fetch(`http://localhost:8080/albums/${id}/`, {
            mode: "cors",
            method: "DELETE",
            headers: {
                "Access-Control-Allow-Origin": "*",
                "Content-Type": "application/json",
            },
        })
            .then(function (response) {
                console.log(response);
                return response.json();
            })
            .then(function (data) {
                console.log(data);
                document.getElementById("delete-album-by-id-result").innerHTML =
                    JSON.stringify(data);
            });
    });
