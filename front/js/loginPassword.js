document.addEventListener("DOMContentLoaded", () => {
    const form = document.getElementById("loginForm");
    const loginField = document.getElementById("login");
    const passwordField = document.getElementById("password");
    const errorBox = document.getElementById("errorMessage");
    const submitBtn = form.querySelector(".btn");

    let timeoutId;

    // Прогресс-бар (как в регистрации)
    const progressBar = document.createElement("div");
    progressBar.className = "progress-bar";
    progressBar.id = "progressBar";
    progressBar.innerHTML = '<div class="progress-bar-fill"></div>';
    form.appendChild(progressBar);

    // Сообщение об успехе
    const successMessage = document.createElement("div");
    successMessage.className = "success-message";
    successMessage.id = "successMessage";
    successMessage.textContent = "Вход выполнен! Перенаправляем...";
    form.appendChild(successMessage);

    function showError(message) {
        if (timeoutId) clearTimeout(timeoutId);

        errorBox.textContent = message;
        errorBox.classList.add("active");

        loginField.classList.add("error");
        passwordField.classList.add("error");

        timeoutId = setTimeout(() => {
            errorBox.classList.remove("active");
            loginField.classList.remove("error");
            passwordField.classList.remove("error");
        }, 3000);
    }

    function clearError() {
        errorBox.classList.remove("active");
        loginField.classList.remove("error");
        passwordField.classList.remove("error");
    }

    loginField.addEventListener("input", clearError);
    passwordField.addEventListener("input", clearError);

    form.addEventListener("submit", async (e) => {
        e.preventDefault();

        if (!loginField.value.trim()) {
            showError("Введите email или username");
            loginField.focus();
            return;
        }
        if (!passwordField.value.trim()) {
            showError("Введите пароль");
            passwordField.focus();
            return;
        }

        // Запускаем анимацию
        submitBtn.classList.add("loading");
        form.classList.add("form-sending");
        progressBar.classList.add("active");
        clearError();

        try {
            const response = await fetch("/login", {
                method: "POST",
                body: new FormData(form),
            });

            // (Опционально) небольшая задержка, чтобы анимация точно успела показаться
            await new Promise((r) => setTimeout(r, 800));

            if (!response.ok) throw new Error("Login failed");

            // Успех
            submitBtn.classList.remove("loading");
            submitBtn.classList.add("success");
            submitBtn.textContent = "✓ Успешно";

            progressBar.classList.remove("active");
            successMessage.classList.add("active");

            setTimeout(() => {
                window.location.href = "/success" //поменяйте на нужный URL
            }, 1200);
        } catch (err) {
            // Ошибка
            submitBtn.classList.remove("loading");
            submitBtn.classList.add("error");
            submitBtn.textContent = "Ошибка";

            progressBar.classList.remove("active");
            form.classList.remove("form-sending");

            showError("Неверный email/username или пароль");

            setTimeout(() => {
                submitBtn.classList.remove("error");
                submitBtn.textContent = "Войти";
            }, 1500);
        }
    });
});