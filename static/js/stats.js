google.charts.load('current', {
    'packages': ['corechart']
});
google.charts.setOnLoadCallback(drawChart);

function drawChart() {
    var data = google.visualization.arrayToDataTable([
        ['Date', 'Download', 'Upload'],
        ['01-02', 10, 32],
        ['01-03', 12, 28],
        ['01-04', 42, 38],
        ['01-05', 37, 28],
        ['01-06', 57, 27],
        ['01-07', 62, 17],
        ['01-08', 32, 20],
        ['01-09', 48, 21],
        ['01-10', 49, 9],
    ]);

    var options = {
        title: 'Download & Upload',
        curveType: 'function',
        legend: {
            position: 'bottom'
        }
    };

    var chart = new google.visualization.LineChart(document.getElementById('curve_chart'));

    chart.draw(data, options);
}