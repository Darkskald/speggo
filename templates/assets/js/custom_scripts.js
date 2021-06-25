(function () {
    // your page initialization code here
    // the DOM will be available
    console.log("Welcome to custom scripting in speggo!");


    /*const form = document.getElementById('searchForm');


    form.addEventListener('submit', function (){
        let value = document.getElementById('searchField').value;
        console.log("lelelelele", value)
        fetch('/sfg/'+value)
            .then(res => res.json())
            .then((out) => {
                console.log('Output: ', out);
                plotSfg(out);
                populateSfgMeta(out);
            }).catch(err => console.error(err));



    });*/

    fetch('/sfg/3')
        .then(res => res.json())
        .then((out) => {
            console.log('Output: ', out);
            plotSfg(out);
        }).catch(err => console.error(err));

    createCard()


})();


function populateSfgMeta(sfg) {
    let sfgName = document.getElementById("sfg_meta_name");
    sfgName.innerHTML = sfg.name;

    let sfgDate = document.getElementById("sfg_meta_date");
    sfgDate.innerHTML = sfg.measured_time;

    let maxIr = Math.max(sfg.ir);
    let maxSfg = Math.max(sfg.sfg_intenstiy);
    let maxVis = Math.max(sfg.vis);
    document.getElementById("sfg_meta_ir").innerHTML = maxIr;
    document.getElementById("sfg_meta_sfg").innerHTML = maxSfg;
    document.getElementById("sfg_meta_vis").innerHTML = maxVis;
}

function searchHandler() {
    console.log("lelele")
}

function createCard() {
    let d1 = document.createElement("div");
    d1.classList.add("w-full", "md:w-1/2", "p-3");

    let d2 = document.createElement("div");
    d2.classList.add("bg-gray-900", "border", "border-gray-800", "rounded", "shadow");

    let d3 = document.createElement("div");
    d3.classList.add("border-b", "border-gray-800", "p-3");

    let h = document.createElement("h5");
    // todo: add id
    h.classList.add("font-bold", "uppercase", "text-gray-600");

    let d4 = document.createElement("div");
    d4.classList.add("p-5");

    let d5 = document.createElement("div");
    d5.id = "tester2"

    d4.appendChild(d5)
    d3.appendChild(h)
    d2.appendChild(d3)
    d2.appendChild(d4)
    d1.appendChild(d2)

    let cc = document.getElementById("card_container")
    cc.appendChild(d1)
}

