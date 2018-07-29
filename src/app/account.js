import m from 'mithril';
import Siema from 'siema';

import menu from './#menu.js';
import footer from './#footer.js';

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
			<section style="" class=" mw8 center">
				<article class="vh-25 dt w-100">
					<div class="dtc v-mid tc ph2 relative">
						<div id="particles-js" class="w-100 vh-25 absolute top-0 left-0"></div>
						<span class="tc f4 fw6 white tracked" style=""> MY ACCOUNTS </span>
					</div>
				</article>

				<span class="list bottom-0 z-9 fl bg-white w-100 br3 br--top  overflow-y-scroll vh-75">
				  <div class="flex items-center lh-copy pa3 bb b--near-white hover-bg-yellow">
				      <img class="w2 h2 w3-ns h3-ns br-100" src="http://tachyons.io/img/avatar-jxnblk.jpg" />
				      <div class="pl3 flex-auto">
				        <span class="f6 db black-70">Jxnblk</span>
								<small class="f7 truncate db black-70">0x6c3b33497C0d5eb19084dcd3e18cFf8C1ad4842a</small>
				      </div>
				  </div>
				  <div class="flex items-center lh-copy pa3 bb b--near-white hover-bg-yellow">
				      <img class="w2 h2 w3-ns h3-ns br-100" src="http://tachyons.io/img/avatar-jasonli.jpg" />
				      <div class="pl3 flex-auto">
				        <span class="f6 db black-70">Jason Li</span>
								<small class="f7 truncate db black-70">0x6c3b33497C0d5eb19084dcd3e18cFf8C1ad4842a</small>
				      </div>
				  </div>
				  <div class="flex items-center lh-copy pa3 bb b--near-white hover-bg-yellow">
				      <img class="w2 h2 w3-ns h3-ns br-100" src="http://tachyons.io/img/avatar-jxnblk.jpg" />
				      <div class="pl3 flex-auto">
				        <span class="f6 db black-70">Jxnblk</span>
								<small class="f7 truncate db black-70">0x6c3b33497C0d5eb19084dcd3e18cFf8C1ad4842a</small>
				      </div>
				  </div>
				  <div class="flex items-center lh-copy pa3 bb b--near-white hover-bg-yellow">
				      <img class="w2 h2 w3-ns h3-ns br-100" src="http://tachyons.io/img/avatar-jxnblk.jpg" />
				      <div class="pl3 flex-auto">
				        <span class="f6 db black-70">Jxnblk</span>
								<small class="f7 truncate db black-70">0x6c3b33497C0d5eb19084dcd3e18cFf8C1ad4842a</small>
				      </div>
				  </div>
				  <div class="flex items-center lh-copy pa3 bb b--near-white hover-bg-yellow">
				      <img class="w2 h2 w3-ns h3-ns br-100" src="http://tachyons.io/img/avatar-jxnblk.jpg" />
				      <div class="pl3 flex-auto">
				        <span class="f6 db black-70">Jxnblk</span>
								<small class="f7 truncate db black-70">0x6c3b33497C0d5eb19084dcd3e18cFf8C1ad4842a</small>
				      </div>
				  </div>
				  <div class="flex items-center lh-copy pa3 bb b--near-white hover-bg-yellow">
				      <img class="w2 h2 w3-ns h3-ns br-100" src="http://tachyons.io/img/avatar-jasonli.jpg" />
				      <div class="pl3 flex-auto">
				        <span class="f6 db black-70">Jason Li</span>
								<small class="f7 truncate db black-70">0x6c3b33497C0d5eb19084dcd3e18cFf8C1ad4842a</small>
				      </div>
				  </div>
				  <div class="flex items-center lh-copy pa3 bb b--near-white hover-bg-yellow">
				      <img class="w2 h2 w3-ns h3-ns br-100" src="http://tachyons.io/img/avatar-yavor.jpg" />
				      <div class="pl3 flex-auto">
				        <span class="f6 db black-70">Yavor</span>
								<small class="f7 truncate db black-70">0x6c3b33497C0d5eb19084dcd3e18cFf8C1ad4842a</small>
				      </div>
				  </div>
				</span>
			</section>
		)
	}
}

export default page;
