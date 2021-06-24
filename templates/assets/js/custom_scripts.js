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

function plotSfg(sfg) {
    document.getElementById("headline").innerHTML = sfg.name;
    const layout = {
        paper_bgcolor: 'rgb(180,180,199)',
        plot_bgcolor: 'rgb(180,180,199)',
        autosize: true,
        //margin: {l: 0.1, t: 0.1, r: 0.1, b: 0.1},
        yaxis: {
            automargin: true,
            title: 'wavenumber / cm-1'
        },
        xaxis: {
            automargin: true,
            title: 'norm SFG intensity/ arb.u.'
        },

    }

    TESTER = document.getElementById('tester');
    Plotly.newPlot(TESTER, [{
            x: sfg.wavenumber,
            y: sfg.sfg_intensity,

        }], layout,
        {responsive: true});
}


const config = {
    responsive: true
}
