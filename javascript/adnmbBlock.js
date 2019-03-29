// ==UserScript==
// @name         匿名版板块过滤
// @namespace    http://tampermonkey.net/
// @version      0.1
// @description  板块屏蔽
// @author       btjoker
// @match        https://adnmb1.com/f/时间线
// @match        https://adnmb1.com/f/%E6%97%B6%E9%97%B4%E7%BA%BF
// @grant        none
// ==/UserScript==

(function() {
    'use strict';
    let postItem = document.querySelectorAll("div.h-threads-item");
    let getlateName = x => x.querySelector("a.h-threads-info-id").nextElementSibling.textContent.toLowerCase()
    let getlateID = x => x.querySelector("a.h-threads-info-id").textContent;
    let filterItem = ['lgbt', '小说', 'dota', 'lol'];
    let filterID = ['No.16321668', 'No.14500641'];

    for (let i of postItem) {
        let Name = getlateName(i);
        let id = getlateID(i)
        if (filterID.includes(id) || filterItem.includes(Name)) {
            i.style = 'display:none';
            // console.log(id,Name);
        }
    }
})();

// document.querySelector("div.quan_div1").querySelector("h3")