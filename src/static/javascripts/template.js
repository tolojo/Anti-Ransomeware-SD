async function saveFile() {
    try {
        let response = await $.ajax({
            contentType: 'application/json',
            url:`/save`,
            method: "POST",
            data: JSON.stringify({"name":document.getElementById("nome").value}),
            dataType: "json",
        })
        if (response.status=200){
            alert("Guardado");
        }
    } catch (e) {
        console.log(e)
    }
}

async function compareFile() {
    try {
        let response = await $.ajax({
            contentType: 'application/json',
            url:`/validate`,
            method: "POST",
            data: JSON.stringify({"name":document.getElementById("nome").value}),
            dataType: "json",
        })
    } catch (e) {
        console.log(e)
    }
}