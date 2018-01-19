// ==UserScript==
// @name         寂月神社
// @namespace    http://tampermonkey.net/
// @version      0.1
// @description  将模糊的图片url修改为正常图url
// @author       btjoker
// @match        https://www.jiyue.com/category/*
// @grant        none
// ==/UserScript==


(function() {
    'use strict';
    let imgList1 = document.querySelectorAll('a.background-container');
    let imgList2 = document.querySelectorAll('div.post-box');
    for (let i of imgList1) {
        i.style.backgroundImage = i.style.backgroundImage.replace('-15x15', '-350x350');
    }
    for (let i of imgList2) {
        i.style.backgroundImage = i.style.backgroundImage.replace('-15x15', '-350x350');
    }
})();