// ==UserScript==
// @name         FastGetMagnet
// @namespace    https://github.com/btjoker
// @version      0.1
// @description  快速获取 acgnx 的种子链接, 仅限搜索页
// @author       btjoker
// @match        https://share.acgnx.se/search.php*
// @grant        none
// ==/UserScript==

(function() {
    'use strict';
    // 等待加载后执行
    setTimeout(function() {
        let button = createElement("button", "dokidoki", appNode);
        document.querySelector("#bangumi").appendChild(button);
    }, 2000);
})();

function createElement(elm, text, fn) {
    let element = document.createElement(elm);
    let elementText = document.createTextNode(text);
    element.appendChild(elementText);
    if (typeof fn === "function") {
        element.addEventListener('click', fn, false);
    }
    return element
}

function appNode() {
    let list = [...document.querySelectorAll("tbody#data_list tr")];
    let targetNode = document.querySelector("div.box.clear");
    let magnets = list.map(x=>x.querySelector("td:nth-child(5) a").href);
    let p = createElement("textarea", magnets.join("\n"));
    targetNode.appendChild(p);
}