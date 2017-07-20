	<hr />

	<nav class="level">
	  <div class="level-item has-text-centered">
	    <div>
	      <p class="heading"><img src="/static/img/logo-msu.png"></p>
	      <p class="title"><a href="http://www.msu.ru/" target="_blank">Lomonosov MSU</a></p>
	    </div>
	  </div>
	  <div class="level-item has-text-centered">
	    <div>
	      <p class="heading"><img src="/static/img/logo.png"></p>
	      <p class="title"><a href="http://msu.uz/" target="_blank">Lomonosov MSU Tashkent Branch</a></p>
	    </div>
	  </div>
	  <div class="level-item has-text-centered">
	    <div>
	      <p class="heading"><img src="/static/img/acm.jpeg"></p>
	      <p class="title"><a href="https://icpc.baylor.edu/" target="_blank">ACM ICPC</a></p>
	    </div>
	  </div>
	  <div class="level-item has-text-centered">
	    <div>
	      <p class="heading"><img src="/static/img/codefoces.png"></p>
	      <p class="title"><a href="http://codeforces.com" target="_blank">codeforces.com</a></p>
	    </div>
	  </div>
	</nav>

	<footer class="footer">
	  <div class="container">
	    <div class="content has-text-centered">
	      <p>
	        <strong>CodeStats</strong> by <a href="https://github.com/MrPark97">MrPark97</a>. The source code is licensed
	        <a href="http://opensource.org/licenses/mit-license.php">MIT</a>.
	      </p>
	      <p>
	        <a class="icon" href="https://github.com/MrPark97/codestats">
	          <i class="fa fa-github"></i>
	        </a>
	      </p>
	    </div>
	  </div>
	</footer>


	<script>
		$(document).ready(function() {
			$('.navbar-burger').click(function() {
				if($('.navbar-burger').hasClass('is-active')) {
					$('.navbar-menu').removeClass('is-active');
					$('.navbar-burger').removeClass('is-active');
				} else {
					$('.navbar-menu').addClass('is-active');
					$('.navbar-burger').addClass('is-active');
				}
			});
		});
	</script>
	
</body>
</html>
