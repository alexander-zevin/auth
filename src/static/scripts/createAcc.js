let data;
ChangeURL();

// Добавляем события
let createAcc = document.getElementsByClassName("createAcc")[0];
createAcc.addEventListener('click', showSignUp);

let button = document.getElementsByClassName("auth")[0];
button.addEventListener('click', signin);

function showSignUp() {

  // Меняем заголовок
  let titleBlock = document.getElementsByClassName("titleBlock")[0];
  titleBlock.innerHTML = "Create your account";

  // Меняем текст кнопки
  let button = document.getElementsByClassName("auth")[0];
  button.innerHTML = "Sign up";

  // Меняем событие кнопки
  button.removeEventListener('click', signin);
  button.addEventListener('click', signup);

  // Меняем событие span
  createAcc.removeEventListener('click', showSignUp);
  createAcc.addEventListener('click', showSignIn);
  createAcc.innerHTML = "Sign in"
}

function showSignIn() {

  // Меняем заголовок
  let titleBlock = document.getElementsByClassName("titleBlock")[0];
  titleBlock.innerHTML = "Sign in to site";

  // Меняем текст кнопки
  let button = document.getElementsByClassName("auth")[0];
  button.innerHTML = "Sign in";

  // Меняем событие кнопки
  button.removeEventListener('click', signup);
  button.addEventListener('click', signin);

  // Меняем событие span
  createAcc.removeEventListener('click', showSignIn);
  createAcc.addEventListener('click', showSignUp);
  createAcc.innerHTML = "Create an account"
}

function request(params, url) { //Запрос к серверу
  let xhttp = new XMLHttpRequest();
  xhttp.open("POST", url, true);
  xhttp.setRequestHeader("Content-type", "application/x-www-form-urlencoded");
  xhttp.send(params);
  return xhttp;
}

function signin() {
  
}

function signup() {

  // Устанавливаем параметры запросу
  let login = document.getElementsByName("login")[0].value;
  let password = document.getElementsByName("password")[0].value;
  let params = 'login=' + encodeURIComponent(login) + '&password=' + encodeURIComponent(password);
  let url = "http://localhost:8181/signup";

  //Отправляем
  let xhttp = request(params, url)

  //Получаем ответ
  xhttp.onreadystatechange = function() {
      if (xhttp.readyState != 4) return;
      if (xhttp.status != 200) {
          alert(xhttp.status + ': ' + xhttp.statusText);
      } else {
          let data = JSON.parse(this.responseText);
          if (data.ReLogin == false) {
            document.getElementsByTagName("main")[0].outerHTML = data.HTML;
            document.getElementsByTagName("title")[0].innerHTML = "Main";
            document.cookie = "token="+data.Token;
            ChangeURL();
          } else {
            alert("Пользователь с данным логином уже зарешестрирован на этом сайте.");
          }
      }
  };
}

function ChangeURL() { //Меняем URL
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

// window.onpopstate = function(event) {
//   if (document.location == "http://localhost:8181/") {
    
//   }
// };
  