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
			<section style="" class="  center w-100 ">
				<article class=" dt w-100">
				  <div class="dtc v-top tc white ph2 relative">
						<div id="particles-js" class="w-100 vh-75 absolute top-0 left-0"></div>
						<div class="measure center tc pa3">

							<img class="db center h4" src="../../assets/img/icon.svg" />
							<div class="fl w-100">
								<p class="f5 tj center tracked ">
									<i class="b">vaccine8</i> is a mobile dapp browser and wallet for the Ethereum blockchain
								</p>
								<p class="f5 tj center tracked">
									It allows you to interact with dapps powered by Ethereum on your mobile device and makes it easy for you to securely store, send and receive Ether and ERC20 tokens.
								</p>
							</div>
						</div>

				  </div>
				</article>
			</section>
		)
	}
}

export default page;
