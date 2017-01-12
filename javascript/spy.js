// ==UserScript==
// @name        A岛大侦探
// @namespace   https://api.github.com/users/btjoker
// @description 简单的添加为id赋予url属性
// @include     https://h.nimingban.com/*
// @include     http://h.adnmb.com/*
// @version     1
// @grant       none
// ==/UserScript==
(function() {
  "use strict";
  var info = document.getElementsByClassName("h-threads-info-uid");
  for (var i=0; i < info.length; i++) {
    urlcreate(info[i]);
  }
})();

function urlcreate(node) {
  var uid = node.innerText;
  var id = uid.split(':')[1];
  var url = "https://www.so.com/s?q=site%3Ah.nimingban.com+" + id + " target=\"_blank\">";
  var urlnode = "<a href=" + url + uid +"</a>";
  node.innerHTML = urlnode;
}