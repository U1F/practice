
const url = "http://localhost:8081/";
document.addEventListener("DOMContentLoaded", initApp);
function initApp() {
  const navHome = document.getElementById("nav-home");
  if (!navHome) {
    return;
  }
  const appContent = document.getElementsByTagName("app-content")[0];
  if (!appContent) {
    return;
  }

  navHome.addEventListener("click", () => {
    request(url + "button", appContent);
  });
}

function request(url: string, rootElement: Element) {
  fetch(url)
    .then((response) => response.text())
    .then((html) => {
      console.log(html);
      rootElement.innerHTML = html;
    });
}
