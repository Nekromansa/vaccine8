import m from 'mithril';
import Icons from './#icons.js';

export var page = {
	oninit:function(vnode){},
	oncreate:function(vnode){
		particlesJS.load('particles-js', '../assets/bin/particles.json', function() {
	 		console.log('callback - particles.js config loaded');
	 	});
	},
	view:function(vnode){
		return (
			<section style="" class="min-vh-100 mw8 center w-100 ">
				<div id="particles-js" class="w-100 vh-100"></div>
				<article class="min-vh-100 dt w-100 absolute left-0 top-0">
				  <div class="dtc v-mid tc white ph2  w-80-l">
						<div class="center w-50-l w-70-m">



							<div class="fl w-100 pa3">
								<div class=" w-100 br2">

									<div class="monaco center yellow flex flex-row pb2">
										<div class={"w-100 fw5  bb b--gold pointer "}>
											Restore Your Account
										</div>
									</div>


									<div class={" f6 avenir white tl cf p pv2 "}>
										<span class="fl w-100 center br2 relative">
											{m("textarea", {class:"f6 h4 tracked bn black bg-light-gray pa3 br2 w-100", placeholder:"Enter Account Seed"})}
										</span>
										<div class="cf mv2"></div>
										<span class="fl w-100 center br2 flex items-center">
											{m("input", {class:"f6 tracked bn black pa3 br2 br--left w-100", placeholder:"Enter Password"})}
											<Icons name="lock-locked" class="bg-orange h1 w2 pv3 ph3 ph2-ns white pointer br2 br--right"/>
										</span>
										<div class="cf mv2"></div>
										<span class="fl w-100 center br2 flex items-center">
											{m("input", {class:"f6 tracked bn black pa3 br2 br--left w-100", placeholder:"Confirm Password"})}
											<Icons name="lock-locked" class="bg-orange h1 w2 pv3 ph3 ph2-ns white pointer br2 br--right"/>
										</span>
										 <div class="cf mv3"></div>
										<div class=" tc">
											<span class="bg-gold ttu b tracked near-white shadow-4 pointer fl w-100 dim pv3 br2" onclick={page.Submit}>Restore Account</span>
										</div>
									</div>
								</div>
								<div class="w-100 br2">
									<div class="f6 avenir tl cf">
										<a href="/" oncreate={m.route.link} class="yellow no-underline ph1 br1">
											<span class="fw5 db">Already Registered?</span>
											<span class="fw5 db gold">Login using password</span>
										</a>
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
