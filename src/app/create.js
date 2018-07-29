import m from 'mithril';
import menu from './#menu.js';
import Icons from '../#icons.js';
import {appAlert} from '../#utils.js';
import {checkRedirect} from '../#utils.js';

var page = {
	oninit:function(vnode){
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
				<div id="particles-js" class="absolute vh-100 left-0 w-100"></div>
				<article class="min-vh-100 dt w-100 absolute left-0 top-0">
				  <div class="dtc v-mid tc white ph2">
					<div class="center w-80-l">

						<div class="flex items-center">
							<span class="center flex items-center b tracked">
								CREATE A PROFILE
							</span>
						</div>

						<div class="mv2 cf"></div>

						<span class="fl w-100 center br2 flex items-center">
							{m("input", {class:"f6 tracked bn black pa3 br2 br--left w-100", placeholder:"First Name"})}
							<Icons name="person" class="bg-orange h1 w2 pv3 ph3 ph2-ns white pointer br2 br--right"/>
						</span>
						<div class="cf mv2"></div>
						<span class="fl w-100 center br2 flex items-center">
							{m("input", {type:"email", class:"f6 tracked bn black pa3 br2 br--left w-100", placeholder:"Last Name"})}
							<Icons name="envelope-closed" class="bg-orange h1 w2 pv3 ph3 ph2-ns white pointer br2 br--right"/>
						</span>
						<div class="cf mv2"></div>
						<span class="fl w-100 center br2 flex items-center">
							{m("input", {type:"password", class:"f6 tracked bn black pa3 br2 br--left w-100", placeholder:"Relationship"})}
							<Icons name="lock-locked" class="bg-orange h1 w2 pv3 ph3 ph2-ns white pointer br2 br--right"/>
						</span>
						<div class="cf mv2"></div>
						<span class="fl w-100 center br2 flex items-center">
							{m("input", {type:"password", class:"f6 tracked bn black pa3 br2 br--left w-100", placeholder:"Confirm Password"})}
							<Icons name="lock-locked" class="bg-orange h1 w2 pv3 ph3 ph2-ns white pointer br2 br--right"/>
						</span>
						<div class="cf mv3"></div>
						<div class=" tc">
							<a class="link bg-gold ttu b tracked near-white shadow-4 pointer fl w-100 dim pv3 br2" href="/app">Sign Up</a>
						</div>
					
					</div>
				  </div>
				</article>
			</section>
		)
	}
}

export default page;
