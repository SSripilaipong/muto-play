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


const defaultCode = `
((map F B) (H)) = B
(map F (G Ys...)) (H X Xs...) = (map F (G Ys... (F X))) (H Xs...)
map F A = (map F ($)) A

((filter P B) (H)) = B
(filter P B) (H X Xs...) = (filter-case P B (H Xs...)) (P X) X
(filter-case P (G Ys...) A) true X = (filter P (G Ys... X)) A
(filter-case P B A) false X = (filter P B) A
filter P A = (filter P ($)) A

((fold F Z) (G)) = Z
(fold F Z) (G X Xs...) = (fold F (F Z X)) (G Xs...)
fold F Z A = (fold F Z) A

(curry F S...) X... = F S... X...



try-this-in-the-query = (filter string? ($ 1 "2" 3 4 "5"))
`.trim()

// load + save code
codeArea.value = localStorage.getItem(RULE_KEY) || defaultCode;
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
