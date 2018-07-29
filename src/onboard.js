import m from 'mithril';
import Siema from 'siema';

import Icons from './#icons.js';
import {appAlert} from './#utils.js';
import {checkRedirect} from './#utils.js';


var page = { sliderContainer:[],
	sliderItem: { view: function(vnode) {
		return(
			<article class="min-vh-100 dt w-100" style="">
				<div class="dtc v-mid tc black ph2 relative">

				<a href="/login" class="link absolute right-0 top-0 pa3 tracked f6 b tc purple">
					GET STARTED
				</a>

				<div class="measure center tc relative">

					<div class="w-100 tc">
						<Icons name="media-record" class={"h1 "+vnode.attrs.slide1Status}/>
						<Icons name="media-record" class={"h1 "+vnode.attrs.slide2Status}/>
						<Icons name="media-record" class={"h1 "+vnode.attrs.slide3Status}/>
						<Icons name="media-record" class={"h1 "+vnode.attrs.slide4Status}/>
						<Icons name="media-record" class={"h1 "+vnode.attrs.slide5Status}/>
						<Icons name="media-record" class={"h1 "+vnode.attrs.slide6Status}/>
					</div>

					<div class="cf w-100 pv1"></div>

						<img class="center h5 h4-ns" src={"../../assets/img/"+vnode.attrs.slideIcon} />

					<div class="fl w-100 tc pt3">
						<p class="f4 b dark-gray center tracked ">
							{vnode.attrs.slideTitle}
						</p>

						<p class="dark-gray f5 center tracked">
							{vnode.attrs.slideDetails}
						</p>
					</div>
				</div>
			</div>
		</article>
		)
	}},
	sliderInit: function(vnode){
		var searchList = [];
		searchList.push(m(page.sliderItem,{
			slide1Status:"purple",slide2Status:"near-white", slide3Status:"near-white",
			slide4Status:"near-white", slide5Status:"near-white", slide6Status:"near-white",
			slideIcon: "logo.png", slideTitle:"ONE SLIDE",
			slideDetails:"Join the community in identifying, mapping, collecting and recycling litter.",
		}));

		searchList.push(m(page.sliderItem,{
			slide1Status:"near-white",slide2Status:"purple", slide3Status:"near-white",
			slide4Status:"near-white", slide5Status:"near-white", slide6Status:"near-white",
			slideIcon: "logo.png", slideTitle:"TWO SLIDE",
			slideDetails:"Join the community in identifying, mapping, collecting and recycling litter.",
		}));

		searchList.push(m(page.sliderItem,{
			slide1Status:"near-white",slide2Status:"near-white", slide3Status:"purple",
			slide4Status:"near-white", slide5Status:"near-white", slide6Status:"near-white",
			slideIcon: "logo.png", slideTitle:"THREE SLIDE",
			slideDetails:"Earn Tokens redeemable to cash as you recycle",
		}));

		searchList.push(m(page.sliderItem,{
			slide1Status:"near-white",slide2Status:"near-white", slide3Status:"near-white",
			slide4Status:"purple", slide5Status:"near-white", slide6Status:"near-white",
			slideIcon: "logo.png", slideTitle:"4TH SLIDE",
			slideDetails:"Earn Tokens redeemable to cash as you recycle",
		}));

		searchList.push(m(page.sliderItem,{
			slide1Status:"near-white",slide2Status:"near-white", slide3Status:"near-white",
			slide4Status:"near-white", slide5Status:"purple", slide6Status:"near-white",
			slideIcon: "logo.png", slideTitle:"5TH SLIDE",
			slideDetails:"Earn Tokens redeemable to cash as you recycle",
		}));

		searchList.push(m(page.sliderItem,{
			slide1Status:"near-white",slide2Status:"near-white", slide3Status:"near-white",
			slide4Status:"near-white", slide5Status:"near-white", slide6Status:"purple",
			slideIcon: "logo.png", slideTitle:"6TH SLIDE",
			slideDetails:"Earn Tokens redeemable to cash as you recycle",
		}));


		if(searchList.length > 0) {
			page.sliderContainer = searchList; m.redraw();
			page.mySiema = new Siema({loop:true,});
			document.getElementById("html").classList.toggle('overflow-hidden');
		}
	},
	oninit:function(vnode){
		setTimeout(function(){page.sliderInit()},250);
	},
	oncreate:function(vnode){},
	view:function(vnode){
		return (
			<section style="" class="min-vh-100 bg-white  center w-100 ">
			<div class="flex flex-column flex-row-m">
				<div class="w-100 center relative">
						<div class="siema w-100">
							{page.sliderContainer}
						</div>
					</div>
				</div>
			</section>
		)
	}
}

export default page;
