

var my_main = function () {
    'use strict';

    let timer = setInterval(function() {
        let xhr = new XMLHttpRequest();
        xhr.open(method, "localhost:12000/pi")
        xhr.send()
        xhr.onload = function() { 
            var element = document.getElementById("pi");
            element.textContent = xhr.response;
        }
    }, 20) // 50 fps ;-)
};

window.onload = my_main;