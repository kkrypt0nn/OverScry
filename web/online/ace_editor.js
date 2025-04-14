var editor = ace.edit("editor");
editor.setTheme("ace/theme/tomorrow_night");
editor.getSession().setMode("ace/mode/yaml");
editor.setOptions({
  tabSize: 2,
  useSoftTabs: true,
});
editor.getSession().setValue(`version: "0.0.7-dev"
author: "Krypton (@kkrypt0nn)"
description: "A query to get apartments with 2 levels"
node:
  building:
    value: apartments
    levels:
      value: 2`);

var output = ace.edit("output");
output.setTheme("ace/theme/tomorrow_night");
output.setOptions({
  readOnly: true,
  showLineNumbers: false,
  showGutter: false,
});
output.container.style.color = "gray";
output.getSession().setValue("// The generated query will appear here");
