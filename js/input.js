(function () {
    "use strict";

    document.addEventListener("click", (event) => {
        const toggleBtn = event.target.closest("[data-tui-input-toggle-password]");
        if (!toggleBtn) return;

        event.preventDefault();

        const inputId = toggleBtn.getAttribute("data-tui-input-toggle-password");
        const passwordInput = document.getElementById(inputId);
        if (!passwordInput) return;

        const eye = toggleBtn.querySelector(".icon-open");
        const eyeOff = toggleBtn.querySelector(".icon-closed");

        if (passwordInput.type === "password") {
            // SHOW PASSWORD
            passwordInput.type = "text";

            eye?.classList.remove("hidden");
            eyeOff?.classList.add("hidden");
        } else {
            // HIDE PASSWORD
            passwordInput.type = "password";

            eye?.classList.add("hidden");
            eyeOff?.classList.remove("hidden");
        }
    });
})();
