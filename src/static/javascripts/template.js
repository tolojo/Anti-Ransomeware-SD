async function saveSha256() {
    try {
        let abc = await $.ajax({
            url: '/save',
            method: "get",
            dataType: "json"
        })
    } catch (e) {
        console.log(e)
    }

    alert("Botão carregado");
}