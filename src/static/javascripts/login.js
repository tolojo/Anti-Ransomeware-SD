async function login(){
    try{
        const ipServer2FA = "http://20.238.103.0:80";
        console.log(ipServer2FA);
        let obj = {
            secret: document.getElementById("key").value,
            email: document.getElementById("email").value,
        }

        let user = await $.ajax({
            url: ipServer2FA + '/api/users/verifytoken',
            method: 'post',
            dataType: 'json',
            data: JSON.stringify(obj),
            contentType: 'application/json'
        });
        console.log(user);
    }
    catch (e){
        
    }
    
    
    
}