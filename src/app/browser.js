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
	view:function(vnode){
		return (
			<section style="" class="vh-100 bg-white mw8 center w-100 ">
				<article class="min-vh-100 dt w-100">
				  <div class="dtc v-mid tc dark-gray ph2">
						<div class="measure w-30-l center tc pa3">
							<img class="db center h4" src="../../assets/img/icon.svg" />
							<span class="b i">vaccine8</span>
						</div>
						
				  </div>
				</article>
			</section>
		)
	}
}

export default page;
