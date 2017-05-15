//
// native.js
//

export class Native {

  /*
  *   get HTML bu template string
  */

  static getHTMLDom(component, data, parent, isRemove = false) {
    let temp = document.createElement('template');
    let result;

    if (temp.content && this.isElement(component)) {
      temp.innerHTML = eval('`' + component.innerHTML + '`');

      if (this.isElement(parent)) {
        parent.appendChild(temp.content);
        result = parent.parentElement.lastElementChild;
      } else {
        component.parentElement.appendChild(temp.content);
        result = component.parentElement.lastElementChild;
      }

      if (isRemove) {
        component.parentNode.removeChild(component);
      }

    } else {
      console.log(`It's not dom component: ${ component }`);
    }
    return result;
  }


  /*
  *   add parse to Dynamic Component
  */

  static reChangeDomDynamically(component) {
    if (this.isElement(component)) {
      Parse.parsComponents(component);
    } else {
      console.log(`This component is not dom: `, component);
    }
  }


  /*
  *   current Id dynamically page
  */

  static get getIdCurrentPage() {
    return Parse.idCurrentPage;
  }


  /*
  *   parse JSON is safely
  */

  static jsonParse(response) {
    try {
      return JSON.parse(response);
    } catch(e) {
      console.log(e);
      alert(e); // error in the above string (in this case, yes)!
    }
  }


  /*
   *  get and post request with callback
   */

  static request(url, callback, params = {}) {

    let method = 'GET';

    if (params) {
      method = 'POST';
    }

    const XHR = ('onload' in new XMLHttpRequest()) ? XMLHttpRequest : XDomainRequest;
    const xhr = new XHR();
    xhr.open(method, url, true);

    xhr.setRequestHeader("X-Requested-With", "XMLHttpRequest");

    xhr.send(params);

    xhr.onload = (response) => {
      if (callback) {
        callback(response.currentTarget.responseText, url);
      } else {
        Observer.emit(Variables.responseToRequest, response.currentTarget.responseText, url);
      }
    };

    xhr.onerror = function () {
      console.log(`Error API to url ${ url } : ${ this }`);
    };

  }


  /*
  *   Set Value Data By Attribute to Dom
  */

  static setValueDataByAttr(data = {}) {

    let obj = data['fields'];
    ParseJSON.parseDataGet(obj, ParseJSON.setAttrToComponent.bind(ParseJSON));

    obj = data['form'];
    const element = document.getElementById(obj['id']);

    if (this.isElement(element)) {
      for (let key in obj) {
        element.setAttribute(key, obj[key]);
      }
    }

    obj = data['data'];
    for (let key in obj) {
      ParseJSON.parseDataGet(obj[key], ParseJSON.insertValueToComponent.bind(ParseJSON), '', true);
    }

  }


  /*
   *  get Value Data By Attributes from Dom
   */

  static getValueDataByAttributes(dom, attr = '', data = {}) {
    const elements = dom.getAttribute(attr);
    for (let element of elements) {
      for (let key in data) {
        element.setAttribute(key, data[key]);
      }
    }
  }


  /*
   *  returns true if it is a DOM element
  */

  static isElement(obj) {
    return (
      typeof HTMLElement === "object" ? obj instanceof HTMLElement : //DOM2
        obj && typeof obj === "object" && obj !== null && obj.nodeType === 1 && typeof obj.nodeName==="string"
    );
  }


  /*
  *   Find first ancestor by class
  */

  static findAncestorByClass(element, className) {
    if (this.isElement(element) && typeof className === 'string') {
      while (!element.classList.contains(className) && (element = element.parentElement));
    }
    return element;
  }



}
