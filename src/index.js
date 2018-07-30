import m from 'mithril';
import Icons from './#icons.js';
import {appAlert} from './#utils.js';
import {checkRedirect} from './#utils.js';

var page = {
	classLogin:"",classRegister:"dn",
	classBtnLogin:"yellow",
	classBtnRegister:"gray",

	switchFormLogin: function() {
		page.classLogin = "";
		page.classRegister = "dn";
		page.classBtnLogin = "yellow";
		page.classBtnRegister = "gray";
	},
	switchFormRegister: function() {
		page.classLogin = "dn";
		page.classRegister = "";
		page.classBtnLogin = "gray";
		page.classBtnRegister = "yellow";
	},

	Submit: function() {
		var actionFields = [
			{validationType : '', fieldID : 'username'},
			{validationType : '', fieldID : 'password'},
		]
		validateSubmit( "/api/login", actionFields);
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
			<section style="" class="min-vh-100 mw8 center w-100">
				<div id="particles-js" class="absolute left-0 w-100"></div>
				<article class="min-vh-100 dt w-100 absolute left-0 top-0">
				  <div class="dtc v-mid tc white ph2">
						<div class="center w-80-l">

						  <div class="fl w-100 w-50-l relative">
								<div class="ph2 w-100 br2 ">
									<article class="h5-l dt w-100">
										<div class="dtc v-mid tc white monaco ">
											<img class="db center h4" src="../../assets/img/logo.png" />
											<span class="b f1-l i">vaccine8</span>
											<div class="pv2 white tc f6 f5-l tracked i fw3">
												 Medical Health Records
											</div>
											<div class="cf mv2"></div>

										</div>
									</article>
								</div>
							</div>

							<div class="fl w-100 w-50-l pa2">

								<div class="w-100 br2">
									<div class="monaco center black flex flex-row">
										<div class={"w-50 fw5 pa3 br b--gold pointer "+page.classBtnLogin} onclick={page.switchFormLogin}>
											Login
										</div>
										<div class={"w-50 fw5 pa3 pointer "+page.classBtnRegister} onclick={page.switchFormRegister}>
											Register
										</div>
									</div>

									<div class={"ph3 f6 avenir black tl cf pt4 pb2 "+page.classLogin}>

										<span class="fl w-100 center br2 flex items-center">
											{m("input", {type:"email", class:"f6 tracked bn black pa3 br2 br--left w-100", placeholder:"Enter Email"})}
											<Icons name="envelope-closed" class="bg-gold h1 w2 pv3 ph3 ph2-ns white pointer br2 br--right"/>
										</span>
										<div class="cf mv2"></div>
										<span class="fl w-100 center br2 flex items-center">
											{m("input", {class:"f6 tracked bn black pa3 br2 br--left w-100", placeholder:"Enter Password"})}
											<Icons name="chevron-right" class="bg-gold h1 w2 pv3 ph3 ph2-ns white pointer br2 br--right"/>
										</span>
										<div class="cf mv3"></div>
										<div class=" tc">
											<a class="link bg-gold ttu b tracked near-white shadow-4 pointer fl w-100 dim pv3 br2" href="/app/">Log In</a>
										</div>
									</div>

									<div class={"ph3 f6 avenir black tl cf pt4 pb2 "+page.classRegister}>
										<span class="fl w-100 center br2 flex items-center">
											{m("input", {class:"f6 tracked bn black pa3 br2 br--left w-100", placeholder:"Enter Name"})}
											<Icons name="person" class="bg-orange h1 w2 pv3 ph3 ph2-ns white pointer br2 br--right"/>
										</span>
										<div class="cf mv2"></div>
										<span class="fl w-100 center br2 flex items-center">
											{m("input", {type:"email", class:"f6 tracked bn black pa3 br2 br--left w-100", placeholder:"Enter Email"})}
											<Icons name="envelope-closed" class="bg-orange h1 w2 pv3 ph3 ph2-ns white pointer br2 br--right"/>
										</span>
										<div class="cf mv2"></div>
										<span class="fl w-100 center br2 flex items-center">
											{m("input", {type:"password", class:"f6 tracked bn black pa3 br2 br--left w-100", placeholder:"Enter Password"})}
											<Icons name="lock-locked" class="bg-orange h1 w2 pv3 ph3 ph2-ns white pointer br2 br--right"/>
										</span>
										<div class="cf mv2"></div>
										<span class="fl w-100 center br2 flex items-center">
											{m("input", {type:"password", class:"f6 tracked bn black pa3 br2 br--left w-100", placeholder:"Confirm Password"})}
											<Icons name="lock-locked" class="bg-orange h1 w2 pv3 ph3 ph2-ns white pointer br2 br--right"/>
										</span>
										<div class="cf mv3"></div>
										<div class=" tc">
											<a class="link bg-gold ttu b tracked near-white shadow-4 pointer fl w-100 dim pv3 br2" href="/app/">Sign Up</a>
										</div>
									</div>
								</div>

						  </div>



						</div>
				  </div>
				</article>
			</section>
		)
	}
}

export default page;
