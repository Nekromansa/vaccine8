var m = require("mithril");
import Icons from '../#icons.js';

// export function menu() {
// 	m.render(document.getElementById('appMenu'), m(menu))
// }

export var menu = {
	menuFixed: "bg-body",
	oninit: function() {
		window.addEventListener('scroll', function() {
			var menuFixed;
			var shrinkOn = 160;
			var distanceY = window.pageYOffset || document.documentElement.scrollTop;
			if (distanceY > shrinkOn) { menuFixed = "bg-body fadeIn fixed"; }
			else { menuFixed = "bg-body"; }

			if (menuFixed !== menu.menuFixed) {
				menu.menuFixed = menuFixed;
				m.redraw();
			}
		});
	},
	linkItem : {
		view: function(vnode) {
			return(
				<a class="link" href={vnode.attrs.href}>
					<li class="tr" onclick={menu.toggle}>
						<p class="ph2 pv3 mv0 white hover-bg-white hover-blue fw5 tracked">{vnode.children}</p>
					</li>
				</a>
			)
		}
	},
	menuItem : {
		view: function(vnode) {
			return(
				<a class="link f5" oncreate={m.route.link} href={vnode.attrs.href}>
					<li class="tr" onclick={menu.toggle}>
						<p class="ph2 pv3 mv0 dark-green hover-bg-gradient hover-white fw5 tracked">
							{vnode.children}
						</p>
					</li>
				</a>
			)
		}
	},
	toggle: function() {
		var appmenuToggle = document.getElementById("menuToggle");
		var appmenuCover = document.getElementById("menuCover");
		appmenuCover.classList.toggle('dn');
		appmenuToggle.classList.toggle('animated');
		appmenuToggle.classList.toggle('bounceInRight');

		// document.getElementById("nav").classList.toggle('dn');
		// document.getElementById("menuBlur").classList.toggle('vh-100');
		document.getElementById("html").classList.toggle('overflow-hidden');
	},
	view: function(vnode) {
		return (
			<section id="menuBlur" class={"z-max w-100 "+menu.menuFixed}>
				<div id="menuCover"  class="fixed z-max right-0 w-100 vh-100 fr dn pa0" style="">
					<ul id="menuToggle" class="fr tr list pl0 vh-100 ma0 bg-yellow overflow-y-scroll overflow-y-hidden-l" style="">
						<li class="w-100 cf pa3">
							<Icons name="cancel" class="h1 dim dib white" onclick={menu.toggle}/>
						</li>
						<a oncreate={m.route.link} class="link white f5" href="/app/account">
							<li class="w-100 ttu tracked cf pa3 bb b--white-30" onclick={menu.toggle}>
								Accounts
							</li>
						</a>
						<a oncreate={m.route.link} class="link white f5" href="/app/transaction">
							<li class="w-100 ttu tracked cf pa3 bb b--white-30" onclick={menu.toggle}>
								Transactions
							</li>
						</a>
						<a oncreate={m.route.link} class="link white f5" href="/app/browser">
							<li class="w-100 ttu tracked cf pa3 bb b--white-30" onclick={menu.toggle}>
								Dapp Browser
							</li>
						</a>
						<a oncreate={m.route.link} class="link white f5" href="/app/recipient">
							<li class="w-100 ttu tracked cf pa3 bb b--white-30" onclick={menu.toggle}>
								Recipients
							</li>
						</a>
						<a oncreate={m.route.link} class="link white f5" href="/app/wallet">
							<li class="w-100 ttu tracked cf pa3 bb b--white-30" onclick={menu.toggle}>
								Wallets
							</li>
						</a>
						<a oncreate={m.route.link} class="link white f5" href="/app/wallet">
							<li class="w-100 ttu tracked cf pa3" onclick={menu.toggle}>
								Settings
							</li>
						</a>
						<li class="tc pa3 fw3 f6 gray" onclick={menu.toggle}>
							<small>Rinkeby Testnet via<br/> Infura Network</small>
						</li>
					</ul>
				</div>

				<nav id="nav" class=" mw8 center w-80-l w-100 pv2	 z-9999 tc">
					<div class=" w-80-l w-100 center">
						<div class="flex items-center">
							<span class="center flex items-center b tracked">
								<img class="center h2 pb2" src="../../assets/img/logo.png" /> 
								VACCINE8
							</span>
							<Icons name="bell" class="white h1 fr pa2 br1" onclickX={menu.toggle}/>
						</div>
					</div>
				</nav>


				<div id="appAlert"></div>
			</section>
		)
	}
}

export default menu;
