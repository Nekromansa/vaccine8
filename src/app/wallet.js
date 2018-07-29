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
												My Records
											</div>
											<div class={"w-50 fw5 pa3 pointer "+page.classBtnRegister} onclick={page.switchFormRegister}>
												Information
											</div>
										</div>

										<div class={"ph3 f6 avenir black tl cf pt4 pb2 "+page.classLogin}>

										<span class="list bottom-0 z-9 fl bg-white w-100 br3 br--top  overflow-y-scroll ">
											<div class="flex items-center lh-copy pa2 bb b--near-white hover-bg-yellow">
													<Icons class="w2 h2 red" name="arrow-thick-top" />
													<div class="pl3 flex-auto">
														<span class="f6 db black-70">Jxnblk</span>
														<span class="f6 db black-70 truncate">0.002 <b>ETH</b></span>
													</div>
													<div class="tr">
														<span class="f7 i link dark-green"><b>3</b> confirmations</span>
													</div>
											</div>
										  <div class="flex items-center lh-copy pa2 bb b--near-white hover-bg-yellow">
										      <Icons class="w2 h2 green" name="arrow-thick-bottom" />
										      <div class="pl3 flex-auto">
										        <span class="f6 db black-70">Jxnblk</span>
														<span class="f6 db black-70 truncate">0.342 <b>ETH</b></span>
										      </div>
										      <div class="tr">
										        <span class="f7 i link black dark-red"><b>0</b> confirmations</span>
										      </div>
										  </div>
											<div class="flex items-center lh-copy pa2 bb b--near-white hover-bg-yellow">
													<Icons class="w2 h2 red" name="arrow-thick-top" />
													<div class="pl3 flex-auto">
														<span class="f6 db black-70">Jxnblk</span>
														<span class="f6 db black-70 truncate">0.002 <b>ETH</b></span>
													</div>
													<div class="tr">
														<span class="f7 i link dark-green"><b>3</b> confirmations</span>
													</div>
											</div>
										  <div class="flex items-center lh-copy pa2 bb b--near-white hover-bg-yellow">
										      <Icons class="w2 h2 green" name="arrow-thick-bottom" />
										      <div class="pl3 flex-auto">
										        <span class="f6 db black-70">Jxnblk</span>
														<span class="f6 db black-70 truncate">0.342 <b>ETH</b></span>
										      </div>
										      <div class="tr">
										        <span class="f7 i link black dark-red"><b>0</b> confirmations</span>
										      </div>
										  </div>
										</span>
										</div>

										<div class={"ph3 f6 avenir black tl cf pt4 pb2 "+page.classRegister}>
											<article class="cf">
												  <div class="fl w-50 bg-near-white tc">
												  	<div class="pa2 tc">
												  	  <img
												  	      src="http://tachyons.io/img/logo.jpg"
												  	      class="br-100 ba h3 w3 dib" alt="avatar">
												  	   </img>
												  	</div>
												    	<p class="f5 f3-l">Column One</p>
												  </div>

												  <div class="fl w-50 bg-red tc">
												  	<div class="pa2 tc">
												  	  <img
												  	      src="http://tachyons.io/img/logo.jpg"
												  	      class="br-100 ba h3 w3 dib" alt="avatar">
												  	   </img>
												  	</div>
												    	<p class="f5 f3-l">Column Two</p>
												  </div>
												  <div class="fl w-50 bg-red tc">
												  	<div class="pa2 tc">
												  	  <img
												  	      src="http://tachyons.io/img/logo.jpg"
												  	      class="br-100 ba h3 w3 dib" alt="avatar">
												  	   </img>
												  	</div>
												    	<p class="f5 f3-l">Column Three</p>
												  </div>

												  <div class="fl w-50 bg-near-white tc">
												  	<div class="pa2 tc">
												  	  <img
												  	      src="http://tachyons.io/img/logo.jpg"
												  	      class="br-100 ba h3 w3 dib" alt="avatar">
												  	   </img>
												  	</div>
												    	<p class="f5 f3-l">Column Four</p>
												  </div>

												  <div class="fl w-50 bg-near-white tc">
												  	<div class="pa2 tc">
												  	  <img
												  	      src="http://tachyons.io/img/logo.jpg"
												  	      class="br-100 ba h3 w3 dib" alt="avatar">
												  	   </img>
												  	</div>
												    	<p class="f5 f3-l">Column Five</p>
												  </div>

												  <div class="fl w-50 bg-red tc">
												  	<div class="pa2 tc">
												  	  <img
												  	      src="http://tachyons.io/img/logo.jpg"
												  	      class="br-100 ba h3 w3 dib" alt="avatar">
												  	   </img>
												  	</div>
												    	<p class="f5 f3-l">Column Six</p>
												  </div>
											</article>
									</div>
									
							  </div>


								<div class="fl w-100 ph3">
									<p class="f5 tl center tracked ">
										vaccine8 is a mobile dapp browser and wallet for the Ethereum blockchain
									</p>
									<p class="f5 tl center tracked dn">
										vaccine8 allows you to interact with dapps powered by Ethereum on your mobile device and makes it easy for you to securely store, send and receive Ether and ERC20 tokens.
									</p>
								</div>
							</div>
					  </div>
					  </div>
					</article>
				
				</section>
			</section>
		)
	}
}

export default page;
