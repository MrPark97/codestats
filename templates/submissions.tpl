<section class="hero is-primary is-bold">
  <div class="hero-body">
    <div class="container">
      <h1 class="title">
        Submissions
      </h1>
      <h2 class="subtitle">
        failed and OK
      </h2>
    </div>
  </div>
</section>

<div class="container has-text-centered">

<script type="text/javascript" src="https://www.gstatic.com/charts/loader.js"></script>
<div id="columnchart_values" style="width: 800px; height: 500px;"></div>

<script>
	google.charts.load("current", {packages:['corechart']});
	google.charts.setOnLoadCallback(drawChart);

	function drawChart() {
		{{$sl := .StatsLatestIndex}}

        var data = google.visualization.arrayToDataTable([
          ['Submissions', 'OK', 'Failed', { role: 'annotation' } ],
          {{range $i, $e := .Stats}}

          ['{{$e.Handle}}', {{$e.SuccessCount}}, {{$e.FailedCount}}, '']{{if lt $i $sl}},{{end}}
          {{end}}
        ]);

        var view = new google.visualization.DataView(data);
        view.setColumns([0, 1,
				         { calc: "stringify",
				           sourceColumn: 1,
				           type: "string",
				           role: "annotation" },
				         2]);

        var options = {
          width: 600,
          height: 400,
          legend: { position: 'top', maxLines: 3 },
          bar: { groupWidth: '75%' },
          isStacked: true,
        };

        var chart = new google.visualization.ColumnChart(document.getElementById("columnchart_values"));
        chart.draw(view, options);
	 }
</script>

</div>