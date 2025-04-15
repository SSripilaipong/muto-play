const codeArea = document.getElementById("rules");
const objectInput = document.getElementById("input");
const resultArea = document.getElementById("result");
const runBtn = document.getElementById("run");
document.addEventListener("keydown", (e) => {
    if (e.key === "Enter" && e.ctrlKey) {
        e.preventDefault();
        runBtn.click();
    }
});

const RULE_KEY = "muto.rules";

// load + save code
codeArea.value = localStorage.getItem(RULE_KEY) || "";
codeArea.addEventListener("input", () => {
    localStorage.setItem(RULE_KEY, codeArea.value);
});

const go = new Go();
WebAssembly.instantiateStreaming(fetch("ide.wasm"), go.importObject).then((result) => {
    go.run(result.instance);

    document.getElementById("run").onclick = () => {
        const rules = codeArea.value;
        const query = objectInput.value.trim();
        if (!query) return;

        try {
            loadCode(rules);
        } catch (err) {
            resultArea.value += `⚠️ compile error: ${err}\n\n`;
        }

        resultArea.value += `µ> ${query}\n`;
        try {
            const output = execute(query);
            resultArea.value += output.length > 0 ? `${output}\n\n` : '\n';
            resultArea.scrollTop = resultArea.scrollHeight;
        } catch (err) {
            resultArea.value += `⚠️ execution error: ${err}\n\n`;
        }

        objectInput.value = "";
        objectInput.focus();
    };
});
