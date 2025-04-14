function saveToFile() {
  const content = output.getValue();
  if (
    content !== "" &&
    content !== "// The generated query will appear here" &&
    content !==
      "// Sending request to Overpass, this can take some time depending on the query..."
  ) {
    const mode = output.getSession().getMode().$id;
    const fileName = `result.${mode === "ace/mode/json" ? "json" : "txt"}`;
    const theFile = new Blob([content], {
      type: `${mode === "ace/mode/json" ? "application/json" : "text/plain"}`,
    });
    const element = document.createElement("a");
    element.setAttribute("href", window.URL.createObjectURL(theFile));
    element.setAttribute("download", fileName);
    element.style.display = "none";
    document.body.appendChild(element);
    element.click();
    document.body.removeChild(element);
  }
}

function copyOutput() {
  const content = output.getValue();
  if (
    content !== "" &&
    content !== "// The generated query will appear here" &&
    content !==
      "// Sending request to Overpass, this can take some time depending on the query..."
  ) {
    navigator.clipboard.writeText(content);
    document.getElementById("copy").innerText = "Copied";
    document.getElementById("copy-icon").classList.add("hidden");
    document.getElementById("copied-icon").classList.remove("hidden");

    setTimeout(() => {
      document.getElementById("copy").innerText = "Copy";
      document.getElementById("copied-icon").classList.add("hidden");
      document.getElementById("copy-icon").classList.remove("hidden");
    }, 2000);
  }
}

function editButtons(enable) {
  const buttons = [
    document.getElementById("save-btn"),
    document.getElementById("copy-btn"),
  ];

  for (btn of buttons) {
    switch (enable) {
      case true:
        btn.classList.add("hover:bg-gray-700", "hover:cursor-pointer");
        btn.classList.remove("opacity-50", "cursor-not-allowed");
        btn.disabled = false;
        break;
      case false:
        btn.classList.remove("hover:bg-gray-700", "hover:cursor-pointer");
        btn.classList.add("opacity-50", "cursor-not-allowed");
        btn.disabled = true;
        break;
    }
  }
}

const go = new Go();

async function execute() {
  const input = editor.getValue();
  let result = "";
  let query = await window.convertYAMLToOQL(input);
  if (query.startsWith("Error parsing YAML: ")) {
    output.container.style.color = "#eb2a2a";
    query = query.replace("Error parsing YAML: ", "");
    let match = query.match(/\(allowed: ([^)]+)\)/);
    if (match) {
      let list = match[1];
      let items = list.split(", ");
      let possibleValues = items.join("\n* ");
      query = query.replace(` ${match[0]}`, `, allowed:\n* ${possibleValues}`);
    }
    query = query.charAt(0).toUpperCase() + query.slice(1);
    result = query;
    editButtons(false);
  } else {
    output.container.style.color = "gray";
    result = "// Sending request to Overpass, this can take some time...";
    let formData = new FormData();
    formData.append("data", query);
    fetch("https://overpass-api.de/api/interpreter", {
      body: new URLSearchParams(formData),
      method: "POST",
      headers: {
        "Content-Type": "application/x-www-form-urlencoded",
      },
    })
      .then((r) => {
        return r.json();
      })
      .then((r) => {
        output.container.style.color = "white";
        output.getSession().setMode("ace/mode/json");
        output.getSession().setValue(JSON.stringify(r.elements, null, "  "));
        editButtons(true);
      })
      .catch((e) => {
        output.container.style.color = "#eb2a2a";
        output.getSession().setMode("ace/mode/text");
        output.getSession().setValue(e.toString());
        editButtons(false);
      });
  }

  editButtons(false);
  output.getSession().setMode("ace/mode/text");
  output.getSession().setValue(result);
}

async function generate() {
  const input = editor.getValue();
  let result = await window.convertYAMLToOQL(input);
  if (result.startsWith("Error parsing YAML: ")) {
    output.container.style.color = "#eb2a2a";
    result = result.replace("Error parsing YAML: ", "");
    let match = result.match(/\(allowed: ([^)]+)\)/);
    if (match) {
      let list = match[1];
      let items = list.split(", ");
      let possibleValues = items.join("\n* ");
      result = result.replace(
        ` ${match[0]}`,
        `, allowed:\n* ${possibleValues}`,
      );
    }
    result = result.charAt(0).toUpperCase() + result.slice(1);
    editButtons(false);
  } else {
    output.container.style.color = "white";
    editButtons(true);
  }

  output.getSession().setMode("ace/mode/text");
  output.getSession().setValue(result);
}

async function renderApp() {
  const result = await window.getVersion();
  document.getElementById("version").innerText = `(v${result})`;
  document.getElementById("app").classList.remove("hidden");
  editButtons(false);
}

WebAssembly.instantiateStreaming(fetch("overscry.wasm"), go.importObject)
  .then((result) => {
    go.run(result.instance);
    document.getElementById("loading").classList.add("hidden");
    renderApp();
  })
  .catch((e) => {
    document.getElementById("loading").classList.add("hidden");
    document.getElementById("error-loading").classList.remove("hidden");
  });
