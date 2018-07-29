<!doctype html>
<!--[if lt IE 7]>      <html class="no-js lt-ie9 lt-ie8 lt-ie7" lang=""> <![endif]-->
<!--[if IE 7]>         <html class="no-js lt-ie9 lt-ie8" lang=""> <![endif]-->
<!--[if IE 8]>         <html class="no-js lt-ie9" lang=""> <![endif]-->
<!--[if gt IE 8]><!-->
<html class="no-js " id="html" lang="en"> <!--<![endif]-->
	<head>
		<meta charset="UTF-8">
		<link rel="manifest" href="assets/manifest.json">
		<meta name="theme-color" content="#ffffff">

		<title id="title"> Dashboard  </title>
		<meta name="viewport" content="width=device-width, initial-scale=1.0, maximum-scale=5">
		<link rel="stylesheet" href="../assets/css/tachyonsloadingio.css?">
		<link rel="stylesheet" href="../assets/css/custom.css?">

		<link rel="shortcut icon" id="favicon" href="../assets/img/favicon.png">
		<script src="../assets/bin/custom.js?" type="text/javascript" charset="utf-8"></script>
	</head>



	<body id="body" class="bg-body avenir w-100">


		<!--[if lt IE 8]>
			<p class="browserupgrade">You are using an <strong>outdated</strong> browser. Please <a href="http://browsehappy.com/">upgrade your browser</a> to improve your experience.</p>
		<![endif]-->

		<span id="top"></span>
		<section id="appMenu" class="top-0 z-max"></section>


		<section id="appLoading" style="" class="min-vh-100 w-100 absolute dn z-max bg-body">
			<article class="dt center">
				<div class="dtc v-mid tc f6 fw3 vh-75 near-white">
					<img class="h3" src="../assets/img/logo.png"/> <br/>
					<div class="spinner">
						<div class="dot1"></div>
						<div class="dot2"></div>
					</div>
				</div>
			</article>
		</section>

		

		<section id="appContent" style="" class="bg-patternx">
			<script>startLoader("Index")</script>

			<section id="Index" style="" class="min-vh-100 w-100 absolute bg-body">
				<article class="dt center">
					<div class="dtc v-mid tc f6 fw3 vh-75 near-white">
						<img class="h3" src="../assets/img/logo.png"/> <br/>
						<div class="spinner center">
							<div class="dot1"></div>
							<div class="dot2"></div>
						</div>
					</div>
				</article>
			</section>
		</section>

		<section id="appFooter" class=""></section>

  </body>
	<script src="../assets/bin/particles.min.js"></script>
	<script type="text/javascript">
	(function(){
	     // var randomh=Math.random();
	     var randomh=new Date().valueOf();
	     var e = document.getElementsByTagName("script")[0];
	     var d = document.createElement("script");
	     d.src = "../assets/bin/app.min.js?"+randomh+"";
	     d.type = "text/javascript";
	     d.async = true;
	     d.defer = true;
	     e.parentNode.insertBefore(d,e);
	 })();


	</script>
</html>
