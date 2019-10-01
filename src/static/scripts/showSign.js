

function showSignUp() {

  // Меняем заголовок
  let title = document.getElementsByClassName("title")[0];
  title.innerHTML = "Create your account";

  // Меняем текст кнопки
  let button = document.getElementsByClassName("auth")[0];
  button.innerHTML = "Sign up";

  // Меняем событие кнопки
  button.removeEventListener('click', signin);
  button.addEventListener('click', signup);

  // Меняем событие span
  let createAcc = document.getElementsByClassName("createAcc")[0];
  createAcc.removeEventListener('click', showSignUp);
  createAcc.addEventListener('click', showSignIn);
  createAcc.innerHTML = "Sign in"

  document.getElementsByName("login")[0].value = "";
  document.getElementsByName("password")[0].value = "";
  document.getElementsByClassName("Err")[0].innerHTML = "";
}

function showSignIn() {

  // Меняем заголовок
  let title = document.getElementsByClassName("title")[0];
  title.innerHTML = "Sign in to your account";

  // Меняем текст кнопки
  let button = document.getElementsByClassName("auth")[0];
  button.innerHTML = "Sign in";

  // Меняем событие кнопки
  button.removeEventListener('click', signup);
  button.addEventListener('click', signin);

  // Меняем событие span
  let createAcc = document.getElementsByClassName("createAcc")[0];
  createAcc.removeEventListener('click', showSignIn);
  createAcc.addEventListener('click', showSignUp);
  createAcc.innerHTML = "Create an account"

  document.getElementsByName("login")[0].value = "";
  document.getElementsByName("password")[0].value = "";
  document.getElementsByClassName("Err")[0].innerHTML = "";
}