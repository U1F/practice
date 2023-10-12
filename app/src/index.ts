const url = "http://localhost:8081/";

let appContent: Element;

function initApp() {
  appContent = document.getElementsByTagName("app-content")[0];
  if (!appContent) {
    return;
  }

  const navHome = document.getElementById("nav-home");
  if (!navHome) {
    return;
  }

  navHome.addEventListener("click", () => {
    request(url + "home", appContent);
  });

  const navAbout = document.getElementById("nav-about");
  if (!navAbout) {
    return;
  }
  navAbout.addEventListener("click", () => {
    request(url + "about", appContent);
  });

  const navContact = document.getElementById("nav-contact");
  if (!navContact) {
    return;
  }
  navContact.addEventListener("click", () => {
    request(url + "contact", appContent);
  });

  const observer = new MutationObserver(observeContent);

  observer.observe(appContent, { childList: true, subtree: true });
}

function request(url: string, rootElement: Element) {
  fetch(url)
    .then((response) => response.text())
    .then((html) => {
      console.log(html);
      rootElement.innerHTML = html;
    });
}
function observeContent(mutations: any[]) {
  mutations.forEach((mutation) => {
    if (mutation.type === "childList") {
      // Get all elements that have data-function attribute
      const elements = document.querySelectorAll("[data-function]");

      // attach a click event listener to each element
      elements.forEach((element) => {
        const functionName = element.getAttribute("data-function");

        let eventName: string;
        if (null === element.getAttribute("data-event")) {
          eventName = "click";
        } else {
          eventName = element.getAttribute("data-event")!;
        }

        let targetID: string;
        if (null === element.getAttribute("data-target")) {
          targetID = "";
        } else {
          targetID = element.getAttribute("data-target")!;
        }

        console.log(functionName);
        const isListenerAttached = element.getAttribute(
          "data-listener-attached"
        );

        if (isListenerAttached) {
          return;
        }
        element.setAttribute("data-listener-attached", "true");
        element.addEventListener(eventName, () => {
          request(url + functionName, appContent);
        });
      });
    }
  });
}

document.addEventListener("DOMContentLoaded", initApp);
