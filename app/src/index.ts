
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

  
  const observer = new MutationObserver((mutations) => {
    mutations.forEach((mutation) => {
      if (mutation.type === 'childList') {
        console.log('Content changed!');
        // Get all elements that have data-function attribute
        const elements = document.querySelectorAll('[data-function]');
        console.log(elements);
        // attach a click event listener to each element
        elements.forEach((element) => {
          const functionName = element.getAttribute('data-function');
          console.log(functionName);
          element.addEventListener('click', () => {
            request(url + functionName, appContent);
          });
        });
      }
    });
  });
  
  observer.observe(appContent, { childList: true, subtree: true });
  
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
