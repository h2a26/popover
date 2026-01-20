(() => {
    (function () {
        "use strict";

        function u(a, e) {
            document.querySelectorAll(`[data-tui-tabs-trigger][data-tui-tabs-id="${a}"]`).forEach((t) => {
                let i = t.getAttribute("data-tui-tabs-value") === e;
                t.setAttribute("data-tui-tabs-state", i ? "active" : "inactive");
            }),
                document.querySelectorAll(`[data-tui-tabs-content][data-tui-tabs-id="${a}"]`).forEach((t) => {
                    let i = t.getAttribute("data-tui-tabs-value") === e;
                    t.setAttribute("data-tui-tabs-state", i ? "active" : "inactive"), t.classList.toggle("hidden", !i);
                });
        }

        document.addEventListener("click", (a) => {
            let e = a.target.closest("[data-tui-tabs-trigger]");
            if (!e) return;
            let t = e.getAttribute("data-tui-tabs-id"),
                i = e.getAttribute("data-tui-tabs-value");
            t && i && u(t, i);
        });

        function s() {
            document.querySelectorAll("[data-tui-tabs]").forEach((tabs) => {
                const tabsId = tabs.getAttribute("data-tui-tabs-id");
                const active = tabs.querySelector('[data-tui-tabs-trigger][data-tui-tabs-state="active"]');
                const input = document.querySelector(`[data-tui-tabs-input="${tabsId}"]`);

                if (active && input) {
                    input.value = active.getAttribute("data-tui-tabs-value");
                }
            });
        }

        document.addEventListener("DOMContentLoaded", s),
            new MutationObserver(s).observe(document.body, {childList: !0, subtree: !0}),
            (window.tui = window.tui || {}),
            (window.tui.tabs = {setActive: u});

        document.addEventListener("click", (e) => {
            const trigger = e.target.closest("[data-tui-tabs-trigger]");
            if (!trigger) return;

            const tabsId = trigger.getAttribute("data-tui-tabs-id");
            const value = trigger.getAttribute("data-tui-tabs-value");

            if (!tabsId || !value) return;

            // convention: <input id="{tabsId}-value">
            const input = document.querySelector(
                `input[type="hidden"][data-tui-tabs-input="${tabsId}"]`
            );

            if (input) {
                input.value = value;
            }
        });

    })();
})();


(function () {
    let lastGenerationType = null;

    function toggleCIF() {
        const input = document.querySelector(
            "input[data-tui-tabs-input='generation_type-tabs']"
        );
        const container = document.getElementById("cif-field-container");

        if (!input || !container) return;

        if (input.value === "existing") {
            container.classList.remove("hidden");
        } else {
            container.classList.add("hidden");
        }
    }

    // run once on load (ONLY show/hide, no reset)
    document.addEventListener("DOMContentLoaded", () => {
        const input = document.querySelector(
            "input[data-tui-tabs-input='generation_type-tabs']"
        );
        if (input) {
            lastGenerationType = input.value;
        }
        toggleCIF();
    });

    // listen ONLY to real tab clicks
    document.addEventListener("click", (e) => {
        const trigger = e.target.closest("[data-tui-tabs-trigger]");
        if (!trigger) return;

        const tabsId = trigger.getAttribute("data-tui-tabs-id");
        const value = trigger.getAttribute("data-tui-tabs-value");

        if (tabsId !== "generation_type-tabs" || !value) return;

        const input = document.querySelector(
            "input[data-tui-tabs-input='generation_type-tabs']"
        );
        const cifInput = document.querySelector("input[name='cif']");

        if (!input) return;

        const previous = lastGenerationType;
        lastGenerationType = value;

        // let tabs.js update the hidden input first
        setTimeout(() => {
            toggleCIF();

            // ✅ reset CIF ONLY when switching TO existing
            if (value === "existing" && previous !== "existing") {
                if (cifInput) {
                    cifInput.value = "";
                }
            }
        }, 0);
    });
})();
