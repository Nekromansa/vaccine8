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
						<span class="tc f4 fw6 white tracked" style=""> SCHEDULE </span>
					</div>
				</article>

				<span class="list bottom-0 z-9 fl bg-white w-100 br3 br--top  overflow-y-scroll vh-75">
					<div class="flex items-center lh-copy pa2 bb b--near-white hover-bg-yellow">
							<div class="pl3 flex-auto">
								<span class="f6 db black-70">Diseases: Polio</span>
								<span class="f6 db black-70 truncate"><b>Vaccine: OPV</b></span>
							</div>
							<div class="tr">
								<span class="f7 i link dark-green"><b>25 May 2020</b></span>
							</div>
					</div>
					<div class="flex items-center lh-copy pa2 bb b--near-white hover-bg-yellow">
							<div class="pl3 flex-auto">
								<span class="f6 db black-70">Diseases: Yellow fever</span>
								<span class="f6 db black-70 truncate"><b>Vaccine: Yellow fever vaccine</b></span>
							</div>
							<div class="tr">
								<span class="f7 i link dark-green"><b>25 May 2011</b> - <b class="red">OVERDUE</b> </span>
							</div>
					</div>
					<div class="flex items-center lh-copy pa2 bb b--near-white hover-bg-yellow">
							<div class="pl3 flex-auto">
								<span class="f6 db black-70">Diseases: Measles Vaccine</span>
								<span class="f6 db black-70 truncate"><b>Vaccine: Measles Vaccine</b></span>
							</div>
							<div class="tr">
								<span class="f7 i link dark-green"><b>25 May 2011</b> - <b class="red">OVERDUE</b> </span>
							</div>
					</div>
					<div class="flex items-center lh-copy pa2 bb b--near-white hover-bg-yellow">
							<div class="pl3 flex-auto">
								<span class="f6 db black-70">Diseases: Polio</span>
								<span class="f6 db black-70 truncate"><b>Vaccine: OPV</b></span>
							</div>
							<div class="tr">
								<span class="f7 i link dark-green"><b>25 May 2011</b> - <b class="red">OVERDUE</b> </span>
							</div>
					</div>
					<div class="flex items-center lh-copy pa2 bb b--near-white hover-bg-yellow">
							<div class="pl3 flex-auto">
								<span class="f6 db black-70">Diseases: Yellow fever</span>
								<span class="f6 db black-70 truncate"><b>Vaccine: Yellow fever vaccine</b></span>
							</div>
							<div class="tr">
								<span class="f7 i link dark-green"><b>25 May 2020</b></span>
							</div>
					</div>
					<div class="flex items-center lh-copy pa2 bb b--near-white hover-bg-yellow">
							<div class="pl3 flex-auto">
								<span class="f6 db black-70">Diseases: Measles Vaccine</span>
								<span class="f6 db black-70 truncate"><b>Vaccine: Measles Vaccine</b></span>
							</div>
							<div class="tr">
								<span class="f7 i link dark-green"><b>25 May 2040</b></span>
							</div>
					</div><div class="flex items-center lh-copy pa2 bb b--near-white hover-bg-yellow">
							<div class="pl3 flex-auto">
								<span class="f6 db black-70">Diseases: Polio</span>
								<span class="f6 db black-70 truncate"><b>Vaccine: OPV</b></span>
							</div>
							<div class="tr">
								<span class="f7 i link dark-green"><b>25 May 2020</b></span>
							</div>
					</div>
					<div class="flex items-center lh-copy pa2 bb b--near-white hover-bg-yellow">
							<div class="pl3 flex-auto">
								<span class="f6 db black-70">Diseases: Yellow fever</span>
								<span class="f6 db black-70 truncate"><b>Vaccine: Yellow fever vaccine</b></span>
							</div>
							<div class="tr">
								<span class="f7 i link dark-green"><b>25 May 2020</b></span>
							</div>
					</div>
					<div class="flex items-center lh-copy pa2 bb b--near-white hover-bg-yellow">
							<div class="pl3 flex-auto">
								<span class="f6 db black-70">Diseases: Measles Vaccine</span>
								<span class="f6 db black-70 truncate"><b>Vaccine: Measles Vaccine</b></span>
							</div>
							<div class="tr">
								<span class="f7 i link dark-green"><b>25 May 2040</b></span>
							</div>
					</div>
				</span>
			</section>
		)
	}
}

export default page;
