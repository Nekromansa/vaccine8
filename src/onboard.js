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

				<a oncreate={m.route.link} href="/login" class="link absolute right-0 top-0 pa3 tracked f6 b tc hot-pink">
					LOGIN/REGISTER
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

						<img class="center h5 h4-ns" src={"assets/img/"+vnode.attrs.slideIcon} />

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
			slide1Status:"purple",slide2Status:"pink", slide3Status:"pink",
			slide4Status:"pink", slide5Status:"pink", slide6Status:"pink",
			slideIcon: "family.svg", slideTitle:"Track your family immunization record.",
			slideDetails:""
		}));

		searchList.push(m(page.sliderItem,{
			slide1Status:"pink",slide2Status:"purple", slide3Status:"pink",
			slide4Status:"pink", slide5Status:"pink", slide6Status:"pink",
			slideIcon: "community.svg", slideTitle:"Track rural community immunization record",
			slideDetails:""
		}));

		searchList.push(m(page.sliderItem,{
			slide1Status:"pink",slide2Status:"pink", slide3Status:"purple",
			slide4Status:"pink", slide5Status:"pink", slide6Status:"pink",
			slideIcon: "access.svg", slideTitle:"Access records from anywhere and on any device",
			slideDetails:""
		}));

		searchList.push(m(page.sliderItem,{
			slide1Status:"pink",slide2Status:"pink", slide3Status:"pink",
			slide4Status:"purple", slide5Status:"pink", slide6Status:"pink",
			slideIcon: "reminders.svg", slideTitle:"Recieve notifications when immunization is due",
			slideDetails:""
		}));

		searchList.push(m(page.sliderItem,{
			slide1Status:"pink",slide2Status:"pink", slide3Status:"pink",
			slide4Status:"pink", slide5Status:"purple", slide6Status:"pink",
			slideIcon: "doctor.svg", slideTitle:"Get informations about immunization in your territory",
			slideDetails:""
		}));

		searchList.push(m(page.sliderItem,{
			slide1Status:"pink",slide2Status:"pink", slide3Status:"pink",
			slide4Status:"pink", slide5Status:"pink", slide6Status:"purple",
			slideIcon: "notification.svg", slideTitle:"Be notified about disease outbreaks in your area",
			slideDetails:""
		}));

		// searchList.push(m(page.sliderItem,{
		// 	slide1Status:"pink",slide2Status:"pink", slide3Status:"pink",
		// 	slide4Status:"pink", slide5Status:"pink", slide6Status:"purple",
		// 	slideIcon: "travel.svg", slideTitle:"Update travel vaccines before each trip",
		// 	slideDetails:""
		// }));



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
