import m from 'mithril';
import Siema from 'siema';

import menu from './#menu.js';
import footer from './#footer.js';

import Icons from '../#icons.js';
import {appAlert} from '../#utils.js';
import {checkRedirect} from '../#utils.js';


var page = {
		oninit:function(vnode){
			var menuCustom = { view: function() {
				return m(menu, {iconClass:"",iconName:"home",iconHref:"/app/"})
			}}
			m.mount(document.getElementById('appMenu'), menuCustom)
		},
	oncreate:function(vnode){
		particlesJS.load('particles-js', '../assets/bin/particles.json', function() {
	 		console.log('callback - particles.js config loaded');
	 	});
	},
	view:function(vnode){
		return (
			<section style="" class=" mw8 center w-100">
				<article class="vh-25 dt w-100">
					<div class="dtc v-mid tc ph2 relative">
						<div id="particles-js" class="w-100 vh-25 absolute top-0 left-0"></div>
						<span class="tc f4 fw6 white tracked" style=""> CREATE A PROFILE </span>
					</div>
				</article>

				<article class=" z-9 ph3 ph0-l w-80-l center br3 br--top  overflow-y-scroll vh-75">
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
            {m("input", {type:"password", class:"f6 tracked bn black pa3 br2 br--left w-100", placeholder:"Sex"})}
            <Icons name="lock-locked" class="bg-orange h1 w2 pv3 ph3 ph2-ns white pointer br2 br--right"/>
          </span>
          <div class="cf mv2"></div>
          <span class="fl w-100 center br2 flex items-center">
            {m("input", {type:"password", class:"f6 tracked bn black pa3 br2 br--left w-100", placeholder:"Birthdate"})}
            <Icons name="lock-locked" class="bg-orange h1 w2 pv3 ph3 ph2-ns white pointer br2 br--right"/>
          </span>
          <div class="cf mv2"></div>
          <span class="fl w-100 center br2 flex items-center">
            {m("input", {type:"password", class:"f6 tracked bn black pa3 br2 br--left w-100", placeholder:"Local Govt area"})}
            <Icons name="lock-locked" class="bg-orange h1 w2 pv3 ph3 ph2-ns white pointer br2 br--right"/>
          </span>
          <div class="cf mv2"></div>
          <span class="fl w-100 center br2 flex items-center">
            {m("input", {type:"password", class:"f6 tracked bn black pa3 br2 br--left w-100", placeholder:"State"})}
            <Icons name="lock-locked" class="bg-orange h1 w2 pv3 ph3 ph2-ns white pointer br2 br--right"/>
          </span>
          <div class="cf mv2"></div>
          <span class="fl w-100 center br2 flex items-center">
            {m("input", {type:"password", class:"f6 tracked bn black pa3 br2 br--left w-100", placeholder:"Note"})}
            <Icons name="lock-locked" class="bg-orange h1 w2 pv3 ph3 ph2-ns white pointer br2 br--right"/>
          </span>
          <div class="cf mv3"></div>
          <div class=" tc">
            <a class="link bg-gold ttu b tracked near-white shadow-4 pointer fl w-100 dim pv3 br2" href="/app">Sign Up</a>
          </div>
				</article>
			</section>
		)
	}
}

export default page;
