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

const textAreaKeyDownHandler = function(e) {
    if (e.key === 'Tab') {
        e.preventDefault();

        const textarea = e.target;
        const start = textarea.selectionStart;
        const end = textarea.selectionEnd;

        // Insert two spaces
        const value = textarea.value;
        textarea.value = value.substring(0, start) + '  ' + value.substring(end);

        // Move cursor after the inserted spaces
        textarea.selectionStart = textarea.selectionEnd = start + 2;
    }
}

codeArea.addEventListener('keydown', textAreaKeyDownHandler);
objectInput.addEventListener('keydown', textAreaKeyDownHandler);

const RULE_KEY = "muto.rules";


const defaultCode = String.raw`
example-1 = + 1 2


example-2 = f1 3
f1 A = string (* 10 A)


example-3 = f2 9
f2 = * 100


example-4 = sum ($ 1 2 3)
sum (_ Y)        = ret Y
sum (T A B S...) = sum (T (+ A B) S...)


example-5 = g 5
g = (compose string f2)


example-6 = \X [+ 100 X] 23


example-7 = (compose
  (curry + 6)
  \X [* 10 X]
) 45


example-8 = $
  (try t1 999)
  (try t2 555)
  (try t2 555 666)
  (try t3 222)

t1 X   = ret 123
t2 X Y = ret "abc"
t3 111 = ret true


example-9 = h ($ 1 "2" 3 4 "5")
h = (compose
  (curry map string)
  (curry map (curry * 10))
  (curry filter number?)
)


example-10 = queries ({.x: 123, .y: 999.888} (.set .y "replaced!"))
queries D = $
  (D (.get .x))
  (D (.get .y))


example-11 = (new-person "Shane" "Lnw" 1998) (.age-at 2025)

new-person FirstName LastName BirthYear = person {
  .first:      FirstName,
  .last:       LastName,
  .birth-year: BirthYear,
}

(person P) M = (match
  \.first         [P (.get .first)]
  \.last          [P (.get .last)]
  \(.age-at Year) [- Year (P (.get .birth-year))]
) M
`.trimStart()

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

        const { err: compileError } = loadCode(rules);
        if (compileError) {
            resultArea.value += `⚠️ compile error: ${compileError}\n\n`;
        } else {
            resultArea.value += `µ> (${query})\n`;
            const {result: output, err: executeError} = execute(query);
            if (executeError) {
                resultArea.value += `⚠️ execution error: ${executeError}\n\n`;
            } else {
                resultArea.value += output.length > 0 ? `${output}\n\n` : '\n';
            }
        }

        resultArea.scrollTop = resultArea.scrollHeight;
        objectInput.value = "";
        objectInput.focus();
    };
});
