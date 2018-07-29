import m from 'mithril';
import Siema from 'siema';

import menu from './#menu.js';
import footer from './#footer.js';

import Icons from '../#icons.js';
import {appAlert} from '../#utils.js';
import {checkRedirect} from '../#utils.js';


var page = {
	classLogin:"",classRegister:"dn",
	classBtnLogin:"white",
	classBtnRegister:"gray",

	switchFormLogin: function() {
		page.classLogin = "";
		page.classRegister = "dn";
		page.classBtnLogin = "white";
		page.classBtnRegister = "gray";
	},
	switchFormRegister: function() {
		page.classLogin = "dn";
		page.classRegister = "";
		page.classBtnLogin = "gray";
		page.classBtnRegister = "white";
	},

	sliderItem: { view: function(vnode) {
	return(m("div",{class:"w-100 vh-75 vh-50 parallaxBG", style:"background-image:url('../../assets/img/"+vnode.attrs.filepath+"');"},))
	}},
	sliderInit: function(vnode){
		var searchList = [];
		searchList.push(m(page.sliderItem,{ filepath: "polio1.jpg" }));

		searchList.push(m(page.sliderItem,{ filepath: "polio2.jpg" }));

		searchList.push(m(page.sliderItem,{ filepath: "polio3.jpg" }));

		if(searchList.length > 0) {
			page.sliderContainer = searchList; m.redraw();
			page.mySiema = new Siema({loop:true,});
		}
	},
	oninit:function(vnode){
		setTimeout(function(){page.sliderInit()},250);
		m.mount(document.getElementById('appMenu'), menu)
	},
	oncreate:function(vnode){
		particlesJS.load('particles-js', '../assets/bin/particles.json', function() {
	 		console.log('callback - particles.js config loaded');
	 	});
	},
	view:function(vnode){
	return (
		<section style="" class="min-vh-100 mw8 center w-100">
			<div id="particles-js" class="absolute left-0 w-100"></div>
				<div class="flex flex-row">
					<div class="w-80-l w-100 center relative">
						<div class="siema w-100">
							{page.sliderContainer}
						</div>
					</div>
				</div>
			<article class="w-100 relative left-0 top-0">
			  <div class="tc white">
					<div class="center w-80-l">

						<div class="fl w-100">

							<div class="w-100 br2">


								<div class="monaco pv3  center black flex flex-row">
									<div class={"w-50 fw5 bw1 pa3 br b--white pointer "+page.classBtnLogin} onclick={page.switchFormLogin}>
										My Records
									</div>
									<div class={"w-50 fw5 pa3 pointer "+page.classBtnRegister} onclick={page.switchFormRegister}>
										Information
									</div>
								</div>

								<div class={" f6 avenir black tl cf "+page.classLogin}>

									<span class="list bottom-0 z-9 fl bg-white w-100 br3 br--top  overflow-y-scroll " style="min-height:50vh">
										<a class="link" href="/app/schedule">
										<div class="flex items-center lh-copy pa3 bb b--near-white hover-bg-green  green hover-gold ">
											<div class="pa1 br-100 bg-red">
												<Icons name="bell" class="white h1 fr br1" onclickX={menu.toggle}/>
											</div>
											<div class="pl3 flex-auto">
												<span class="f6 db black-70"><b>DANIEL - My SON</b></span>
											</div>
											<div class="tr">
												<Icons name="chevron-right" class=" h1 fr br1"/>
											</div>
										</div>
										</a>

										<a class="link" href="/app/schedule">
										<div class="flex items-center lh-copy pa3 bb b--near-white hover-bg-green  green hover-gold ">
											<div class="pa1 br-100 bg-green">
												<Icons name="bell" class="white h1 fr br1" onclickX={menu.toggle}/>
											</div>
											<div class="pl3 flex-auto">
												<span class="f6 db black-70"><b>DANIEL - My SON</b></span>
											</div>
											<div class="tr">
												<Icons name="chevron-right" class=" h1 fr br1"/>
											</div>
										</div>
										</a>
									</span>
								</div>

								<div class={" f6 avenir black tl cf "+page.classRegister}>
									<article class="cf br3 br--top bg-white w-100">
										  <div class="fl w-50 bb br b--near-white tc">
										  	<div class="pa4 tc">
										  	  <img src="http://tachyons.io/img/logo.jpg"
										  	      class="br-100 ba h3 w3 dib" alt="avatar">
										  	   </img>
										    	<label class="db f5 f4-l">Column One</label>
										  	</div>
										  </div>

										  <div class="fl w-50 bb b--near-white tc">
										  	<div class="pa4 tc">
										  	  <img src="http://tachyons.io/img/logo.jpg"
										  	      class="br-100 ba h3 w3 dib" alt="avatar">
										  	   </img>
										    	<label class="db f5 f4-l">Column Two</label>
										  	</div>
										  </div>
										  <div class="fl w-50 bb br b--near-white tc">
										  	<div class="pa4 tc">
										  	  <img src="http://tachyons.io/img/logo.jpg"
										  	      class="br-100 ba h3 w3 dib" alt="avatar">
										  	   </img>
										    	<label class="db f5 f4-l">Column Three</label>
										  	</div>
										  </div>

										  <div class="fl w-50 bb b--near-white tc">
										  	<div class="pa4 tc">
										  	  <img src="http://tachyons.io/img/logo.jpg"
										  	      class="br-100 ba h3 w3 dib" alt="avatar">
										  	   </img>
										    	<label class="db f5 f4-l">Column Four</label>
										  	</div>
										  </div>

										  <div class="fl w-50 bb br b--near-white tc">
										  	<div class="pa4 tc">
										  	  <img src="http://tachyons.io/img/logo.jpg"
										  	      class="br-100 ba h3 w3 dib" alt="avatar">
										  	   </img>
										    	<label class="db f5 f4-l">Column Five</label>
										  	</div>
										  </div>

										  <div class="fl w-50 bb b--near-white tc">
										  	<div class="pa4 tc">
										  	  <img src="http://tachyons.io/img/logo.jpg"
										  	      class="br-100 ba h3 w3 dib" alt="avatar">
										  	   </img>
										    	<label class="db f5 f4-l">Column Six</label>
										  	</div>
										  </div>
									</article>
							</div>

					  </div>
					</div>
			  </div>
			  </div>
			</article>

			<a class={"link "+page.classLogin} href="/app/create">
				<div class="pa3 br-100 bg-green fixed bottom-2 right-2">
					<Icons name="medical-cross" class="white h1 fr br1" />
				</div>
			</a>
		</section>
	)
	}
}

export default page;
