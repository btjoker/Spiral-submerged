// hacg右键选中神秘代码，然后点击本书签跳转到云，再次点击自动完成输入搬运8酱的默认密码。
javascript:if(window.location.pathname=="/share/init"){(function(){document.getElementById("accessCode").value="8酱";document.getElementById("submitBtn").onclick();})();}else{var q=window.getSelection();window.open('http://pan.baidu.com/s/'+q,'_blank');}

//获取acfun主页 banner图片，需要在acfun主页点击本书签，打开的新标签就是banner
javascript:var z = document.querySelector('.header-banner style').innerHTML; var x = z.split('url(')[1].split(');')[0]; window.open(x,'_blakn');



// acfix的备份
javascript:(function(){var f=document.createElement('script');f.src='http://www.talkshowcn.com/js/acfunfix.js?ran='+new Date().getTime();document.body.appendChild(f);})();


// 一个Dm1080自动获取神秘代码，在神秘代码后面添加跳转链接的油猴脚本

// ==UserScript==
// @name         Dm1080
// @namespace    http://tampermonkey.net/
// @version      0.1
// @description  try to take over the world!
// @author       You
// @match        http://dm1080p.com/archives/*
// @grant        none
// ==/UserScript==

(function() {
var dmzz = /[A-Z0-9]{7}/;
var dmbq = document.querySelectorAll('span');
for(let i=0;i<dmbq.length;i++){
	if(dmzz.test(dmbq[i].outerText)){
		var tj = dmzz.exec(dmbq[i].outerText);
		var v1 = document.createElement('a');
		v1.setAttribute('href','http://zzzpan.com/?/file/view-'+tj[0]+'.html');
		v1.innerText = '按钮';
		dmbq[i].appendChild(v1);
	}
}
})();