function submitData(event){
    event.preventDefault();
    let name = document.getElementById('input-name').value;
    let email = document.getElementById('input-email').value;
    let phone = document.getElementById('input-number').value;
    let subject = document.getElementById('input-subject').value;
    let massage = document.getElementById('input-massage').value;

    if(name== ""){
        return alert("Nama Harus Diisi!")
    }else if(email==""){
        return alert("Email Harus Diisi!")
    }else if(phone ==""){
        return alert("Nomor Telepon Harus Diisi!!")
    }else if(subject==""){
        return alert("Subject Harus Diisi!!")
    }else if(massage ==""){
        return alert("Massage Harus Diisi!!")
    }

    console.log(name);
    console.log(email);
    console.log(phone);
    console.log(subject);
    console.log(massage);

    let emailReceiver="fahmiburhan56@gmail.com";
    let a = document.createElement('a');
    a.href =`mailto:${emailReceiver} ?subject=${subject}&body= Hello, nama saya ${name},${massage}.Silahkan kontak saya dinomor${phone}.`;
    a.click();
    
    //object
    let emailer={
        name,
        email,
        phone,
        subject,
        massage,
    };

    console.log(emailer);


}

