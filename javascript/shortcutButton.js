// ==UserScript==
// @name         我的钢铁快捷按钮
// @namespace    http://tampermonkey.net/
// @version      0.4
// @description  快捷按钮, 不需要再去下拉查看!
// @author       btjojker
// @match        https://*.mysteel.com/m/*
// @grant        none
// ==/UserScript==

(function () {
    'use strict';
    let blue = "color:#0099FF";
    let black = "color:#000000";
    var data = Array.from(document.querySelectorAll("ul#placeUL input"));
    var a4 = document.querySelector("#a4");

    function change(name) {
        // console.log(name);
        for (let i of data) {
            if (name === i.value) {
                i.click();
                clickColorOn()
            };
        };
    };

    function clickColorOn() {
        // 做个筛选, 获取选中状态的项, 然后返回他们的 值
        let filterSelectList = data.filter(x=>x.checked).map(x=>x.value);

        for (let i of allItem) {
            if (filterSelectList.includes(i.textContent)) {
                i.style = blue;
            } else {
                i.style = black;
            }
        }
    }

    // createF4 生成厂家按钮
    function createF4() {
        if (document.querySelectorAll(".f4").length != 0) {
            return
        }

        // 可以做个排序, 排序字符串, 然后生成按钮
        for (let i of data) {
            let btn = createElement("button", i.value, () => change(i.value));
            btn.className = "f4";
            a4.appendChild(btn);
        }
    };

    createF4();

    let allItem = Array.from(document.querySelectorAll(".f4"));

    // 初始化
    function restoreAll() {
        for (let i of data) {
            if (i.checked) {
                i.click();
            }
        }
        clickColorOn()
    };

    function show() {
        for (let i of allItem) {
            if (i.style === '') return
            i.style = '';
        }
        clickColorOn()
    }

    function hidden() {
        for (let i of allItem) {
            if (i.style === 'display:none') return
            i.style = 'display:none';
        }
    }

    // createButton 创建按钮
    function createButton() {
        let naver = document.querySelector("#naver");
        let restore = createElement("button", "初始化", restoreAll);
        let showBtn = createElement("button", "显示按钮", show);
        let hiddenBtn = createElement("button", "隐藏按钮", hidden);

        naver.appendChild(restore);
        naver.appendChild(hiddenBtn);
        naver.appendChild(showBtn);
    }

    createButton()
})();

function createElement(elm, text, fn) {
    let element = document.createElement(elm);
    let elementText = document.createTextNode(text);
    element.appendChild(elementText);
    element.addEventListener('click', fn, false);
    return element
}

// let a = document.querySelector("tr#ctr293");
// a.style = 'display:none' // 隐藏元素
// a.style = ''; // 取消隐藏

//初始化不冲突 - 不需要改
// 隐藏 -> style = "display:none"
// 显示 -> 显示函数修改 -> 不创建资源 修改 style 值为 "" 并且调用 clickColorOn函数 为按钮文字附上颜色

// 隐藏后就不会有点击了
// 先判断是否隐藏 -> 如果隐藏了, 就对点击操作不做反应
// 初始化的时候会重置选中状态, 并且显示所有的按钮