const fs = require("fs");
const path = require("path");

const replacements = {
  accordion: "Accordion",
  alert: "Alert",
  breadcrumb: "Breadcrumb",
  "scroll-area": "ScrollArea",
  textarea: "Textarea",
  avatar: "Avatar",
  table: "Table",
  card: "Card",
  button: "Button",
  tabs: "Tabs",
  skeleton: "Skeleton",
  sheet: "Sheet",
  sidebar: "Sidebar",
  select: "Select",
  input: "Input",
  checkbox: "Checkbox",
  form: "Form",
  dialog: "Dialog",
  badge: "Badge",
  carousel: "Carousel",
  label: "Label",
  switch: "Switch",
  "dropdown-menu": "DropdownMenu",
};

function replaceInFile(filePath) {
  let content = fs.readFileSync(filePath, "utf8");
  let modified = false;

  Object.entries(replacements).forEach(([oldName, newName]) => {
    const oldImport = `from "@/components/ui/${oldName}"`;
    const newImport = `from "@/components/ui/${newName}"`;

    if (content.includes(oldImport)) {
      content = content.replace(new RegExp(oldImport, "g"), newImport);
      modified = true;
    }
  });

  if (modified) {
    fs.writeFileSync(filePath, content);
    console.log(`Updated: ${filePath}`);
  }
}

function scanDirectory(dir) {
  const files = fs.readdirSync(dir);

  files.forEach((file) => {
    const filePath = path.join(dir, file);
    const stat = fs.statSync(filePath);

    if (stat.isDirectory()) {
      scanDirectory(filePath);
    } else if (/\.(jsx?|tsx?)$/.test(file)) {
      replaceInFile(filePath);
    }
  });
}

scanDirectory("./src");
console.log("Import replacement completed!");
