// ==UserScript==
// @name hash_id_get
// @namespace Violentmonkey Scripts
// @match http://www.miobt.com/show-*.html
// @match http://kisssub.org/show-*.html
// @description  匹配分流站帖子, 拼接磁链
// @author       btjoker
// @match http://www.comicat.org/show-*.html
// @grant none
// ==/UserScript==

(function () {
  let reg = /.*?show-(.*?).html/;
  let result = reg.exec(document.URL);
  if (result.length != 2) {
    console.log("正则匹配出错");
  }

  let target = document.querySelector(".location");
  let new_ele = createElement("a", "我是磁链", magnet_gen(result[1]));
  target.appendChild(createElement("br"));
  target.appendChild(new_ele);
})();

function magnet_gen(hash_id) {
  return `magnet:?xt=urn:btih:${hash_id}&tr=http://open.acgtracker.com:1096/announce`
}

function createElement(elm, text, hash_id) {
  let element = document.createElement(elm);
  let elementText = document.createTextNode(text);
  element.appendChild(elementText);
  if (hash_id !== "") {
    element.setAttribute("href", hash_id)
  }
  return element
}