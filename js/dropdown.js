(function () {
    "use strict";

    document.addEventListener("click", function (e) {
        const item = e.target.closest("[data-tui-dropdown-item]");
        if (!item) return;

        // Ignore submenu triggers
        if (item.hasAttribute("data-tui-dropdown-submenu-trigger")) return;

        // Find enclosing popover
        const popover = item.closest("[data-tui-popover-id]");
        if (!popover) return;

        const popoverID = popover.getAttribute("data-tui-popover-id");
        if (!popoverID) return;

        // Find trigger associated with this popover
        const trigger = document.querySelector(
            `[data-tui-popover-trigger="${popoverID}"]`
        );
        if (!trigger) return;

        // ---- Resolve value & label ----
        const value = item.getAttribute("data-value");
        const label =
            item.getAttribute("data-label") ??
            item.querySelector("[data-tui-dropdown-item-label]")?.textContent ??
            item.textContent.trim();

        // ---- Update trigger label ----
        const labelEl = trigger.querySelector("[data-tui-dropdown-trigger-label]");
        if (labelEl && label) labelEl.textContent = label;

        // ---- Update hidden input scoped to current dropdown ----
        // Look for hidden input in the closest dropdown wrapper
        const wrapper = trigger.closest("[data-tui-dropdown-wrapper]") || trigger.parentElement;
        if (wrapper && value !== null) {
            const hiddenInput = wrapper.querySelector("[data-tui-dropdown-value-target]");
            if (hiddenInput) {
                hiddenInput.value = value;
                hiddenInput.dispatchEvent(new Event("change", {bubbles: true}));
            }
        }

        /// ---- Prevent auto-close ----
        if (item.getAttribute("data-tui-dropdown-prevent-close") === "true") return;


        if (typeof window.closePopover === "function") {
            window.closePopover(popoverID);
        }
    });
})();
