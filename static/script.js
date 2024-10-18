const ID_JOINER = "->";
const EDITOR_CONTAINER = "editor";

function getState() {
  const elements = document.getElementById("base").children;

  const stateObject = {};
  for (const element of elements) {
    const key = element.getAttribute("data-button-key");
    const vals = element.getAttribute("data-button-key-values");
    console.log(key, vals);
    stateObject[key] = vals;
  }
  console.log(stateObject);

  return stateObject;
}

function getChildrenIds(divElement) {
  const ids = [];
  const children = divElement.children;
  for (let i = 0; i < children.length; i++) {
    const child = children[i];
    if (child?.id) ids.push(child.id);
  }
  return ids;
}
function removeSubElements(parentKey) {
  const parentParts = parentKey.split(ID_JOINER);
  const lastParentPart = parentParts.length - 1;

  const elements = Array.from(
    document.getElementById(EDITOR_CONTAINER)?.children || []
  );

  for (const element of elements) {
    if (element.id) {
      const id = element.id;
      const idParts = id.split(ID_JOINER);

      if (
        parentParts.length < idParts.length &&
        idParts.includes(parentParts[lastParentPart])
      ) {
        element.remove();
      }
    }
  }
}
function createNewCanvas(parentKey, key) {
  removeSubElements(parentKey);
  const newContainer = document.createElement("div");
  newContainer.setAttribute("class", "val_container");
  newContainer.setAttribute("id", parentKey + ID_JOINER + key);
  return newContainer;
}

function getParentId(element) {
  const parent = element.parentElement;
  if (parent && parent.id) {
    return parent.id;
  }
  return null;
}

function parseDataInContainer(containerId) {
  const childElement = document.getElementById(containerId).children;
  const obj = {};
  for (const child of childElement) {
    const key = child.getAttribute("data-button-key");
    const val = child.getAttribute("data-button-key-values");
    obj[key] = JSON.parse(val);
  }

  return obj;
}

function getElementByDataKey(id) {
  return document.querySelector(`[data-button-key='${id}']`);
}

function updateParent(element) {
  const parentId = getParentId(element);
  if (parentId === "base") return;
  const idParts = parentId.split(ID_JOINER);
  const parsedDataInContainer = parseDataInContainer(parentId);
  const targetId = idParts[idParts.length - 1];
  const targetElement = getElementByDataKey(targetId);
  targetElement.setAttribute(
    "data-button-key-values",
    JSON.stringify(parsedDataInContainer)
  );

  updateParent(targetElement);
}

function updateEntry() {
  const enElement = document.querySelector("[data-lang='en']");
  const arElement = document.querySelector("[data-lang='ar']");
  const enVal = enElement.value;
  const arVal = arElement.value;
  const elementId = enElement.getAttribute("data-key");
  const obj = {
    en: enVal,
    ar: arVal,
  };
  const elementToUpdate = document.querySelector(
    `[data-button-key='${elementId}']`
  );
  elementToUpdate.setAttribute("data-button-key-values", JSON.stringify(obj));
  updateParent(elementToUpdate);
}

function createUpdateButton(id) {
  const newBtn = document.createElement("button");
  newBtn.setAttribute("class", "button");
  newBtn.setAttribute("data-button-id", id);
  newBtn.innerText = "Update";
  newBtn.addEventListener("click", updateEntry);
  return newBtn;
}

function createTextarea(id, lang, value) {
  const newTextarea = document.createElement("textarea");
  newTextarea.setAttribute("class", "textarea");
  newTextarea.setAttribute("data-lang", lang);
  newTextarea.setAttribute("data-lang", lang);
  newTextarea.setAttribute("data-key", id);
  newTextarea.value = value;
  return newTextarea;
}
function createContainer(event, entries) {
  const parentId = getParentId(event.target);
  const buttonId = event.target.getAttribute("data-button-key");
  const values = event.target.getAttribute("data-button-key-values");
  const vals = JSON.parse(values);
  const newContainer = createNewCanvas(parentId, buttonId);

  if (typeof vals === "object" && vals?.ar && vals.en) {
    const en = createTextarea(buttonId, "en", vals.en);
    const ar = createTextarea(buttonId, "ar", vals.ar);
    const newBtn = createUpdateButton(buttonId);
    newContainer.appendChild(en);
    newContainer.appendChild(ar);
    newContainer.appendChild(newBtn);
  } else if (typeof vals === "object") {
    Object.keys(vals).forEach((val) => {
      const newBtn = document.createElement("button");
      newBtn.className = "valButton";
      newBtn.innerText = val;
      newBtn.setAttribute("data-button-key", val);
      newBtn.setAttribute("data-button-key-values", JSON.stringify(vals[val]));
      newContainer.appendChild(newBtn);
    });
  } else {
    const newEl = document.createElement("button");
    newEl.className = "valButton";
    newEl.innerText = vals;
    newContainer.appendChild(newEl);
  }

  const parentElement = document.getElementById(EDITOR_CONTAINER);
  parentElement.appendChild(newContainer);

  if (typeof vals === "object" && !vals?.ar && !vals.en)
    addEventListenerToKeyButtons(".valButton");
}

function buttonClicked(event, entries) {
  const buttonId = event.target.getAttribute("data-button-key");
  const values = event.target.getAttribute("data-button-key-values");
  const vals = JSON.parse(values);
  const elemt = document.getElementById("values_container");
  elemt.innerHTML = "";
  if (typeof vals === "object") {
    Object.keys(vals).forEach((val) => {
      const newBtn = document.createElement("button");
      newBtn.className = "valButton";
      newBtn.innerText = val;
      newBtn.setAttribute("data-button-key", val);
      newBtn.setAttribute("data-button-key-values", vals[val]);
      elemt.appendChild(newBtn);
    });
  } else {
    const newEl = document.createElement("button");
    newEl.className = "valButton";
    newEl.innerText = vals;
    elemt.appendChild(newEl);
  }
}

function addEventListenerToKeyButtons(buttonSelector) {
  const keyButtons = document.querySelectorAll(buttonSelector);

  keyButtons.forEach((btn) => {
    btn.addEventListener("click", createContainer);
  });
}

document.addEventListener("DOMContentLoaded", function () {
  const keyButtons = document.querySelectorAll(".keyButton");
  keyButtons.forEach((btn) => {
    btn.addEventListener("click", createContainer);
  });
});
