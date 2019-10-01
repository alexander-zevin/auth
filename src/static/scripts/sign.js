function signin() {

    // Устанавливаем параметры запросу
    let login = document.getElementsByName("login")[0].value;
    let password = document.getElementsByName("password")[0].value;
    let obj = '{"login":"'+login+'", "password":"'+password+'"}'
    let params = 'obj=' + encodeURIComponent(obj);
    let url = "http://localhost:8181/signin";
  
    //Отправляем
    let xhttp = request(params, url)
  
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
  
          //Убираем спинер
          blockLoader.style.display = "none";
  
          let data = JSON.parse(this.responseText);
          if (data.MsgErr == "") {
            main.outerHTML = data.HTML;
            document.getElementsByTagName("title")[0].innerHTML = data.Title;
            сhangeURL();
            event();
          } else {
            document.getElementsByClassName("Err")[0].innerHTML = data.MsgErr;
          }
        }
      };
  }

function signup() {

    // Устанавливаем параметры запросу
    let login = document.getElementsByName("login")[0].value;
    let password = document.getElementsByName("password")[0].value;
    let obj = '{"login":"'+login+'", "password":"'+password+'"}'
    let params = 'obj=' + encodeURIComponent(obj);
    let url = "http://localhost:8181/signup";

    if ((login.length < 4) || (password.length < 4)) {
        document.getElementsByClassName("Err")[0].innerHTML = "Login and password must not be less than 4 characters";
    } else {

    //Отправляем
    let xhttp = request(params, url)

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
            let data = JSON.parse(this.responseText);
            if (data.ReLogin == false) {
                document.getElementsByTagName("main")[0].outerHTML = data.HTML;
                document.getElementsByTagName("title")[0].innerHTML = "Main";
                document.cookie = "token="+data.Token;
                сhangeURL();
                event();
            } else {
                document.getElementsByClassName("Err")[0].innerHTML = "A user with this login is already registered";
            }
        }
    };
    }
}

function signOut(){
    changeHTML("auth", true);
}