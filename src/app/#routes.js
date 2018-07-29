var m = require("mithril")

//Generic Website Pages
import indexPage from './index.js';
import browserPage from './browser.js';
import accountPage from './account.js';
import recipientPage from './recipient.js';

import createPage from './create.js';
import schedulePage from './schedule.js';

//


m.route.setOrig = m.route.set;
m.route.set = function(path, data, options){
	m.route.setOrig(path, data, options);
	window.scrollTo(0,0);
}

m.route.linkOrig = m.route.link;
m.route.link = function(vnode){
	m.route.linkOrig(vnode);
	window.scrollTo(0,0);
}

m.route.prefix("")
m.route.mode = "pathname"
m.route(document.getElementById('appContent'), "/app", {
	"/app":{ view: function(vnode) { return m(indexPage);},},

	"/app/browser":{ view: function(vnode) { return m(browserPage);},},
	"/app/account":{ view: function(vnode) { return m(accountPage);},},

	"/app/recipient":{ view: function(vnode) { return m(recipientPage);},},
	"/app/create":{ view: function(vnode) { return m(createPage);},},
	"/app/schedule":{ view: function(vnode) { return m(schedulePage);},},

});
