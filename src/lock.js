import m from 'mithril';
import Icons from './#icons.js';
import {appAlert} from './#utils.js';
import {checkRedirect} from './#utils.js';


var page = {
	Submit: function() {
		var actionFields = [
			{validationType : '', fieldID : 'username'},
			{validationType : '', fieldID : 'password'},
		]
		validateSubmit( "/api/login", actionFields);
	},

	alert:function(vnode){
		alert("BAM BA");
	},
	oninit:function(vnode){
		// m.mount(document.getElementById('appMenu'), menu)
	},
	oncreate:function(vnode){
		particlesJS.load('particles-js', '../assets/bin/particles.json', function() {
	 		console.log('callback - particles.js config loaded');
	 	});
	},
	view:function(vnode){
		return (
			<section style="" class="vh-100 mw8 center w-100 ">
				<div id="particles-js" class="w-100 vh-100"></div>
				<article class="min-vh-100 dt w-100 absolute top-0 left-0">
				  <div class="dtc v-mid tc white ph2">
						<div class="measure w-30-l center tc pa3">
							<img class="db center h4" src="../../assets/img/icon.svg" />
							<span class="b i">vaccine8</span>
							<div class="cf mv3"></div>
							<span class="fl w-100 center br2 flex items-center z-max">
								{m("input", {class:"f6 tracked bn black pa3 br2 br--left w-100", placeholder:"Enter Password"})}
								<Icons name="chevron-right" class="bg-gold h1 w2 pv3 ph3 ph2-ns white pointer br2 br--right"
								onclick={page.alert}/>
							</span>
						</div>
				  </div>
				</article>

			</section>
		)
	}
}

export default page;
