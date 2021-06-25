const layout = {
    paper_bgcolor: 'rgb(180,180,199)',
    plot_bgcolor: 'rgb(180,180,199)',
    autosize: true,
    //margin: {l: 0.1, t: 0.1, r: 0.1, b: 0.1},
    yaxis: {
        automargin: true,
        title: 'norm SFG intensity/ arb.u.'
    },
    xaxis: {
        automargin: true,
        title: 'wavenumber / cm-1'
    }
}
const config =  {responsive: true}

function plotSfg(sfg) {
    document.getElementById("headline").innerHTML = sfg.name;
    data = [{
        x: sfg.wavenumbers,
        y: sfg.sfg_intensity,
        mode: 'lines+markers'

    }];
    let TESTER = document.getElementById('tester');

    Plotly.newPlot(TESTER,data, layout, config);

}


