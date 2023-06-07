let myProject = [];

function projectPost(event) {
    // untuk menghilangkan refresh atau reload
    event.preventDefault();

    // untuk memasukan elemen berdasarkan id dengamn nilai ( value ) yang diinput ke variable
    let projectName = document.getElementById('input-title').value;
    let startDate = document.getElementById('start-date').value;
    let endDate = document.getElementById('end-date').value;
    let message = document.getElementById('input-textarea').value;
    let image = document.getElementById('input-image').files;

    //  untuk memasukan icon kedalam variabel
    let nodeJs = '<i class="fa-brands fa-node"></i>';
    let javaScript = '<i class="fa-brands fa-js"></i>';
    let reactJs = '<i class="fa-brands fa-react"></i>';
    let golang = '<i class="fa-brands fa-golang"></i>';

    //  untuk memasukan element berasarkan id dengan checkbox yang dipilih kedalam variable
    let iconNode = document.getElementById('node').checked ? nodeJs : "";
    let iconJavascript = document.getElementById('javascript').checked ? javaScript : "";
    let iconreact = document.getElementById('react').checked ? reactJs : "";
    let iconGolang = document.getElementById('golang').checked ? golang : "";

    // perhitungan durasi 
    let inputStartDate = new Date(startDate);
    let inputEndDate = new Date(endDate);
    let selisih = Math.abs(inputEndDate.getMonth() - inputStartDate.getMonth() + ((inputEndDate.getFullYear() - inputStartDate.getFullYear())));
    console.log(selisih);


    // membuat url image
    image = URL.createObjectURL(image[0]);
    console.log(image);

    // // menampung data project
    let projectView = {
        projectName,
        startDate,
        endDate,
        selisih,
        // postAt :new Date(),
        // auhor :"fahmi burhan",
        message,
        iconNode,
        iconreact,
        iconJavascript,
        iconGolang,
        image,
    };
    // memasukan value ke myproject
    myProject.push(projectView);
    console.log(myProject);
    // memanggil function ketika blok project ditambahkan
    renderProject();
}

function renderProject() {
    document.getElementById('content').innerHTML = "";

    for (let i = 0; i < myProject.length; i++) {
        document.getElementById('content').innerHTML += `

        <div class="col col-md-4 d-flex justify-content-center my-4 ">
        <div class="card shadow" style="width: 25rem;">
          <img src="${myProject[i].image}" class="card-img-top p-3" alt="gambar">
          <div class="card-body">
            <h5 class="card-title">${myProject[i].projectName}</h5>
            <div>
              <p class="durasi">${myProject[i].selisih} Month</p>
              <p class="durasi-1">${myProject[i].message}</p>
            </div> 
            <div class="gep-3">
                ${myProject[i].iconNode}
                ${myProject[i].iconreact}
                ${myProject[i].iconJavascript} 
                ${myProject[i].iconGolang}
            </div>
          </div>
          <div class="mt-4 d-flex justify-content-center gap-3 mb-3">
            <button class="btn btn-dark text-light px-5">Detail</button>
            <button class="btn  btn-dark text-light px-5">Delete</button>
          </div>
        </div>
      </div>
     `;
    }
}