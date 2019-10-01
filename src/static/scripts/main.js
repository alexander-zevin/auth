
function event() {
    // Добавляем события
    if (document.getElementsByTagName("title")[0].innerHTML == "Auth") {
      let createAcc = document.getElementsByClassName("createAcc")[0];
      createAcc.addEventListener('click', showSignUp);
  
      let button = document.getElementsByClassName("auth")[0];
      button.addEventListener('click', signin);
    }
  
    if (document.getElementsByTagName("title")[0].innerHTML == "Main") {
      let signOutBtn = document.getElementsByClassName("signOut")[0];
      signOutBtn.addEventListener('click', signOut);
    }
}

function request(params, url) { //Запрос к серверу
    let xhttp = new XMLHttpRequest();
    xhttp.open("POST", url, true);
    xhttp.setRequestHeader("Content-type", "application/x-www-form-urlencoded");
    xhttp.send(params);
    return xhttp;
}

function сhangeURL() { //Меняем URL
    let title = document.getElementsByTagName("title")[0].innerHTML;
    switch(title) {
      case "Auth":
        history.pushState(null, null, '/auth');
        break;
      case "Main":
        history.pushState(null, null, '/');
        break;
    }
}

window.onpopstate = function(event) {
    let page;
    if (document.location == "http://localhost:8181/auth") {
      page = "auth";
    } else if (document.location == "http://localhost:8181/") {
      page = "main";
    }
    this.changeHTML(page, false);
};

function changeHTML(page, delCookie) {
    // Устанавливаем параметр запросу
    let url = "http://localhost:8181/changeHTML";
    var obj = '{ "page":"' + page + '", "delCookie":' + delCookie + '}';
    let params = 'obj=' + encodeURIComponent(obj);
    
    //Отправляем
    let xhttp = request(params, url);
  
    //Показываем спинер
    let main = document.getElementsByTagName("main")[0];
    let blockLoader = document.getElementsByClassName("blockLoader")[0];
    blockLoader.style.display = "flex";
  
    //Получаем ответ
    xhttp.onreadystatechange = function() {
      if (xhttp.readyState != 4) return;
      if (xhttp.status != 200) {
          alert(xhttp.status + ': ' + xhttp.statusText);
      } else {
          blockLoader.style.display = "none";
          let d = JSON.parse(this.responseText);
          main.outerHTML = d.HTML;
          document.getElementsByTagName("title")[0].innerHTML = d.Title;
          сhangeURL();
          event();
        }
    };
}

event();

сhangeURL();
